package contentkeypolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentKeyPolicyRestrictionTokenKey interface {
	ContentKeyPolicyRestrictionTokenKey() BaseContentKeyPolicyRestrictionTokenKeyImpl
}

var _ ContentKeyPolicyRestrictionTokenKey = BaseContentKeyPolicyRestrictionTokenKeyImpl{}

type BaseContentKeyPolicyRestrictionTokenKeyImpl struct {
	OdataType string `json:"@odata.type"`
}

func (s BaseContentKeyPolicyRestrictionTokenKeyImpl) ContentKeyPolicyRestrictionTokenKey() BaseContentKeyPolicyRestrictionTokenKeyImpl {
	return s
}

var _ ContentKeyPolicyRestrictionTokenKey = RawContentKeyPolicyRestrictionTokenKeyImpl{}

// RawContentKeyPolicyRestrictionTokenKeyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawContentKeyPolicyRestrictionTokenKeyImpl struct {
	contentKeyPolicyRestrictionTokenKey BaseContentKeyPolicyRestrictionTokenKeyImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawContentKeyPolicyRestrictionTokenKeyImpl) ContentKeyPolicyRestrictionTokenKey() BaseContentKeyPolicyRestrictionTokenKeyImpl {
	return s.contentKeyPolicyRestrictionTokenKey
}

func UnmarshalContentKeyPolicyRestrictionTokenKeyImplementation(input []byte) (ContentKeyPolicyRestrictionTokenKey, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ContentKeyPolicyRestrictionTokenKey into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyRsaTokenKey") {
		var out ContentKeyPolicyRsaTokenKey
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyRsaTokenKey: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicySymmetricTokenKey") {
		var out ContentKeyPolicySymmetricTokenKey
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicySymmetricTokenKey: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyX509CertificateTokenKey") {
		var out ContentKeyPolicyX509CertificateTokenKey
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyX509CertificateTokenKey: %+v", err)
		}
		return out, nil
	}

	var parent BaseContentKeyPolicyRestrictionTokenKeyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseContentKeyPolicyRestrictionTokenKeyImpl: %+v", err)
	}

	return RawContentKeyPolicyRestrictionTokenKeyImpl{
		contentKeyPolicyRestrictionTokenKey: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
