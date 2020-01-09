package spinnakerpb

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		obj      interface{}
		expected string
	}{
		{"stage",
			Stage{
				Stage: &Stage_ManualJudgment{
					ManualJudgment: &ManualJudgmentStage{
						Type:  "manualJudgment",
						RefId: "derp",
						Name:  "Derp derp",
					},
				},
			},
			`{"type":"manualJudgment","refId":"derp","name":"Derp derp"}`,
		},
		{
			"notification",
			Notification{
				Notification: &Notification_Pubsub{
					Pubsub: &PubsubNotification{
						Type:          "pubsub",
						Level:         "stage",
						When:          []string{"stage.starting", "stage.failed"},
						PublisherName: "derp",
					},
				},
			},
			`{"type":"pubsub","level":"stage","when":["stage.starting","stage.failed"],"publisherName":"derp"}`,
		},
		{
			"trigger",
			Trigger{
				Trigger: &Trigger_Webhook{
					Webhook: &WebhookTrigger{
						Type:                "webhook",
						Enabled:             true,
						ExpectedArtifactIds: []string{"herp", "derp"},
						Source:              "asdf",
						PayloadConstraints: map[string]string{
							"foo": "bar",
						},
					},
				},
			},
			`{"type":"webhook","enabled":true,"expectedArtifactIds":["herp","derp"],"source":"asdf","payloadConstraints":{"foo":"bar"}}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf, err := json.Marshal(test.obj)
			require.NoError(t, err)
			assert.Equal(t, string(buf), test.expected)
		})
	}
}
