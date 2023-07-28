package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Webhook interface {
}

func unmarshalWebhookImplementation(input []byte) (Webhook, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Webhook into map[string]interface: %+v", err)
	}

	value, ok := temp["webhookType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureDevOps") {
		var out AzureDevOpsWebhook
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureDevOpsWebhook: %+v", err)
		}
		return out, nil
	}

	type RawWebhookImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawWebhookImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
