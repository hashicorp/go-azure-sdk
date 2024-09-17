package videos

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenKey interface {
	TokenKey() BaseTokenKeyImpl
}

var _ TokenKey = BaseTokenKeyImpl{}

type BaseTokenKeyImpl struct {
	Kid  string `json:"kid"`
	Type string `json:"@type"`
}

func (s BaseTokenKeyImpl) TokenKey() BaseTokenKeyImpl {
	return s
}

var _ TokenKey = RawTokenKeyImpl{}

// RawTokenKeyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTokenKeyImpl struct {
	tokenKey BaseTokenKeyImpl
	Type     string
	Values   map[string]interface{}
}

func (s RawTokenKeyImpl) TokenKey() BaseTokenKeyImpl {
	return s.tokenKey
}

func UnmarshalTokenKeyImplementation(input []byte) (TokenKey, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TokenKey into map[string]interface: %+v", err)
	}

	value, ok := temp["@type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#Microsoft.VideoAnalyzer.EccTokenKey") {
		var out EccTokenKey
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EccTokenKey: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.VideoAnalyzer.RsaTokenKey") {
		var out RsaTokenKey
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RsaTokenKey: %+v", err)
		}
		return out, nil
	}

	var parent BaseTokenKeyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTokenKeyImpl: %+v", err)
	}

	return RawTokenKeyImpl{
		tokenKey: parent,
		Type:     value,
		Values:   temp,
	}, nil

}
