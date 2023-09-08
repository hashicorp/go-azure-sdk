package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsSettingValueType string

const (
	ManagedTenantsSettingValueTypeboolean           ManagedTenantsSettingValueType = "Boolean"
	ManagedTenantsSettingValueTypebooleanCollection ManagedTenantsSettingValueType = "BooleanCollection"
	ManagedTenantsSettingValueTypeguid              ManagedTenantsSettingValueType = "Guid"
	ManagedTenantsSettingValueTypeguidCollection    ManagedTenantsSettingValueType = "GuidCollection"
	ManagedTenantsSettingValueTypeinteger           ManagedTenantsSettingValueType = "Integer"
	ManagedTenantsSettingValueTypeintegerCollection ManagedTenantsSettingValueType = "IntegerCollection"
	ManagedTenantsSettingValueTypestring            ManagedTenantsSettingValueType = "String"
	ManagedTenantsSettingValueTypestringCollection  ManagedTenantsSettingValueType = "StringCollection"
)

func PossibleValuesForManagedTenantsSettingValueType() []string {
	return []string{
		string(ManagedTenantsSettingValueTypeboolean),
		string(ManagedTenantsSettingValueTypebooleanCollection),
		string(ManagedTenantsSettingValueTypeguid),
		string(ManagedTenantsSettingValueTypeguidCollection),
		string(ManagedTenantsSettingValueTypeinteger),
		string(ManagedTenantsSettingValueTypeintegerCollection),
		string(ManagedTenantsSettingValueTypestring),
		string(ManagedTenantsSettingValueTypestringCollection),
	}
}

func parseManagedTenantsSettingValueType(input string) (*ManagedTenantsSettingValueType, error) {
	vals := map[string]ManagedTenantsSettingValueType{
		"boolean":           ManagedTenantsSettingValueTypeboolean,
		"booleancollection": ManagedTenantsSettingValueTypebooleanCollection,
		"guid":              ManagedTenantsSettingValueTypeguid,
		"guidcollection":    ManagedTenantsSettingValueTypeguidCollection,
		"integer":           ManagedTenantsSettingValueTypeinteger,
		"integercollection": ManagedTenantsSettingValueTypeintegerCollection,
		"string":            ManagedTenantsSettingValueTypestring,
		"stringcollection":  ManagedTenantsSettingValueTypestringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsSettingValueType(input)
	return &out, nil
}
