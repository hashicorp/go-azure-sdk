package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartialNotificationSetting struct {
	WebHooks *map[string]Webhook `json:"webhooks,omitempty"`
}

var _ json.Unmarshaler = &PartialNotificationSetting{}

func (s *PartialNotificationSetting) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PartialNotificationSetting into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["webhooks"]; ok {
		var dictionaryTemp map[string]json.RawMessage
		if err := json.Unmarshal(v, &dictionaryTemp); err != nil {
			return fmt.Errorf("unmarshaling WebHooks into dictionary map[string]json.RawMessage: %+v", err)
		}

		output := make(map[string]Webhook)
		for key, val := range dictionaryTemp {
			impl, err := unmarshalWebhookImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling key %q field 'WebHooks' for 'PartialNotificationSetting': %+v", key, err)
			}
			output[key] = impl
		}
		s.WebHooks = &output
	}
	return nil
}
