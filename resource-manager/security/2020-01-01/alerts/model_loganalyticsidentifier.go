package alerts

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ResourceIdentifier = LogAnalyticsIdentifier{}

type LogAnalyticsIdentifier struct {
	AgentId                 *string `json:"agentId,omitempty"`
	WorkspaceId             *string `json:"workspaceId,omitempty"`
	WorkspaceResourceGroup  *string `json:"workspaceResourceGroup,omitempty"`
	WorkspaceSubscriptionId *string `json:"workspaceSubscriptionId,omitempty"`

	// Fields inherited from ResourceIdentifier

	Type ResourceIdentifierType `json:"type"`
}

func (s LogAnalyticsIdentifier) ResourceIdentifier() BaseResourceIdentifierImpl {
	return BaseResourceIdentifierImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = LogAnalyticsIdentifier{}

func (s LogAnalyticsIdentifier) MarshalJSON() ([]byte, error) {
	type wrapper LogAnalyticsIdentifier
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LogAnalyticsIdentifier: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LogAnalyticsIdentifier: %+v", err)
	}

	decoded["type"] = "LogAnalytics"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LogAnalyticsIdentifier: %+v", err)
	}

	return encoded, nil
}
