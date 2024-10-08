package customdomains

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateSourceParameters interface {
	CertificateSourceParameters() BaseCertificateSourceParametersImpl
}

var _ CertificateSourceParameters = BaseCertificateSourceParametersImpl{}

type BaseCertificateSourceParametersImpl struct {
	TypeName CertificateSourceParametersType `json:"typeName"`
}

func (s BaseCertificateSourceParametersImpl) CertificateSourceParameters() BaseCertificateSourceParametersImpl {
	return s
}

var _ CertificateSourceParameters = RawCertificateSourceParametersImpl{}

// RawCertificateSourceParametersImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCertificateSourceParametersImpl struct {
	certificateSourceParameters BaseCertificateSourceParametersImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawCertificateSourceParametersImpl) CertificateSourceParameters() BaseCertificateSourceParametersImpl {
	return s.certificateSourceParameters
}

func UnmarshalCertificateSourceParametersImplementation(input []byte) (CertificateSourceParameters, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CertificateSourceParameters into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["typeName"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "CdnCertificateSourceParameters") {
		var out CdnCertificateSourceParameters
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CdnCertificateSourceParameters: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "KeyVaultCertificateSourceParameters") {
		var out KeyVaultCertificateSourceParameters
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KeyVaultCertificateSourceParameters: %+v", err)
		}
		return out, nil
	}

	var parent BaseCertificateSourceParametersImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCertificateSourceParametersImpl: %+v", err)
	}

	return RawCertificateSourceParametersImpl{
		certificateSourceParameters: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
