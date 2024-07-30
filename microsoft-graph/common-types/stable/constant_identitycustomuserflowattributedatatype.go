package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityCustomUserFlowAttributeDataType string

const (
	IdentityCustomUserFlowAttributeDataType_Boolean          IdentityCustomUserFlowAttributeDataType = "boolean"
	IdentityCustomUserFlowAttributeDataType_DateTime         IdentityCustomUserFlowAttributeDataType = "dateTime"
	IdentityCustomUserFlowAttributeDataType_Int64            IdentityCustomUserFlowAttributeDataType = "int64"
	IdentityCustomUserFlowAttributeDataType_String           IdentityCustomUserFlowAttributeDataType = "string"
	IdentityCustomUserFlowAttributeDataType_StringCollection IdentityCustomUserFlowAttributeDataType = "stringCollection"
)

func PossibleValuesForIdentityCustomUserFlowAttributeDataType() []string {
	return []string{
		string(IdentityCustomUserFlowAttributeDataType_Boolean),
		string(IdentityCustomUserFlowAttributeDataType_DateTime),
		string(IdentityCustomUserFlowAttributeDataType_Int64),
		string(IdentityCustomUserFlowAttributeDataType_String),
		string(IdentityCustomUserFlowAttributeDataType_StringCollection),
	}
}

func (s *IdentityCustomUserFlowAttributeDataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityCustomUserFlowAttributeDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityCustomUserFlowAttributeDataType(input string) (*IdentityCustomUserFlowAttributeDataType, error) {
	vals := map[string]IdentityCustomUserFlowAttributeDataType{
		"boolean":          IdentityCustomUserFlowAttributeDataType_Boolean,
		"datetime":         IdentityCustomUserFlowAttributeDataType_DateTime,
		"int64":            IdentityCustomUserFlowAttributeDataType_Int64,
		"string":           IdentityCustomUserFlowAttributeDataType_String,
		"stringcollection": IdentityCustomUserFlowAttributeDataType_StringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityCustomUserFlowAttributeDataType(input)
	return &out, nil
}
