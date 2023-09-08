package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDisableUserResponseActionIdentifier string

const (
	SecurityDisableUserResponseActionIdentifieraccountSid                  SecurityDisableUserResponseActionIdentifier = "AccountSid"
	SecurityDisableUserResponseActionIdentifierinitiatingProcessAccountSid SecurityDisableUserResponseActionIdentifier = "InitiatingProcessAccountSid"
	SecurityDisableUserResponseActionIdentifieronPremSid                   SecurityDisableUserResponseActionIdentifier = "OnPremSid"
	SecurityDisableUserResponseActionIdentifierrequestAccountSid           SecurityDisableUserResponseActionIdentifier = "RequestAccountSid"
)

func PossibleValuesForSecurityDisableUserResponseActionIdentifier() []string {
	return []string{
		string(SecurityDisableUserResponseActionIdentifieraccountSid),
		string(SecurityDisableUserResponseActionIdentifierinitiatingProcessAccountSid),
		string(SecurityDisableUserResponseActionIdentifieronPremSid),
		string(SecurityDisableUserResponseActionIdentifierrequestAccountSid),
	}
}

func parseSecurityDisableUserResponseActionIdentifier(input string) (*SecurityDisableUserResponseActionIdentifier, error) {
	vals := map[string]SecurityDisableUserResponseActionIdentifier{
		"accountsid":                  SecurityDisableUserResponseActionIdentifieraccountSid,
		"initiatingprocessaccountsid": SecurityDisableUserResponseActionIdentifierinitiatingProcessAccountSid,
		"onpremsid":                   SecurityDisableUserResponseActionIdentifieronPremSid,
		"requestaccountsid":           SecurityDisableUserResponseActionIdentifierrequestAccountSid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDisableUserResponseActionIdentifier(input)
	return &out, nil
}
