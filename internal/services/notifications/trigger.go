package notifications

import (
	"github.com/Nahom101-1/assignment-2/internal/models"
	"github.com/Nahom101-1/assignment-2/internal/storage"
	"github.com/Nahom101-1/assignment-2/utils"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
)

// TriggerWebhooks Sender post request to specified webhook given event and country match
func TriggerWebhooks(w http.ResponseWriter, r *http.Request, event string, country string) {
	log.Printf("Triggering webhooks for event %s", event)

	// get iterator for collection
	iter := storage.GetClient().Collection("notifications").Documents(r.Context())

	for {
		// Get one document at a time (each iteration)
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error retrieving document: %v", err)
			continue
		}

		// Decode the document into a Webhook struct
		var temp models.Webhook
		if err := doc.DataTo(&temp); err != nil {
			log.Printf("Error decoding document: %v", err)
			continue
		}

		// Check if webhook matches event or country
		if temp.Event == event {
			if temp.Country == "" || temp.Country == country {
				// Construct payload with correct data
				payload := models.Notification{
					ID:      temp.ID,
					Country: country,
					Event:   event,
					Time:    utils.GetTimestamp(),
				}

				// Send POST request to webhook URL
				utils.SendPostRequest(temp.URL, payload)
			}
		}
	}
}
