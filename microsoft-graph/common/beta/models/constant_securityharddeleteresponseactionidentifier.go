package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHardDeleteResponseActionIdentifier string

const (
	SecurityHardDeleteResponseActionIdentifiernetworkMessageId      SecurityHardDeleteResponseActionIdentifier = "NetworkMessageId"
	SecurityHardDeleteResponseActionIdentifierrecipientEmailAddress SecurityHardDeleteResponseActionIdentifier = "RecipientEmailAddress"
)

func PossibleValuesForSecurityHardDeleteResponseActionIdentifier() []string {
	return []string{
		string(SecurityHardDeleteResponseActionIdentifiernetworkMessageId),
		string(SecurityHardDeleteResponseActionIdentifierrecipientEmailAddress),
	}
}

func parseSecurityHardDeleteResponseActionIdentifier(input string) (*SecurityHardDeleteResponseActionIdentifier, error) {
	vals := map[string]SecurityHardDeleteResponseActionIdentifier{
		"networkmessageid":      SecurityHardDeleteResponseActionIdentifiernetworkMessageId,
		"recipientemailaddress": SecurityHardDeleteResponseActionIdentifierrecipientEmailAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHardDeleteResponseActionIdentifier(input)
	return &out, nil
}
