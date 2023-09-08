package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsPropertyType string

const (
	ExternalConnectorsPropertyTypeboolean            ExternalConnectorsPropertyType = "Boolean"
	ExternalConnectorsPropertyTypedateTime           ExternalConnectorsPropertyType = "DateTime"
	ExternalConnectorsPropertyTypedateTimeCollection ExternalConnectorsPropertyType = "DateTimeCollection"
	ExternalConnectorsPropertyTypedouble             ExternalConnectorsPropertyType = "Double"
	ExternalConnectorsPropertyTypedoubleCollection   ExternalConnectorsPropertyType = "DoubleCollection"
	ExternalConnectorsPropertyTypeint64              ExternalConnectorsPropertyType = "Int64"
	ExternalConnectorsPropertyTypeint64Collection    ExternalConnectorsPropertyType = "Int64Collection"
	ExternalConnectorsPropertyTypestring             ExternalConnectorsPropertyType = "String"
	ExternalConnectorsPropertyTypestringCollection   ExternalConnectorsPropertyType = "StringCollection"
)

func PossibleValuesForExternalConnectorsPropertyType() []string {
	return []string{
		string(ExternalConnectorsPropertyTypeboolean),
		string(ExternalConnectorsPropertyTypedateTime),
		string(ExternalConnectorsPropertyTypedateTimeCollection),
		string(ExternalConnectorsPropertyTypedouble),
		string(ExternalConnectorsPropertyTypedoubleCollection),
		string(ExternalConnectorsPropertyTypeint64),
		string(ExternalConnectorsPropertyTypeint64Collection),
		string(ExternalConnectorsPropertyTypestring),
		string(ExternalConnectorsPropertyTypestringCollection),
	}
}

func parseExternalConnectorsPropertyType(input string) (*ExternalConnectorsPropertyType, error) {
	vals := map[string]ExternalConnectorsPropertyType{
		"boolean":            ExternalConnectorsPropertyTypeboolean,
		"datetime":           ExternalConnectorsPropertyTypedateTime,
		"datetimecollection": ExternalConnectorsPropertyTypedateTimeCollection,
		"double":             ExternalConnectorsPropertyTypedouble,
		"doublecollection":   ExternalConnectorsPropertyTypedoubleCollection,
		"int64":              ExternalConnectorsPropertyTypeint64,
		"int64collection":    ExternalConnectorsPropertyTypeint64Collection,
		"string":             ExternalConnectorsPropertyTypestring,
		"stringcollection":   ExternalConnectorsPropertyTypestringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsPropertyType(input)
	return &out, nil
}
