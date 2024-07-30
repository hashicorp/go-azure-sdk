package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityImpactedMailboxAssetIdentifier string

const (
	SecurityImpactedMailboxAssetIdentifier_AccountUpn                  SecurityImpactedMailboxAssetIdentifier = "accountUpn"
	SecurityImpactedMailboxAssetIdentifier_FileOwnerUpn                SecurityImpactedMailboxAssetIdentifier = "fileOwnerUpn"
	SecurityImpactedMailboxAssetIdentifier_InitiatingProcessAccountUpn SecurityImpactedMailboxAssetIdentifier = "initiatingProcessAccountUpn"
	SecurityImpactedMailboxAssetIdentifier_LastModifyingAccountUpn     SecurityImpactedMailboxAssetIdentifier = "lastModifyingAccountUpn"
	SecurityImpactedMailboxAssetIdentifier_RecipientEmailAddress       SecurityImpactedMailboxAssetIdentifier = "recipientEmailAddress"
	SecurityImpactedMailboxAssetIdentifier_SenderDisplayName           SecurityImpactedMailboxAssetIdentifier = "senderDisplayName"
	SecurityImpactedMailboxAssetIdentifier_SenderFromAddress           SecurityImpactedMailboxAssetIdentifier = "senderFromAddress"
	SecurityImpactedMailboxAssetIdentifier_SenderMailFromAddress       SecurityImpactedMailboxAssetIdentifier = "senderMailFromAddress"
	SecurityImpactedMailboxAssetIdentifier_TargetAccountUpn            SecurityImpactedMailboxAssetIdentifier = "targetAccountUpn"
)

func PossibleValuesForSecurityImpactedMailboxAssetIdentifier() []string {
	return []string{
		string(SecurityImpactedMailboxAssetIdentifier_AccountUpn),
		string(SecurityImpactedMailboxAssetIdentifier_FileOwnerUpn),
		string(SecurityImpactedMailboxAssetIdentifier_InitiatingProcessAccountUpn),
		string(SecurityImpactedMailboxAssetIdentifier_LastModifyingAccountUpn),
		string(SecurityImpactedMailboxAssetIdentifier_RecipientEmailAddress),
		string(SecurityImpactedMailboxAssetIdentifier_SenderDisplayName),
		string(SecurityImpactedMailboxAssetIdentifier_SenderFromAddress),
		string(SecurityImpactedMailboxAssetIdentifier_SenderMailFromAddress),
		string(SecurityImpactedMailboxAssetIdentifier_TargetAccountUpn),
	}
}

func (s *SecurityImpactedMailboxAssetIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityImpactedMailboxAssetIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityImpactedMailboxAssetIdentifier(input string) (*SecurityImpactedMailboxAssetIdentifier, error) {
	vals := map[string]SecurityImpactedMailboxAssetIdentifier{
		"accountupn":                  SecurityImpactedMailboxAssetIdentifier_AccountUpn,
		"fileownerupn":                SecurityImpactedMailboxAssetIdentifier_FileOwnerUpn,
		"initiatingprocessaccountupn": SecurityImpactedMailboxAssetIdentifier_InitiatingProcessAccountUpn,
		"lastmodifyingaccountupn":     SecurityImpactedMailboxAssetIdentifier_LastModifyingAccountUpn,
		"recipientemailaddress":       SecurityImpactedMailboxAssetIdentifier_RecipientEmailAddress,
		"senderdisplayname":           SecurityImpactedMailboxAssetIdentifier_SenderDisplayName,
		"senderfromaddress":           SecurityImpactedMailboxAssetIdentifier_SenderFromAddress,
		"sendermailfromaddress":       SecurityImpactedMailboxAssetIdentifier_SenderMailFromAddress,
		"targetaccountupn":            SecurityImpactedMailboxAssetIdentifier_TargetAccountUpn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityImpactedMailboxAssetIdentifier(input)
	return &out, nil
}
