package channel

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Channel = TelephonyChannel{}

type TelephonyChannel struct {
	Properties *TelephonyChannelProperties `json:"properties,omitempty"`

	// Fields inherited from Channel
	Etag              *string `json:"etag,omitempty"`
	Location          *string `json:"location,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

var _ json.Marshaler = TelephonyChannel{}

func (s TelephonyChannel) MarshalJSON() ([]byte, error) {
	type wrapper TelephonyChannel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TelephonyChannel: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TelephonyChannel: %+v", err)
	}
	decoded["channelName"] = "TelephonyChannel"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TelephonyChannel: %+v", err)
	}

	return encoded, nil
}
