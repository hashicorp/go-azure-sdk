package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityIntelligenceProfile{}

type SecurityIntelligenceProfile struct {
	// A list of commonly-known aliases for the threat intelligence included in the intelligenceProfile.
	Aliases *[]string `json:"aliases,omitempty"`

	// The country/region of origin for the given actor or threat associated with this intelligenceProfile.
	CountriesOrRegionsOfOrigin *[]SecurityIntelligenceProfileCountryOrRegionOfOrigin `json:"countriesOrRegionsOfOrigin,omitempty"`

	Description *SecurityFormattedContent `json:"description,omitempty"`

	// The date and time when this intelligenceProfile was first active. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	FirstActiveDateTime *string `json:"firstActiveDateTime,omitempty"`

	// Includes an assemblage of high-fidelity network indicators of compromise.
	Indicators *[]SecurityIntelligenceProfileIndicator `json:"indicators,omitempty"`

	Kind    *SecurityIntelligenceProfileKind `json:"kind,omitempty"`
	Summary *SecurityFormattedContent        `json:"summary,omitempty"`

	// Known targets related to this intelligenceProfile.
	Targets *[]string `json:"targets,omitempty"`

	// The title of this intelligenceProfile.
	Title *string `json:"title,omitempty"`

	// Formatted information featuring a description of the distinctive tactics, techniques, and procedures (TTP) of the
	// group, followed by a list of all known custom, commodity, and publicly available implants used by the group.
	Tradecraft *SecurityFormattedContent `json:"tradecraft,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SecurityIntelligenceProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityIntelligenceProfile{}

func (s SecurityIntelligenceProfile) MarshalJSON() ([]byte, error) {
	type wrapper SecurityIntelligenceProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityIntelligenceProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityIntelligenceProfile: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.intelligenceProfile"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityIntelligenceProfile: %+v", err)
	}

	return encoded, nil
}
