package quotarequests

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubRequest struct {
	Limit             LimitJsonObject    `json:"limit"`
	Message           *string            `json:"message,omitempty"`
	Name              *ResourceName      `json:"name,omitempty"`
	ProvisioningState *QuotaRequestState `json:"provisioningState,omitempty"`
	ResourceType      *string            `json:"resourceType,omitempty"`
	SubRequestId      *string            `json:"subRequestId,omitempty"`
	Unit              *string            `json:"unit,omitempty"`
}

var _ json.Unmarshaler = &SubRequest{}

func (s *SubRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Message           *string            `json:"message,omitempty"`
		Name              *ResourceName      `json:"name,omitempty"`
		ProvisioningState *QuotaRequestState `json:"provisioningState,omitempty"`
		ResourceType      *string            `json:"resourceType,omitempty"`
		SubRequestId      *string            `json:"subRequestId,omitempty"`
		Unit              *string            `json:"unit,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Message = decoded.Message
	s.Name = decoded.Name
	s.ProvisioningState = decoded.ProvisioningState
	s.ResourceType = decoded.ResourceType
	s.SubRequestId = decoded.SubRequestId
	s.Unit = decoded.Unit

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SubRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["limit"]; ok {
		impl, err := UnmarshalLimitJsonObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Limit' for 'SubRequest': %+v", err)
		}
		s.Limit = impl
	}

	return nil
}
