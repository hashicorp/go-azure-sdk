package job

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Webhook interface {
	Webhook() BaseWebhookImpl
}

var _ Webhook = BaseWebhookImpl{}

type BaseWebhookImpl struct {
	EventType   *string     `json:"eventType,omitempty"`
	WebhookType WebhookType `json:"webhookType"`
}

func (s BaseWebhookImpl) Webhook() BaseWebhookImpl {
	return s
}

var _ Webhook = RawWebhookImpl{}

// RawWebhookImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWebhookImpl struct {
	webhook BaseWebhookImpl
	Type    string
	Values  map[string]interface{}
}

func (s RawWebhookImpl) Webhook() BaseWebhookImpl {
	return s.webhook
}

func UnmarshalWebhookImplementation(input []byte) (Webhook, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Webhook into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["webhookType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureDevOps") {
		var out AzureDevOpsWebhook
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureDevOpsWebhook: %+v", err)
		}
		return out, nil
	}

	var parent BaseWebhookImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWebhookImpl: %+v", err)
	}

	return RawWebhookImpl{
		webhook: parent,
		Type:    value,
		Values:  temp,
	}, nil

}
