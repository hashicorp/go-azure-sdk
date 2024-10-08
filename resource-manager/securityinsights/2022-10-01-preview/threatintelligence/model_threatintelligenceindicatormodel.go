package threatintelligence

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ThreatIntelligenceInformation = ThreatIntelligenceIndicatorModel{}

type ThreatIntelligenceIndicatorModel struct {
	Properties *ThreatIntelligenceIndicatorProperties `json:"properties,omitempty"`

	// Fields inherited from ThreatIntelligenceInformation

	Etag       *string                            `json:"etag,omitempty"`
	Id         *string                            `json:"id,omitempty"`
	Kind       ThreatIntelligenceResourceKindEnum `json:"kind"`
	Name       *string                            `json:"name,omitempty"`
	SystemData *systemdata.SystemData             `json:"systemData,omitempty"`
	Type       *string                            `json:"type,omitempty"`
}

func (s ThreatIntelligenceIndicatorModel) ThreatIntelligenceInformation() BaseThreatIntelligenceInformationImpl {
	return BaseThreatIntelligenceInformationImpl{
		Etag:       s.Etag,
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = ThreatIntelligenceIndicatorModel{}

func (s ThreatIntelligenceIndicatorModel) MarshalJSON() ([]byte, error) {
	type wrapper ThreatIntelligenceIndicatorModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ThreatIntelligenceIndicatorModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ThreatIntelligenceIndicatorModel: %+v", err)
	}

	decoded["kind"] = "indicator"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ThreatIntelligenceIndicatorModel: %+v", err)
	}

	return encoded, nil
}
