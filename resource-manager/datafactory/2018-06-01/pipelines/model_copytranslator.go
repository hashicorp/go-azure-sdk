package pipelines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyTranslator interface {
	CopyTranslator() BaseCopyTranslatorImpl
}

var _ CopyTranslator = BaseCopyTranslatorImpl{}

type BaseCopyTranslatorImpl struct {
	Type string `json:"type"`
}

func (s BaseCopyTranslatorImpl) CopyTranslator() BaseCopyTranslatorImpl {
	return s
}

var _ CopyTranslator = RawCopyTranslatorImpl{}

// RawCopyTranslatorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCopyTranslatorImpl struct {
	copyTranslator BaseCopyTranslatorImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawCopyTranslatorImpl) CopyTranslator() BaseCopyTranslatorImpl {
	return s.copyTranslator
}

func UnmarshalCopyTranslatorImplementation(input []byte) (CopyTranslator, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CopyTranslator into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "TabularTranslator") {
		var out TabularTranslator
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TabularTranslator: %+v", err)
		}
		return out, nil
	}

	var parent BaseCopyTranslatorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCopyTranslatorImpl: %+v", err)
	}

	return RawCopyTranslatorImpl{
		copyTranslator: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
