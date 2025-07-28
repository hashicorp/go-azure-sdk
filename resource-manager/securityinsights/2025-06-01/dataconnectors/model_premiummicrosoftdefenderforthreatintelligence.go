package dataconnectors

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnector = PremiumMicrosoftDefenderForThreatIntelligence{}

type PremiumMicrosoftDefenderForThreatIntelligence struct {
	Properties *PremiumMdtiDataConnectorProperties `json:"properties,omitempty"`

	// Fields inherited from DataConnector

	Etag       *string                `json:"etag,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Kind       DataConnectorKind      `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s PremiumMicrosoftDefenderForThreatIntelligence) DataConnector() BaseDataConnectorImpl {
	return BaseDataConnectorImpl{
		Etag:       s.Etag,
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = PremiumMicrosoftDefenderForThreatIntelligence{}

func (s PremiumMicrosoftDefenderForThreatIntelligence) MarshalJSON() ([]byte, error) {
	type wrapper PremiumMicrosoftDefenderForThreatIntelligence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PremiumMicrosoftDefenderForThreatIntelligence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PremiumMicrosoftDefenderForThreatIntelligence: %+v", err)
	}

	decoded["kind"] = "PremiumMicrosoftDefenderForThreatIntelligence"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PremiumMicrosoftDefenderForThreatIntelligence: %+v", err)
	}

	return encoded, nil
}
