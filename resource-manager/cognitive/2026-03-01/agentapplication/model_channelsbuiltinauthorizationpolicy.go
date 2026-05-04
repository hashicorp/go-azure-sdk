package agentapplication

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApplicationAuthorizationPolicy = ChannelsBuiltInAuthorizationPolicy{}

type ChannelsBuiltInAuthorizationPolicy struct {

	// Fields inherited from ApplicationAuthorizationPolicy

	Type BuiltInAuthorizationScheme `json:"type"`
}

func (s ChannelsBuiltInAuthorizationPolicy) ApplicationAuthorizationPolicy() BaseApplicationAuthorizationPolicyImpl {
	return BaseApplicationAuthorizationPolicyImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = ChannelsBuiltInAuthorizationPolicy{}

func (s ChannelsBuiltInAuthorizationPolicy) MarshalJSON() ([]byte, error) {
	type wrapper ChannelsBuiltInAuthorizationPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChannelsBuiltInAuthorizationPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChannelsBuiltInAuthorizationPolicy: %+v", err)
	}

	decoded["type"] = "Channels"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChannelsBuiltInAuthorizationPolicy: %+v", err)
	}

	return encoded, nil
}
