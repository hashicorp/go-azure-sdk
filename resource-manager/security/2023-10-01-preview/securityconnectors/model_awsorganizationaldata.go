package securityconnectors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsOrganizationalData interface {
	AwsOrganizationalData() BaseAwsOrganizationalDataImpl
}

var _ AwsOrganizationalData = BaseAwsOrganizationalDataImpl{}

type BaseAwsOrganizationalDataImpl struct {
	OrganizationMembershipType OrganizationMembershipType `json:"organizationMembershipType"`
}

func (s BaseAwsOrganizationalDataImpl) AwsOrganizationalData() BaseAwsOrganizationalDataImpl {
	return s
}

var _ AwsOrganizationalData = RawAwsOrganizationalDataImpl{}

// RawAwsOrganizationalDataImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAwsOrganizationalDataImpl struct {
	awsOrganizationalData BaseAwsOrganizationalDataImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawAwsOrganizationalDataImpl) AwsOrganizationalData() BaseAwsOrganizationalDataImpl {
	return s.awsOrganizationalData
}

func UnmarshalAwsOrganizationalDataImplementation(input []byte) (AwsOrganizationalData, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsOrganizationalData into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["organizationMembershipType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Organization") {
		var out AwsOrganizationalDataMaster
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsOrganizationalDataMaster: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Member") {
		var out AwsOrganizationalDataMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsOrganizationalDataMember: %+v", err)
		}
		return out, nil
	}

	var parent BaseAwsOrganizationalDataImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAwsOrganizationalDataImpl: %+v", err)
	}

	return RawAwsOrganizationalDataImpl{
		awsOrganizationalData: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
