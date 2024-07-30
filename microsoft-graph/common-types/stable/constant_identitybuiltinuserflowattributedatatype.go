package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityBuiltInUserFlowAttributeDataType string

const (
	IdentityBuiltInUserFlowAttributeDataType_Boolean          IdentityBuiltInUserFlowAttributeDataType = "boolean"
	IdentityBuiltInUserFlowAttributeDataType_DateTime         IdentityBuiltInUserFlowAttributeDataType = "dateTime"
	IdentityBuiltInUserFlowAttributeDataType_Int64            IdentityBuiltInUserFlowAttributeDataType = "int64"
	IdentityBuiltInUserFlowAttributeDataType_String           IdentityBuiltInUserFlowAttributeDataType = "string"
	IdentityBuiltInUserFlowAttributeDataType_StringCollection IdentityBuiltInUserFlowAttributeDataType = "stringCollection"
)

func PossibleValuesForIdentityBuiltInUserFlowAttributeDataType() []string {
	return []string{
		string(IdentityBuiltInUserFlowAttributeDataType_Boolean),
		string(IdentityBuiltInUserFlowAttributeDataType_DateTime),
		string(IdentityBuiltInUserFlowAttributeDataType_Int64),
		string(IdentityBuiltInUserFlowAttributeDataType_String),
		string(IdentityBuiltInUserFlowAttributeDataType_StringCollection),
	}
}

func (s *IdentityBuiltInUserFlowAttributeDataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityBuiltInUserFlowAttributeDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityBuiltInUserFlowAttributeDataType(input string) (*IdentityBuiltInUserFlowAttributeDataType, error) {
	vals := map[string]IdentityBuiltInUserFlowAttributeDataType{
		"boolean":          IdentityBuiltInUserFlowAttributeDataType_Boolean,
		"datetime":         IdentityBuiltInUserFlowAttributeDataType_DateTime,
		"int64":            IdentityBuiltInUserFlowAttributeDataType_Int64,
		"string":           IdentityBuiltInUserFlowAttributeDataType_String,
		"stringcollection": IdentityBuiltInUserFlowAttributeDataType_StringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityBuiltInUserFlowAttributeDataType(input)
	return &out, nil
}
