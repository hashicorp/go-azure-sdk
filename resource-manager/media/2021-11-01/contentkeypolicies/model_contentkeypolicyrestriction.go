package contentkeypolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentKeyPolicyRestriction interface {
	ContentKeyPolicyRestriction() BaseContentKeyPolicyRestrictionImpl
}

var _ ContentKeyPolicyRestriction = BaseContentKeyPolicyRestrictionImpl{}

type BaseContentKeyPolicyRestrictionImpl struct {
	OdataType string `json:"@odata.type"`
}

func (s BaseContentKeyPolicyRestrictionImpl) ContentKeyPolicyRestriction() BaseContentKeyPolicyRestrictionImpl {
	return s
}

var _ ContentKeyPolicyRestriction = RawContentKeyPolicyRestrictionImpl{}

// RawContentKeyPolicyRestrictionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawContentKeyPolicyRestrictionImpl struct {
	contentKeyPolicyRestriction BaseContentKeyPolicyRestrictionImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawContentKeyPolicyRestrictionImpl) ContentKeyPolicyRestriction() BaseContentKeyPolicyRestrictionImpl {
	return s.contentKeyPolicyRestriction
}

func UnmarshalContentKeyPolicyRestrictionImplementation(input []byte) (ContentKeyPolicyRestriction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ContentKeyPolicyRestriction into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyOpenRestriction") {
		var out ContentKeyPolicyOpenRestriction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyOpenRestriction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyTokenRestriction") {
		var out ContentKeyPolicyTokenRestriction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyTokenRestriction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyUnknownRestriction") {
		var out ContentKeyPolicyUnknownRestriction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyUnknownRestriction: %+v", err)
		}
		return out, nil
	}

	var parent BaseContentKeyPolicyRestrictionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseContentKeyPolicyRestrictionImpl: %+v", err)
	}

	return RawContentKeyPolicyRestrictionImpl{
		contentKeyPolicyRestriction: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
