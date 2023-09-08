package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMarkUserAsCompromisedResponseActionIdentifier string

const (
	SecurityMarkUserAsCompromisedResponseActionIdentifieraccountObjectId                  SecurityMarkUserAsCompromisedResponseActionIdentifier = "AccountObjectId"
	SecurityMarkUserAsCompromisedResponseActionIdentifierinitiatingProcessAccountObjectId SecurityMarkUserAsCompromisedResponseActionIdentifier = "InitiatingProcessAccountObjectId"
	SecurityMarkUserAsCompromisedResponseActionIdentifierrecipientObjectId                SecurityMarkUserAsCompromisedResponseActionIdentifier = "RecipientObjectId"
	SecurityMarkUserAsCompromisedResponseActionIdentifierservicePrincipalId               SecurityMarkUserAsCompromisedResponseActionIdentifier = "ServicePrincipalId"
)

func PossibleValuesForSecurityMarkUserAsCompromisedResponseActionIdentifier() []string {
	return []string{
		string(SecurityMarkUserAsCompromisedResponseActionIdentifieraccountObjectId),
		string(SecurityMarkUserAsCompromisedResponseActionIdentifierinitiatingProcessAccountObjectId),
		string(SecurityMarkUserAsCompromisedResponseActionIdentifierrecipientObjectId),
		string(SecurityMarkUserAsCompromisedResponseActionIdentifierservicePrincipalId),
	}
}

func parseSecurityMarkUserAsCompromisedResponseActionIdentifier(input string) (*SecurityMarkUserAsCompromisedResponseActionIdentifier, error) {
	vals := map[string]SecurityMarkUserAsCompromisedResponseActionIdentifier{
		"accountobjectid":                  SecurityMarkUserAsCompromisedResponseActionIdentifieraccountObjectId,
		"initiatingprocessaccountobjectid": SecurityMarkUserAsCompromisedResponseActionIdentifierinitiatingProcessAccountObjectId,
		"recipientobjectid":                SecurityMarkUserAsCompromisedResponseActionIdentifierrecipientObjectId,
		"serviceprincipalid":               SecurityMarkUserAsCompromisedResponseActionIdentifierservicePrincipalId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMarkUserAsCompromisedResponseActionIdentifier(input)
	return &out, nil
}
