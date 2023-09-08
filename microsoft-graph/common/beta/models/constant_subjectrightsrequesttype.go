package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestType string

const (
	SubjectRightsRequestTypeaccess       SubjectRightsRequestType = "Access"
	SubjectRightsRequestTypedelete       SubjectRightsRequestType = "Delete"
	SubjectRightsRequestTypeexport       SubjectRightsRequestType = "Export"
	SubjectRightsRequestTypetagForAction SubjectRightsRequestType = "TagForAction"
)

func PossibleValuesForSubjectRightsRequestType() []string {
	return []string{
		string(SubjectRightsRequestTypeaccess),
		string(SubjectRightsRequestTypedelete),
		string(SubjectRightsRequestTypeexport),
		string(SubjectRightsRequestTypetagForAction),
	}
}

func parseSubjectRightsRequestType(input string) (*SubjectRightsRequestType, error) {
	vals := map[string]SubjectRightsRequestType{
		"access":       SubjectRightsRequestTypeaccess,
		"delete":       SubjectRightsRequestTypedelete,
		"export":       SubjectRightsRequestTypeexport,
		"tagforaction": SubjectRightsRequestTypetagForAction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestType(input)
	return &out, nil
}
