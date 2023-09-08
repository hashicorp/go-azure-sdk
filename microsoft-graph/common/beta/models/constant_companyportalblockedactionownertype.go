package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyPortalBlockedActionOwnerType string

const (
	CompanyPortalBlockedActionOwnerTypecompany  CompanyPortalBlockedActionOwnerType = "Company"
	CompanyPortalBlockedActionOwnerTypepersonal CompanyPortalBlockedActionOwnerType = "Personal"
	CompanyPortalBlockedActionOwnerTypeunknown  CompanyPortalBlockedActionOwnerType = "Unknown"
)

func PossibleValuesForCompanyPortalBlockedActionOwnerType() []string {
	return []string{
		string(CompanyPortalBlockedActionOwnerTypecompany),
		string(CompanyPortalBlockedActionOwnerTypepersonal),
		string(CompanyPortalBlockedActionOwnerTypeunknown),
	}
}

func parseCompanyPortalBlockedActionOwnerType(input string) (*CompanyPortalBlockedActionOwnerType, error) {
	vals := map[string]CompanyPortalBlockedActionOwnerType{
		"company":  CompanyPortalBlockedActionOwnerTypecompany,
		"personal": CompanyPortalBlockedActionOwnerTypepersonal,
		"unknown":  CompanyPortalBlockedActionOwnerTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompanyPortalBlockedActionOwnerType(input)
	return &out, nil
}
