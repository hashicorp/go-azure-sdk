package codeversion

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingUploadCredentialDto interface {
	PendingUploadCredentialDto() BasePendingUploadCredentialDtoImpl
}

var _ PendingUploadCredentialDto = BasePendingUploadCredentialDtoImpl{}

type BasePendingUploadCredentialDtoImpl struct {
	CredentialType PendingUploadCredentialType `json:"credentialType"`
}

func (s BasePendingUploadCredentialDtoImpl) PendingUploadCredentialDto() BasePendingUploadCredentialDtoImpl {
	return s
}

var _ PendingUploadCredentialDto = RawPendingUploadCredentialDtoImpl{}

// RawPendingUploadCredentialDtoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPendingUploadCredentialDtoImpl struct {
	pendingUploadCredentialDto BasePendingUploadCredentialDtoImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawPendingUploadCredentialDtoImpl) PendingUploadCredentialDto() BasePendingUploadCredentialDtoImpl {
	return s.pendingUploadCredentialDto
}

func UnmarshalPendingUploadCredentialDtoImplementation(input []byte) (PendingUploadCredentialDto, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PendingUploadCredentialDto into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["credentialType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "SAS") {
		var out SASCredentialDto
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SASCredentialDto: %+v", err)
		}
		return out, nil
	}

	var parent BasePendingUploadCredentialDtoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePendingUploadCredentialDtoImpl: %+v", err)
	}

	return RawPendingUploadCredentialDtoImpl{
		pendingUploadCredentialDto: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
