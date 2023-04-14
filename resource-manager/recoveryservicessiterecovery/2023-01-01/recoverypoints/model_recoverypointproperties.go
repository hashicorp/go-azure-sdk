package recoverypoints

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointProperties struct {
	ProviderSpecificDetails *ProviderSpecificRecoveryPointDetails `json:"providerSpecificDetails,omitempty"`
	RecoveryPointTime       *string                               `json:"recoveryPointTime,omitempty"`
	RecoveryPointType       *string                               `json:"recoveryPointType,omitempty"`
}

func (o *RecoveryPointProperties) GetRecoveryPointTimeAsTime() (*time.Time, error) {
	if o.RecoveryPointTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RecoveryPointTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecoveryPointProperties) SetRecoveryPointTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RecoveryPointTime = &formatted
}

var _ json.Unmarshaler = &RecoveryPointProperties{}

func (s *RecoveryPointProperties) UnmarshalJSON(bytes []byte) error {
	type alias RecoveryPointProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into RecoveryPointProperties: %+v", err)
	}

	s.RecoveryPointTime = decoded.RecoveryPointTime
	s.RecoveryPointType = decoded.RecoveryPointType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling RecoveryPointProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := unmarshalProviderSpecificRecoveryPointDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'RecoveryPointProperties': %+v", err)
		}
		s.ProviderSpecificDetails = impl
	}
	return nil
}
