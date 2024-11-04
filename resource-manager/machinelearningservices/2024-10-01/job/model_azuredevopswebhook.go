package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Webhook = AzureDevOpsWebhook{}

type AzureDevOpsWebhook struct {

	// Fields inherited from Webhook

	EventType   *string     `json:"eventType,omitempty"`
	WebhookType WebhookType `json:"webhookType"`
}

func (s AzureDevOpsWebhook) Webhook() BaseWebhookImpl {
	return BaseWebhookImpl{
		EventType:   s.EventType,
		WebhookType: s.WebhookType,
	}
}

var _ json.Marshaler = AzureDevOpsWebhook{}

func (s AzureDevOpsWebhook) MarshalJSON() ([]byte, error) {
	type wrapper AzureDevOpsWebhook
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureDevOpsWebhook: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureDevOpsWebhook: %+v", err)
	}

	decoded["webhookType"] = "AzureDevOps"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureDevOpsWebhook: %+v", err)
	}

	return encoded, nil
}
