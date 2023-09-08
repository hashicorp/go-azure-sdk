package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMoveToDeletedItemsResponseActionIdentifier string

const (
	SecurityMoveToDeletedItemsResponseActionIdentifiernetworkMessageId      SecurityMoveToDeletedItemsResponseActionIdentifier = "NetworkMessageId"
	SecurityMoveToDeletedItemsResponseActionIdentifierrecipientEmailAddress SecurityMoveToDeletedItemsResponseActionIdentifier = "RecipientEmailAddress"
)

func PossibleValuesForSecurityMoveToDeletedItemsResponseActionIdentifier() []string {
	return []string{
		string(SecurityMoveToDeletedItemsResponseActionIdentifiernetworkMessageId),
		string(SecurityMoveToDeletedItemsResponseActionIdentifierrecipientEmailAddress),
	}
}

func parseSecurityMoveToDeletedItemsResponseActionIdentifier(input string) (*SecurityMoveToDeletedItemsResponseActionIdentifier, error) {
	vals := map[string]SecurityMoveToDeletedItemsResponseActionIdentifier{
		"networkmessageid":      SecurityMoveToDeletedItemsResponseActionIdentifiernetworkMessageId,
		"recipientemailaddress": SecurityMoveToDeletedItemsResponseActionIdentifierrecipientEmailAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMoveToDeletedItemsResponseActionIdentifier(input)
	return &out, nil
}
