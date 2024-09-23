package contentkeypolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentKeyPolicyPlayReadyContentKeyLocation interface {
	ContentKeyPolicyPlayReadyContentKeyLocation() BaseContentKeyPolicyPlayReadyContentKeyLocationImpl
}

var _ ContentKeyPolicyPlayReadyContentKeyLocation = BaseContentKeyPolicyPlayReadyContentKeyLocationImpl{}

type BaseContentKeyPolicyPlayReadyContentKeyLocationImpl struct {
	OdataType string `json:"@odata.type"`
}

func (s BaseContentKeyPolicyPlayReadyContentKeyLocationImpl) ContentKeyPolicyPlayReadyContentKeyLocation() BaseContentKeyPolicyPlayReadyContentKeyLocationImpl {
	return s
}

var _ ContentKeyPolicyPlayReadyContentKeyLocation = RawContentKeyPolicyPlayReadyContentKeyLocationImpl{}

// RawContentKeyPolicyPlayReadyContentKeyLocationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawContentKeyPolicyPlayReadyContentKeyLocationImpl struct {
	contentKeyPolicyPlayReadyContentKeyLocation BaseContentKeyPolicyPlayReadyContentKeyLocationImpl
	Type                                        string
	Values                                      map[string]interface{}
}

func (s RawContentKeyPolicyPlayReadyContentKeyLocationImpl) ContentKeyPolicyPlayReadyContentKeyLocation() BaseContentKeyPolicyPlayReadyContentKeyLocationImpl {
	return s.contentKeyPolicyPlayReadyContentKeyLocation
}

func UnmarshalContentKeyPolicyPlayReadyContentKeyLocationImplementation(input []byte) (ContentKeyPolicyPlayReadyContentKeyLocation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ContentKeyPolicyPlayReadyContentKeyLocation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader") {
		var out ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier") {
		var out ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier: %+v", err)
		}
		return out, nil
	}

	var parent BaseContentKeyPolicyPlayReadyContentKeyLocationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseContentKeyPolicyPlayReadyContentKeyLocationImpl: %+v", err)
	}

	return RawContentKeyPolicyPlayReadyContentKeyLocationImpl{
		contentKeyPolicyPlayReadyContentKeyLocation: parent,
		Type:   value,
		Values: temp,
	}, nil

}
