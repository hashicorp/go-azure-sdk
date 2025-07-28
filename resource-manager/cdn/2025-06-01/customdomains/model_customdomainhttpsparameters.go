package customdomains

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomDomainHTTPSParameters interface {
	CustomDomainHTTPSParameters() BaseCustomDomainHTTPSParametersImpl
}

var _ CustomDomainHTTPSParameters = BaseCustomDomainHTTPSParametersImpl{}

type BaseCustomDomainHTTPSParametersImpl struct {
	CertificateSource CertificateSource  `json:"certificateSource"`
	MinimumTlsVersion *MinimumTlsVersion `json:"minimumTlsVersion,omitempty"`
	ProtocolType      ProtocolType       `json:"protocolType"`
}

func (s BaseCustomDomainHTTPSParametersImpl) CustomDomainHTTPSParameters() BaseCustomDomainHTTPSParametersImpl {
	return s
}

var _ CustomDomainHTTPSParameters = RawCustomDomainHTTPSParametersImpl{}

// RawCustomDomainHTTPSParametersImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomDomainHTTPSParametersImpl struct {
	customDomainHTTPSParameters BaseCustomDomainHTTPSParametersImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawCustomDomainHTTPSParametersImpl) CustomDomainHTTPSParameters() BaseCustomDomainHTTPSParametersImpl {
	return s.customDomainHTTPSParameters
}

func UnmarshalCustomDomainHTTPSParametersImplementation(input []byte) (CustomDomainHTTPSParameters, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomDomainHTTPSParameters into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["certificateSource"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Cdn") {
		var out CdnManagedHTTPSParameters
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CdnManagedHTTPSParameters: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureKeyVault") {
		var out UserManagedHTTPSParameters
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserManagedHTTPSParameters: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomDomainHTTPSParametersImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomDomainHTTPSParametersImpl: %+v", err)
	}

	return RawCustomDomainHTTPSParametersImpl{
		customDomainHTTPSParameters: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
