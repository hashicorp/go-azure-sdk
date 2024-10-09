package subassessments

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubAssessmentProperties struct {
	AdditionalData  AdditionalData       `json:"additionalData"`
	Category        *string              `json:"category,omitempty"`
	Description     *string              `json:"description,omitempty"`
	DisplayName     *string              `json:"displayName,omitempty"`
	Id              *string              `json:"id,omitempty"`
	Impact          *string              `json:"impact,omitempty"`
	Remediation     *string              `json:"remediation,omitempty"`
	ResourceDetails ResourceDetails      `json:"resourceDetails"`
	Status          *SubAssessmentStatus `json:"status,omitempty"`
	TimeGenerated   *string              `json:"timeGenerated,omitempty"`
}

func (o *SecuritySubAssessmentProperties) GetTimeGeneratedAsTime() (*time.Time, error) {
	if o.TimeGenerated == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TimeGenerated, "2006-01-02T15:04:05Z07:00")
}

func (o *SecuritySubAssessmentProperties) SetTimeGeneratedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TimeGenerated = &formatted
}

var _ json.Unmarshaler = &SecuritySubAssessmentProperties{}

func (s *SecuritySubAssessmentProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Category      *string              `json:"category,omitempty"`
		Description   *string              `json:"description,omitempty"`
		DisplayName   *string              `json:"displayName,omitempty"`
		Id            *string              `json:"id,omitempty"`
		Impact        *string              `json:"impact,omitempty"`
		Remediation   *string              `json:"remediation,omitempty"`
		Status        *SubAssessmentStatus `json:"status,omitempty"`
		TimeGenerated *string              `json:"timeGenerated,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Category = decoded.Category
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.Impact = decoded.Impact
	s.Remediation = decoded.Remediation
	s.Status = decoded.Status
	s.TimeGenerated = decoded.TimeGenerated

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecuritySubAssessmentProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["additionalData"]; ok {
		impl, err := UnmarshalAdditionalDataImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AdditionalData' for 'SecuritySubAssessmentProperties': %+v", err)
		}
		s.AdditionalData = impl
	}

	if v, ok := temp["resourceDetails"]; ok {
		impl, err := UnmarshalResourceDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ResourceDetails' for 'SecuritySubAssessmentProperties': %+v", err)
		}
		s.ResourceDetails = impl
	}

	return nil
}
