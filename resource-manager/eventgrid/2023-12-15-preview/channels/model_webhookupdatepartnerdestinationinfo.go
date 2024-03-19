package channels

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PartnerUpdateDestinationInfo = WebhookUpdatePartnerDestinationInfo{}

type WebhookUpdatePartnerDestinationInfo struct {
	Properties *WebhookPartnerDestinationProperties `json:"properties,omitempty"`

	// Fields inherited from PartnerUpdateDestinationInfo
}

var _ json.Marshaler = WebhookUpdatePartnerDestinationInfo{}

func (s WebhookUpdatePartnerDestinationInfo) MarshalJSON() ([]byte, error) {
	type wrapper WebhookUpdatePartnerDestinationInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WebhookUpdatePartnerDestinationInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WebhookUpdatePartnerDestinationInfo: %+v", err)
	}
	decoded["endpointType"] = "WebHook"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WebhookUpdatePartnerDestinationInfo: %+v", err)
	}

	return encoded, nil
}
