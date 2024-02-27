package alerts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertProperties struct {
	AlertDisplayName     *string               `json:"alertDisplayName,omitempty"`
	AlertType            *string               `json:"alertType,omitempty"`
	AlertUri             *string               `json:"alertUri,omitempty"`
	CompromisedEntity    *string               `json:"compromisedEntity,omitempty"`
	CorrelationKey       *string               `json:"correlationKey,omitempty"`
	Description          *string               `json:"description,omitempty"`
	EndTimeUtc           *string               `json:"endTimeUtc,omitempty"`
	Entities             *[]AlertEntity        `json:"entities,omitempty"`
	ExtendedLinks        *[]map[string]string  `json:"extendedLinks,omitempty"`
	ExtendedProperties   *map[string]string    `json:"extendedProperties,omitempty"`
	Intent               *Intent               `json:"intent,omitempty"`
	IsIncident           *bool                 `json:"isIncident,omitempty"`
	ProcessingEndTimeUtc *string               `json:"processingEndTimeUtc,omitempty"`
	ProductComponentName *string               `json:"productComponentName,omitempty"`
	ProductName          *string               `json:"productName,omitempty"`
	RemediationSteps     *[]string             `json:"remediationSteps,omitempty"`
	ResourceIdentifiers  *[]ResourceIdentifier `json:"resourceIdentifiers,omitempty"`
	Severity             *AlertSeverity        `json:"severity,omitempty"`
	StartTimeUtc         *string               `json:"startTimeUtc,omitempty"`
	Status               *AlertStatus          `json:"status,omitempty"`
	SystemAlertId        *string               `json:"systemAlertId,omitempty"`
	TimeGeneratedUtc     *string               `json:"timeGeneratedUtc,omitempty"`
	VendorName           *string               `json:"vendorName,omitempty"`
}

func (o *AlertProperties) GetEndTimeUtcAsTime() (*time.Time, error) {
	if o.EndTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *AlertProperties) GetProcessingEndTimeUtcAsTime() (*time.Time, error) {
	if o.ProcessingEndTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ProcessingEndTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *AlertProperties) GetStartTimeUtcAsTime() (*time.Time, error) {
	if o.StartTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *AlertProperties) GetTimeGeneratedUtcAsTime() (*time.Time, error) {
	if o.TimeGeneratedUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TimeGeneratedUtc, "2006-01-02T15:04:05Z07:00")
}

var _ json.Unmarshaler = &AlertProperties{}

func (s *AlertProperties) UnmarshalJSON(bytes []byte) error {
	type alias AlertProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into AlertProperties: %+v", err)
	}

	s.AlertDisplayName = decoded.AlertDisplayName
	s.AlertType = decoded.AlertType
	s.AlertUri = decoded.AlertUri
	s.CompromisedEntity = decoded.CompromisedEntity
	s.CorrelationKey = decoded.CorrelationKey
	s.Description = decoded.Description
	s.EndTimeUtc = decoded.EndTimeUtc
	s.Entities = decoded.Entities
	s.ExtendedLinks = decoded.ExtendedLinks
	s.ExtendedProperties = decoded.ExtendedProperties
	s.Intent = decoded.Intent
	s.IsIncident = decoded.IsIncident
	s.ProcessingEndTimeUtc = decoded.ProcessingEndTimeUtc
	s.ProductComponentName = decoded.ProductComponentName
	s.ProductName = decoded.ProductName
	s.RemediationSteps = decoded.RemediationSteps
	s.Severity = decoded.Severity
	s.StartTimeUtc = decoded.StartTimeUtc
	s.Status = decoded.Status
	s.SystemAlertId = decoded.SystemAlertId
	s.TimeGeneratedUtc = decoded.TimeGeneratedUtc
	s.VendorName = decoded.VendorName

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AlertProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["resourceIdentifiers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ResourceIdentifiers into list []json.RawMessage: %+v", err)
		}

		output := make([]ResourceIdentifier, 0)
		for i, val := range listTemp {
			impl, err := unmarshalResourceIdentifierImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ResourceIdentifiers' for 'AlertProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ResourceIdentifiers = &output
	}
	return nil
}
