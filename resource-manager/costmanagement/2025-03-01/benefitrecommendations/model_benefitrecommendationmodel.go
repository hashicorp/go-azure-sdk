package benefitrecommendations

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitRecommendationModel struct {
	Id         *string                         `json:"id,omitempty"`
	Kind       *BenefitKind                    `json:"kind,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Properties BenefitRecommendationProperties `json:"properties"`
	Type       *string                         `json:"type,omitempty"`
}

var _ json.Unmarshaler = &BenefitRecommendationModel{}

func (s *BenefitRecommendationModel) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Id   *string      `json:"id,omitempty"`
		Kind *BenefitKind `json:"kind,omitempty"`
		Name *string      `json:"name,omitempty"`
		Type *string      `json:"type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.Kind = decoded.Kind
	s.Name = decoded.Name
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BenefitRecommendationModel into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := UnmarshalBenefitRecommendationPropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'BenefitRecommendationModel': %+v", err)
		}
		s.Properties = impl
	}

	return nil
}
