package itemlevelrecoveryconnections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ILRRequest interface {
	ILRRequest() BaseILRRequestImpl
}

var _ ILRRequest = BaseILRRequestImpl{}

type BaseILRRequestImpl struct {
	ObjectType string `json:"objectType"`
}

func (s BaseILRRequestImpl) ILRRequest() BaseILRRequestImpl {
	return s
}

var _ ILRRequest = RawILRRequestImpl{}

// RawILRRequestImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawILRRequestImpl struct {
	iLRRequest BaseILRRequestImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawILRRequestImpl) ILRRequest() BaseILRRequestImpl {
	return s.iLRRequest
}

func UnmarshalILRRequestImplementation(input []byte) (ILRRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ILRRequest into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["objectType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureFileShareProvisionILRRequest") {
		var out AzureFileShareProvisionILRRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureFileShareProvisionILRRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IaasVMILRRegistrationRequest") {
		var out IaasVMILRRegistrationRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IaasVMILRRegistrationRequest: %+v", err)
		}
		return out, nil
	}

	var parent BaseILRRequestImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseILRRequestImpl: %+v", err)
	}

	return RawILRRequestImpl{
		iLRRequest: parent,
		Type:       value,
		Values:     temp,
	}, nil

}
