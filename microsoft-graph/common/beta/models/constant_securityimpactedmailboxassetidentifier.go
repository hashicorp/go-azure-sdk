package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityImpactedMailboxAssetIdentifier string

const (
	SecurityImpactedMailboxAssetIdentifieraccountUpn                  SecurityImpactedMailboxAssetIdentifier = "AccountUpn"
	SecurityImpactedMailboxAssetIdentifierfileOwnerUpn                SecurityImpactedMailboxAssetIdentifier = "FileOwnerUpn"
	SecurityImpactedMailboxAssetIdentifierinitiatingProcessAccountUpn SecurityImpactedMailboxAssetIdentifier = "InitiatingProcessAccountUpn"
	SecurityImpactedMailboxAssetIdentifierlastModifyingAccountUpn     SecurityImpactedMailboxAssetIdentifier = "LastModifyingAccountUpn"
	SecurityImpactedMailboxAssetIdentifierrecipientEmailAddress       SecurityImpactedMailboxAssetIdentifier = "RecipientEmailAddress"
	SecurityImpactedMailboxAssetIdentifiersenderDisplayName           SecurityImpactedMailboxAssetIdentifier = "SenderDisplayName"
	SecurityImpactedMailboxAssetIdentifiersenderFromAddress           SecurityImpactedMailboxAssetIdentifier = "SenderFromAddress"
	SecurityImpactedMailboxAssetIdentifiersenderMailFromAddress       SecurityImpactedMailboxAssetIdentifier = "SenderMailFromAddress"
	SecurityImpactedMailboxAssetIdentifiertargetAccountUpn            SecurityImpactedMailboxAssetIdentifier = "TargetAccountUpn"
)

func PossibleValuesForSecurityImpactedMailboxAssetIdentifier() []string {
	return []string{
		string(SecurityImpactedMailboxAssetIdentifieraccountUpn),
		string(SecurityImpactedMailboxAssetIdentifierfileOwnerUpn),
		string(SecurityImpactedMailboxAssetIdentifierinitiatingProcessAccountUpn),
		string(SecurityImpactedMailboxAssetIdentifierlastModifyingAccountUpn),
		string(SecurityImpactedMailboxAssetIdentifierrecipientEmailAddress),
		string(SecurityImpactedMailboxAssetIdentifiersenderDisplayName),
		string(SecurityImpactedMailboxAssetIdentifiersenderFromAddress),
		string(SecurityImpactedMailboxAssetIdentifiersenderMailFromAddress),
		string(SecurityImpactedMailboxAssetIdentifiertargetAccountUpn),
	}
}

func parseSecurityImpactedMailboxAssetIdentifier(input string) (*SecurityImpactedMailboxAssetIdentifier, error) {
	vals := map[string]SecurityImpactedMailboxAssetIdentifier{
		"accountupn":                  SecurityImpactedMailboxAssetIdentifieraccountUpn,
		"fileownerupn":                SecurityImpactedMailboxAssetIdentifierfileOwnerUpn,
		"initiatingprocessaccountupn": SecurityImpactedMailboxAssetIdentifierinitiatingProcessAccountUpn,
		"lastmodifyingaccountupn":     SecurityImpactedMailboxAssetIdentifierlastModifyingAccountUpn,
		"recipientemailaddress":       SecurityImpactedMailboxAssetIdentifierrecipientEmailAddress,
		"senderdisplayname":           SecurityImpactedMailboxAssetIdentifiersenderDisplayName,
		"senderfromaddress":           SecurityImpactedMailboxAssetIdentifiersenderFromAddress,
		"sendermailfromaddress":       SecurityImpactedMailboxAssetIdentifiersenderMailFromAddress,
		"targetaccountupn":            SecurityImpactedMailboxAssetIdentifiertargetAccountUpn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityImpactedMailboxAssetIdentifier(input)
	return &out, nil
}
