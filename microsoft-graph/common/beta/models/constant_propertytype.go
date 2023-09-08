package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PropertyType string

const (
	PropertyTypeboolean            PropertyType = "Boolean"
	PropertyTypedateTime           PropertyType = "DateTime"
	PropertyTypedateTimeCollection PropertyType = "DateTimeCollection"
	PropertyTypedouble             PropertyType = "Double"
	PropertyTypedoubleCollection   PropertyType = "DoubleCollection"
	PropertyTypeint64              PropertyType = "Int64"
	PropertyTypeint64Collection    PropertyType = "Int64Collection"
	PropertyTypestring             PropertyType = "String"
	PropertyTypestringCollection   PropertyType = "StringCollection"
)

func PossibleValuesForPropertyType() []string {
	return []string{
		string(PropertyTypeboolean),
		string(PropertyTypedateTime),
		string(PropertyTypedateTimeCollection),
		string(PropertyTypedouble),
		string(PropertyTypedoubleCollection),
		string(PropertyTypeint64),
		string(PropertyTypeint64Collection),
		string(PropertyTypestring),
		string(PropertyTypestringCollection),
	}
}

func parsePropertyType(input string) (*PropertyType, error) {
	vals := map[string]PropertyType{
		"boolean":            PropertyTypeboolean,
		"datetime":           PropertyTypedateTime,
		"datetimecollection": PropertyTypedateTimeCollection,
		"double":             PropertyTypedouble,
		"doublecollection":   PropertyTypedoubleCollection,
		"int64":              PropertyTypeint64,
		"int64collection":    PropertyTypeint64Collection,
		"string":             PropertyTypestring,
		"stringcollection":   PropertyTypestringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PropertyType(input)
	return &out, nil
}
