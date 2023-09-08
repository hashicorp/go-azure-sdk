package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ColumnDefinitionType string

const (
	ColumnDefinitionTypeapprovalStatus ColumnDefinitionType = "ApprovalStatus"
	ColumnDefinitionTypeboolean        ColumnDefinitionType = "Boolean"
	ColumnDefinitionTypecalculated     ColumnDefinitionType = "Calculated"
	ColumnDefinitionTypechoice         ColumnDefinitionType = "Choice"
	ColumnDefinitionTypecurrency       ColumnDefinitionType = "Currency"
	ColumnDefinitionTypedateTime       ColumnDefinitionType = "DateTime"
	ColumnDefinitionTypegeolocation    ColumnDefinitionType = "Geolocation"
	ColumnDefinitionTypelocation       ColumnDefinitionType = "Location"
	ColumnDefinitionTypelookup         ColumnDefinitionType = "Lookup"
	ColumnDefinitionTypemultichoice    ColumnDefinitionType = "Multichoice"
	ColumnDefinitionTypemultiterm      ColumnDefinitionType = "Multiterm"
	ColumnDefinitionTypenote           ColumnDefinitionType = "Note"
	ColumnDefinitionTypenumber         ColumnDefinitionType = "Number"
	ColumnDefinitionTypeterm           ColumnDefinitionType = "Term"
	ColumnDefinitionTypetext           ColumnDefinitionType = "Text"
	ColumnDefinitionTypethumbnail      ColumnDefinitionType = "Thumbnail"
	ColumnDefinitionTypeurl            ColumnDefinitionType = "Url"
	ColumnDefinitionTypeuser           ColumnDefinitionType = "User"
)

func PossibleValuesForColumnDefinitionType() []string {
	return []string{
		string(ColumnDefinitionTypeapprovalStatus),
		string(ColumnDefinitionTypeboolean),
		string(ColumnDefinitionTypecalculated),
		string(ColumnDefinitionTypechoice),
		string(ColumnDefinitionTypecurrency),
		string(ColumnDefinitionTypedateTime),
		string(ColumnDefinitionTypegeolocation),
		string(ColumnDefinitionTypelocation),
		string(ColumnDefinitionTypelookup),
		string(ColumnDefinitionTypemultichoice),
		string(ColumnDefinitionTypemultiterm),
		string(ColumnDefinitionTypenote),
		string(ColumnDefinitionTypenumber),
		string(ColumnDefinitionTypeterm),
		string(ColumnDefinitionTypetext),
		string(ColumnDefinitionTypethumbnail),
		string(ColumnDefinitionTypeurl),
		string(ColumnDefinitionTypeuser),
	}
}

func parseColumnDefinitionType(input string) (*ColumnDefinitionType, error) {
	vals := map[string]ColumnDefinitionType{
		"approvalstatus": ColumnDefinitionTypeapprovalStatus,
		"boolean":        ColumnDefinitionTypeboolean,
		"calculated":     ColumnDefinitionTypecalculated,
		"choice":         ColumnDefinitionTypechoice,
		"currency":       ColumnDefinitionTypecurrency,
		"datetime":       ColumnDefinitionTypedateTime,
		"geolocation":    ColumnDefinitionTypegeolocation,
		"location":       ColumnDefinitionTypelocation,
		"lookup":         ColumnDefinitionTypelookup,
		"multichoice":    ColumnDefinitionTypemultichoice,
		"multiterm":      ColumnDefinitionTypemultiterm,
		"note":           ColumnDefinitionTypenote,
		"number":         ColumnDefinitionTypenumber,
		"term":           ColumnDefinitionTypeterm,
		"text":           ColumnDefinitionTypetext,
		"thumbnail":      ColumnDefinitionTypethumbnail,
		"url":            ColumnDefinitionTypeurl,
		"user":           ColumnDefinitionTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ColumnDefinitionType(input)
	return &out, nil
}
