package feature

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureDataType string

const (
	FeatureDataTypeBinary   FeatureDataType = "Binary"
	FeatureDataTypeBoolean  FeatureDataType = "Boolean"
	FeatureDataTypeDatetime FeatureDataType = "Datetime"
	FeatureDataTypeDouble   FeatureDataType = "Double"
	FeatureDataTypeFloat    FeatureDataType = "Float"
	FeatureDataTypeInteger  FeatureDataType = "Integer"
	FeatureDataTypeLong     FeatureDataType = "Long"
	FeatureDataTypeString   FeatureDataType = "String"
)

func PossibleValuesForFeatureDataType() []string {
	return []string{
		string(FeatureDataTypeBinary),
		string(FeatureDataTypeBoolean),
		string(FeatureDataTypeDatetime),
		string(FeatureDataTypeDouble),
		string(FeatureDataTypeFloat),
		string(FeatureDataTypeInteger),
		string(FeatureDataTypeLong),
		string(FeatureDataTypeString),
	}
}

func (s *FeatureDataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureDataType(input string) (*FeatureDataType, error) {
	vals := map[string]FeatureDataType{
		"binary":   FeatureDataTypeBinary,
		"boolean":  FeatureDataTypeBoolean,
		"datetime": FeatureDataTypeDatetime,
		"double":   FeatureDataTypeDouble,
		"float":    FeatureDataTypeFloat,
		"integer":  FeatureDataTypeInteger,
		"long":     FeatureDataTypeLong,
		"string":   FeatureDataTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureDataType(input)
	return &out, nil
}
