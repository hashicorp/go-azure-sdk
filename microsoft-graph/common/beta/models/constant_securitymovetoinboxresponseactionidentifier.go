package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMoveToInboxResponseActionIdentifier string

const (
	SecurityMoveToInboxResponseActionIdentifiernetworkMessageId      SecurityMoveToInboxResponseActionIdentifier = "NetworkMessageId"
	SecurityMoveToInboxResponseActionIdentifierrecipientEmailAddress SecurityMoveToInboxResponseActionIdentifier = "RecipientEmailAddress"
)

func PossibleValuesForSecurityMoveToInboxResponseActionIdentifier() []string {
	return []string{
		string(SecurityMoveToInboxResponseActionIdentifiernetworkMessageId),
		string(SecurityMoveToInboxResponseActionIdentifierrecipientEmailAddress),
	}
}

func parseSecurityMoveToInboxResponseActionIdentifier(input string) (*SecurityMoveToInboxResponseActionIdentifier, error) {
	vals := map[string]SecurityMoveToInboxResponseActionIdentifier{
		"networkmessageid":      SecurityMoveToInboxResponseActionIdentifiernetworkMessageId,
		"recipientemailaddress": SecurityMoveToInboxResponseActionIdentifierrecipientEmailAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMoveToInboxResponseActionIdentifier(input)
	return &out, nil
}
