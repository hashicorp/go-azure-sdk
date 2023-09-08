package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityForceUserPasswordResetResponseActionIdentifier string

const (
	SecurityForceUserPasswordResetResponseActionIdentifieraccountSid                  SecurityForceUserPasswordResetResponseActionIdentifier = "AccountSid"
	SecurityForceUserPasswordResetResponseActionIdentifierinitiatingProcessAccountSid SecurityForceUserPasswordResetResponseActionIdentifier = "InitiatingProcessAccountSid"
	SecurityForceUserPasswordResetResponseActionIdentifieronPremSid                   SecurityForceUserPasswordResetResponseActionIdentifier = "OnPremSid"
	SecurityForceUserPasswordResetResponseActionIdentifierrequestAccountSid           SecurityForceUserPasswordResetResponseActionIdentifier = "RequestAccountSid"
)

func PossibleValuesForSecurityForceUserPasswordResetResponseActionIdentifier() []string {
	return []string{
		string(SecurityForceUserPasswordResetResponseActionIdentifieraccountSid),
		string(SecurityForceUserPasswordResetResponseActionIdentifierinitiatingProcessAccountSid),
		string(SecurityForceUserPasswordResetResponseActionIdentifieronPremSid),
		string(SecurityForceUserPasswordResetResponseActionIdentifierrequestAccountSid),
	}
}

func parseSecurityForceUserPasswordResetResponseActionIdentifier(input string) (*SecurityForceUserPasswordResetResponseActionIdentifier, error) {
	vals := map[string]SecurityForceUserPasswordResetResponseActionIdentifier{
		"accountsid":                  SecurityForceUserPasswordResetResponseActionIdentifieraccountSid,
		"initiatingprocessaccountsid": SecurityForceUserPasswordResetResponseActionIdentifierinitiatingProcessAccountSid,
		"onpremsid":                   SecurityForceUserPasswordResetResponseActionIdentifieronPremSid,
		"requestaccountsid":           SecurityForceUserPasswordResetResponseActionIdentifierrequestAccountSid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityForceUserPasswordResetResponseActionIdentifier(input)
	return &out, nil
}
