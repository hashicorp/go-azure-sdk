package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttributeDataType string

const (
	IdentityUserFlowAttributeDataTypeboolean          IdentityUserFlowAttributeDataType = "Boolean"
	IdentityUserFlowAttributeDataTypedateTime         IdentityUserFlowAttributeDataType = "DateTime"
	IdentityUserFlowAttributeDataTypeint64            IdentityUserFlowAttributeDataType = "Int64"
	IdentityUserFlowAttributeDataTypestring           IdentityUserFlowAttributeDataType = "String"
	IdentityUserFlowAttributeDataTypestringCollection IdentityUserFlowAttributeDataType = "StringCollection"
)

func PossibleValuesForIdentityUserFlowAttributeDataType() []string {
	return []string{
		string(IdentityUserFlowAttributeDataTypeboolean),
		string(IdentityUserFlowAttributeDataTypedateTime),
		string(IdentityUserFlowAttributeDataTypeint64),
		string(IdentityUserFlowAttributeDataTypestring),
		string(IdentityUserFlowAttributeDataTypestringCollection),
	}
}

func parseIdentityUserFlowAttributeDataType(input string) (*IdentityUserFlowAttributeDataType, error) {
	vals := map[string]IdentityUserFlowAttributeDataType{
		"boolean":          IdentityUserFlowAttributeDataTypeboolean,
		"datetime":         IdentityUserFlowAttributeDataTypedateTime,
		"int64":            IdentityUserFlowAttributeDataTypeint64,
		"string":           IdentityUserFlowAttributeDataTypestring,
		"stringcollection": IdentityUserFlowAttributeDataTypestringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityUserFlowAttributeDataType(input)
	return &out, nil
}
