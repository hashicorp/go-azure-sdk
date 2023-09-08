package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityBuiltInUserFlowAttributeDataType string

const (
	IdentityBuiltInUserFlowAttributeDataTypeboolean          IdentityBuiltInUserFlowAttributeDataType = "Boolean"
	IdentityBuiltInUserFlowAttributeDataTypedateTime         IdentityBuiltInUserFlowAttributeDataType = "DateTime"
	IdentityBuiltInUserFlowAttributeDataTypeint64            IdentityBuiltInUserFlowAttributeDataType = "Int64"
	IdentityBuiltInUserFlowAttributeDataTypestring           IdentityBuiltInUserFlowAttributeDataType = "String"
	IdentityBuiltInUserFlowAttributeDataTypestringCollection IdentityBuiltInUserFlowAttributeDataType = "StringCollection"
)

func PossibleValuesForIdentityBuiltInUserFlowAttributeDataType() []string {
	return []string{
		string(IdentityBuiltInUserFlowAttributeDataTypeboolean),
		string(IdentityBuiltInUserFlowAttributeDataTypedateTime),
		string(IdentityBuiltInUserFlowAttributeDataTypeint64),
		string(IdentityBuiltInUserFlowAttributeDataTypestring),
		string(IdentityBuiltInUserFlowAttributeDataTypestringCollection),
	}
}

func parseIdentityBuiltInUserFlowAttributeDataType(input string) (*IdentityBuiltInUserFlowAttributeDataType, error) {
	vals := map[string]IdentityBuiltInUserFlowAttributeDataType{
		"boolean":          IdentityBuiltInUserFlowAttributeDataTypeboolean,
		"datetime":         IdentityBuiltInUserFlowAttributeDataTypedateTime,
		"int64":            IdentityBuiltInUserFlowAttributeDataTypeint64,
		"string":           IdentityBuiltInUserFlowAttributeDataTypestring,
		"stringcollection": IdentityBuiltInUserFlowAttributeDataTypestringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityBuiltInUserFlowAttributeDataType(input)
	return &out, nil
}
