package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityImpactedUserAssetIdentifier string

const (
	SecurityImpactedUserAssetIdentifier_AccountDomain               SecurityImpactedUserAssetIdentifier = "accountDomain"
	SecurityImpactedUserAssetIdentifier_AccountId                   SecurityImpactedUserAssetIdentifier = "accountId"
	SecurityImpactedUserAssetIdentifier_AccountName                 SecurityImpactedUserAssetIdentifier = "accountName"
	SecurityImpactedUserAssetIdentifier_AccountObjectId             SecurityImpactedUserAssetIdentifier = "accountObjectId"
	SecurityImpactedUserAssetIdentifier_AccountSid                  SecurityImpactedUserAssetIdentifier = "accountSid"
	SecurityImpactedUserAssetIdentifier_AccountUpn                  SecurityImpactedUserAssetIdentifier = "accountUpn"
	SecurityImpactedUserAssetIdentifier_InitiatingAccountDomain     SecurityImpactedUserAssetIdentifier = "initiatingAccountDomain"
	SecurityImpactedUserAssetIdentifier_InitiatingAccountName       SecurityImpactedUserAssetIdentifier = "initiatingAccountName"
	SecurityImpactedUserAssetIdentifier_InitiatingAccountSid        SecurityImpactedUserAssetIdentifier = "initiatingAccountSid"
	SecurityImpactedUserAssetIdentifier_InitiatingProcessAccountUpn SecurityImpactedUserAssetIdentifier = "initiatingProcessAccountUpn"
	SecurityImpactedUserAssetIdentifier_ProcessAccountObjectId      SecurityImpactedUserAssetIdentifier = "processAccountObjectId"
	SecurityImpactedUserAssetIdentifier_RecipientObjectId           SecurityImpactedUserAssetIdentifier = "recipientObjectId"
	SecurityImpactedUserAssetIdentifier_RequestAccountDomain        SecurityImpactedUserAssetIdentifier = "requestAccountDomain"
	SecurityImpactedUserAssetIdentifier_RequestAccountName          SecurityImpactedUserAssetIdentifier = "requestAccountName"
	SecurityImpactedUserAssetIdentifier_RequestAccountSid           SecurityImpactedUserAssetIdentifier = "requestAccountSid"
	SecurityImpactedUserAssetIdentifier_ServicePrincipalId          SecurityImpactedUserAssetIdentifier = "servicePrincipalId"
	SecurityImpactedUserAssetIdentifier_ServicePrincipalName        SecurityImpactedUserAssetIdentifier = "servicePrincipalName"
	SecurityImpactedUserAssetIdentifier_TargetAccountUpn            SecurityImpactedUserAssetIdentifier = "targetAccountUpn"
)

func PossibleValuesForSecurityImpactedUserAssetIdentifier() []string {
	return []string{
		string(SecurityImpactedUserAssetIdentifier_AccountDomain),
		string(SecurityImpactedUserAssetIdentifier_AccountId),
		string(SecurityImpactedUserAssetIdentifier_AccountName),
		string(SecurityImpactedUserAssetIdentifier_AccountObjectId),
		string(SecurityImpactedUserAssetIdentifier_AccountSid),
		string(SecurityImpactedUserAssetIdentifier_AccountUpn),
		string(SecurityImpactedUserAssetIdentifier_InitiatingAccountDomain),
		string(SecurityImpactedUserAssetIdentifier_InitiatingAccountName),
		string(SecurityImpactedUserAssetIdentifier_InitiatingAccountSid),
		string(SecurityImpactedUserAssetIdentifier_InitiatingProcessAccountUpn),
		string(SecurityImpactedUserAssetIdentifier_ProcessAccountObjectId),
		string(SecurityImpactedUserAssetIdentifier_RecipientObjectId),
		string(SecurityImpactedUserAssetIdentifier_RequestAccountDomain),
		string(SecurityImpactedUserAssetIdentifier_RequestAccountName),
		string(SecurityImpactedUserAssetIdentifier_RequestAccountSid),
		string(SecurityImpactedUserAssetIdentifier_ServicePrincipalId),
		string(SecurityImpactedUserAssetIdentifier_ServicePrincipalName),
		string(SecurityImpactedUserAssetIdentifier_TargetAccountUpn),
	}
}

func (s *SecurityImpactedUserAssetIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityImpactedUserAssetIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityImpactedUserAssetIdentifier(input string) (*SecurityImpactedUserAssetIdentifier, error) {
	vals := map[string]SecurityImpactedUserAssetIdentifier{
		"accountdomain":               SecurityImpactedUserAssetIdentifier_AccountDomain,
		"accountid":                   SecurityImpactedUserAssetIdentifier_AccountId,
		"accountname":                 SecurityImpactedUserAssetIdentifier_AccountName,
		"accountobjectid":             SecurityImpactedUserAssetIdentifier_AccountObjectId,
		"accountsid":                  SecurityImpactedUserAssetIdentifier_AccountSid,
		"accountupn":                  SecurityImpactedUserAssetIdentifier_AccountUpn,
		"initiatingaccountdomain":     SecurityImpactedUserAssetIdentifier_InitiatingAccountDomain,
		"initiatingaccountname":       SecurityImpactedUserAssetIdentifier_InitiatingAccountName,
		"initiatingaccountsid":        SecurityImpactedUserAssetIdentifier_InitiatingAccountSid,
		"initiatingprocessaccountupn": SecurityImpactedUserAssetIdentifier_InitiatingProcessAccountUpn,
		"processaccountobjectid":      SecurityImpactedUserAssetIdentifier_ProcessAccountObjectId,
		"recipientobjectid":           SecurityImpactedUserAssetIdentifier_RecipientObjectId,
		"requestaccountdomain":        SecurityImpactedUserAssetIdentifier_RequestAccountDomain,
		"requestaccountname":          SecurityImpactedUserAssetIdentifier_RequestAccountName,
		"requestaccountsid":           SecurityImpactedUserAssetIdentifier_RequestAccountSid,
		"serviceprincipalid":          SecurityImpactedUserAssetIdentifier_ServicePrincipalId,
		"serviceprincipalname":        SecurityImpactedUserAssetIdentifier_ServicePrincipalName,
		"targetaccountupn":            SecurityImpactedUserAssetIdentifier_TargetAccountUpn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityImpactedUserAssetIdentifier(input)
	return &out, nil
}
