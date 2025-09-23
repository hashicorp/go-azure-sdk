package currentquotalimitbases

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaProperties struct {
	IsQuotaApplicable *bool           `json:"isQuotaApplicable,omitempty"`
	Limit             LimitJsonObject `json:"limit"`
	Name              *ResourceName   `json:"name,omitempty"`
	Properties        *interface{}    `json:"properties,omitempty"`
	QuotaPeriod       *string         `json:"quotaPeriod,omitempty"`
	ResourceType      *string         `json:"resourceType,omitempty"`
	Unit              *string         `json:"unit,omitempty"`
}

var _ json.Unmarshaler = &QuotaProperties{}

func (s *QuotaProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsQuotaApplicable *bool         `json:"isQuotaApplicable,omitempty"`
		Name              *ResourceName `json:"name,omitempty"`
		Properties        *interface{}  `json:"properties,omitempty"`
		QuotaPeriod       *string       `json:"quotaPeriod,omitempty"`
		ResourceType      *string       `json:"resourceType,omitempty"`
		Unit              *string       `json:"unit,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsQuotaApplicable = decoded.IsQuotaApplicable
	s.Name = decoded.Name
	s.Properties = decoded.Properties
	s.QuotaPeriod = decoded.QuotaPeriod
	s.ResourceType = decoded.ResourceType
	s.Unit = decoded.Unit

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling QuotaProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["limit"]; ok {
		impl, err := UnmarshalLimitJsonObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Limit' for 'QuotaProperties': %+v", err)
		}
		s.Limit = impl
	}

	return nil
}
