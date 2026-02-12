package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentGroup struct {
	Attestation GroupAttestation    `json:"attestation"`
	DisplayName string              `json:"displayName"`
	Enabled     *bool               `json:"enabled,omitempty"`
	Etag        *string             `json:"etag,omitempty"`
	Id          *string             `json:"id,omitempty"`
	IdScope     *string             `json:"idScope,omitempty"`
	Type        EnrollmentGroupType `json:"type"`
}

var _ json.Unmarshaler = &EnrollmentGroup{}

func (s *EnrollmentGroup) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName string              `json:"displayName"`
		Enabled     *bool               `json:"enabled,omitempty"`
		Etag        *string             `json:"etag,omitempty"`
		Id          *string             `json:"id,omitempty"`
		IdScope     *string             `json:"idScope,omitempty"`
		Type        EnrollmentGroupType `json:"type"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.Enabled = decoded.Enabled
	s.Etag = decoded.Etag
	s.Id = decoded.Id
	s.IdScope = decoded.IdScope
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EnrollmentGroup into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attestation"]; ok {
		impl, err := UnmarshalGroupAttestationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Attestation' for 'EnrollmentGroup': %+v", err)
		}
		s.Attestation = impl
	}

	return nil
}
