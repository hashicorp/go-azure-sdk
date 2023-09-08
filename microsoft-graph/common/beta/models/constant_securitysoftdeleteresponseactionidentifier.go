package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySoftDeleteResponseActionIdentifier string

const (
	SecuritySoftDeleteResponseActionIdentifiernetworkMessageId      SecuritySoftDeleteResponseActionIdentifier = "NetworkMessageId"
	SecuritySoftDeleteResponseActionIdentifierrecipientEmailAddress SecuritySoftDeleteResponseActionIdentifier = "RecipientEmailAddress"
)

func PossibleValuesForSecuritySoftDeleteResponseActionIdentifier() []string {
	return []string{
		string(SecuritySoftDeleteResponseActionIdentifiernetworkMessageId),
		string(SecuritySoftDeleteResponseActionIdentifierrecipientEmailAddress),
	}
}

func parseSecuritySoftDeleteResponseActionIdentifier(input string) (*SecuritySoftDeleteResponseActionIdentifier, error) {
	vals := map[string]SecuritySoftDeleteResponseActionIdentifier{
		"networkmessageid":      SecuritySoftDeleteResponseActionIdentifiernetworkMessageId,
		"recipientemailaddress": SecuritySoftDeleteResponseActionIdentifierrecipientEmailAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySoftDeleteResponseActionIdentifier(input)
	return &out, nil
}
