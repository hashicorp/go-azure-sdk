package assessments

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAssessmentPropertiesResponse struct {
	AdditionalData  *map[string]string                    `json:"additionalData,omitempty"`
	DisplayName     *string                               `json:"displayName,omitempty"`
	Links           *AssessmentLinks                      `json:"links,omitempty"`
	Metadata        *SecurityAssessmentMetadataProperties `json:"metadata,omitempty"`
	PartnersData    *SecurityAssessmentPartnerData        `json:"partnersData,omitempty"`
	ResourceDetails ResourceDetails                       `json:"resourceDetails"`
	Status          AssessmentStatusResponse              `json:"status"`
}

var _ json.Unmarshaler = &SecurityAssessmentPropertiesResponse{}

func (s *SecurityAssessmentPropertiesResponse) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AdditionalData *map[string]string                    `json:"additionalData,omitempty"`
		DisplayName    *string                               `json:"displayName,omitempty"`
		Links          *AssessmentLinks                      `json:"links,omitempty"`
		Metadata       *SecurityAssessmentMetadataProperties `json:"metadata,omitempty"`
		PartnersData   *SecurityAssessmentPartnerData        `json:"partnersData,omitempty"`
		Status         AssessmentStatusResponse              `json:"status"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AdditionalData = decoded.AdditionalData
	s.DisplayName = decoded.DisplayName
	s.Links = decoded.Links
	s.Metadata = decoded.Metadata
	s.PartnersData = decoded.PartnersData
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityAssessmentPropertiesResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["resourceDetails"]; ok {
		impl, err := UnmarshalResourceDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ResourceDetails' for 'SecurityAssessmentPropertiesResponse': %+v", err)
		}
		s.ResourceDetails = impl
	}

	return nil
}
