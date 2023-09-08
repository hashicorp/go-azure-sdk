package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityCustomUserFlowAttributeDataType string

const (
	IdentityCustomUserFlowAttributeDataTypeboolean          IdentityCustomUserFlowAttributeDataType = "Boolean"
	IdentityCustomUserFlowAttributeDataTypedateTime         IdentityCustomUserFlowAttributeDataType = "DateTime"
	IdentityCustomUserFlowAttributeDataTypeint64            IdentityCustomUserFlowAttributeDataType = "Int64"
	IdentityCustomUserFlowAttributeDataTypestring           IdentityCustomUserFlowAttributeDataType = "String"
	IdentityCustomUserFlowAttributeDataTypestringCollection IdentityCustomUserFlowAttributeDataType = "StringCollection"
)

func PossibleValuesForIdentityCustomUserFlowAttributeDataType() []string {
	return []string{
		string(IdentityCustomUserFlowAttributeDataTypeboolean),
		string(IdentityCustomUserFlowAttributeDataTypedateTime),
		string(IdentityCustomUserFlowAttributeDataTypeint64),
		string(IdentityCustomUserFlowAttributeDataTypestring),
		string(IdentityCustomUserFlowAttributeDataTypestringCollection),
	}
}

func parseIdentityCustomUserFlowAttributeDataType(input string) (*IdentityCustomUserFlowAttributeDataType, error) {
	vals := map[string]IdentityCustomUserFlowAttributeDataType{
		"boolean":          IdentityCustomUserFlowAttributeDataTypeboolean,
		"datetime":         IdentityCustomUserFlowAttributeDataTypedateTime,
		"int64":            IdentityCustomUserFlowAttributeDataTypeint64,
		"string":           IdentityCustomUserFlowAttributeDataTypestring,
		"stringcollection": IdentityCustomUserFlowAttributeDataTypestringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityCustomUserFlowAttributeDataType(input)
	return &out, nil
}
