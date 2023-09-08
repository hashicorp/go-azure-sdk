package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyPortalBlockedActionAction string

const (
	CompanyPortalBlockedActionActionremove  CompanyPortalBlockedActionAction = "Remove"
	CompanyPortalBlockedActionActionreset   CompanyPortalBlockedActionAction = "Reset"
	CompanyPortalBlockedActionActionunknown CompanyPortalBlockedActionAction = "Unknown"
)

func PossibleValuesForCompanyPortalBlockedActionAction() []string {
	return []string{
		string(CompanyPortalBlockedActionActionremove),
		string(CompanyPortalBlockedActionActionreset),
		string(CompanyPortalBlockedActionActionunknown),
	}
}

func parseCompanyPortalBlockedActionAction(input string) (*CompanyPortalBlockedActionAction, error) {
	vals := map[string]CompanyPortalBlockedActionAction{
		"remove":  CompanyPortalBlockedActionActionremove,
		"reset":   CompanyPortalBlockedActionActionreset,
		"unknown": CompanyPortalBlockedActionActionunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompanyPortalBlockedActionAction(input)
	return &out, nil
}
