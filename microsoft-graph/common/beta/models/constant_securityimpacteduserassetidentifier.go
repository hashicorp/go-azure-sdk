package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityImpactedUserAssetIdentifier string

const (
	SecurityImpactedUserAssetIdentifieraccountDomain               SecurityImpactedUserAssetIdentifier = "AccountDomain"
	SecurityImpactedUserAssetIdentifieraccountId                   SecurityImpactedUserAssetIdentifier = "AccountId"
	SecurityImpactedUserAssetIdentifieraccountName                 SecurityImpactedUserAssetIdentifier = "AccountName"
	SecurityImpactedUserAssetIdentifieraccountObjectId             SecurityImpactedUserAssetIdentifier = "AccountObjectId"
	SecurityImpactedUserAssetIdentifieraccountSid                  SecurityImpactedUserAssetIdentifier = "AccountSid"
	SecurityImpactedUserAssetIdentifieraccountUpn                  SecurityImpactedUserAssetIdentifier = "AccountUpn"
	SecurityImpactedUserAssetIdentifierinitiatingAccountDomain     SecurityImpactedUserAssetIdentifier = "InitiatingAccountDomain"
	SecurityImpactedUserAssetIdentifierinitiatingAccountName       SecurityImpactedUserAssetIdentifier = "InitiatingAccountName"
	SecurityImpactedUserAssetIdentifierinitiatingAccountSid        SecurityImpactedUserAssetIdentifier = "InitiatingAccountSid"
	SecurityImpactedUserAssetIdentifierinitiatingProcessAccountUpn SecurityImpactedUserAssetIdentifier = "InitiatingProcessAccountUpn"
	SecurityImpactedUserAssetIdentifierprocessAccountObjectId      SecurityImpactedUserAssetIdentifier = "ProcessAccountObjectId"
	SecurityImpactedUserAssetIdentifierrecipientObjectId           SecurityImpactedUserAssetIdentifier = "RecipientObjectId"
	SecurityImpactedUserAssetIdentifierrequestAccountDomain        SecurityImpactedUserAssetIdentifier = "RequestAccountDomain"
	SecurityImpactedUserAssetIdentifierrequestAccountName          SecurityImpactedUserAssetIdentifier = "RequestAccountName"
	SecurityImpactedUserAssetIdentifierrequestAccountSid           SecurityImpactedUserAssetIdentifier = "RequestAccountSid"
	SecurityImpactedUserAssetIdentifierservicePrincipalId          SecurityImpactedUserAssetIdentifier = "ServicePrincipalId"
	SecurityImpactedUserAssetIdentifierservicePrincipalName        SecurityImpactedUserAssetIdentifier = "ServicePrincipalName"
	SecurityImpactedUserAssetIdentifiertargetAccountUpn            SecurityImpactedUserAssetIdentifier = "TargetAccountUpn"
)

func PossibleValuesForSecurityImpactedUserAssetIdentifier() []string {
	return []string{
		string(SecurityImpactedUserAssetIdentifieraccountDomain),
		string(SecurityImpactedUserAssetIdentifieraccountId),
		string(SecurityImpactedUserAssetIdentifieraccountName),
		string(SecurityImpactedUserAssetIdentifieraccountObjectId),
		string(SecurityImpactedUserAssetIdentifieraccountSid),
		string(SecurityImpactedUserAssetIdentifieraccountUpn),
		string(SecurityImpactedUserAssetIdentifierinitiatingAccountDomain),
		string(SecurityImpactedUserAssetIdentifierinitiatingAccountName),
		string(SecurityImpactedUserAssetIdentifierinitiatingAccountSid),
		string(SecurityImpactedUserAssetIdentifierinitiatingProcessAccountUpn),
		string(SecurityImpactedUserAssetIdentifierprocessAccountObjectId),
		string(SecurityImpactedUserAssetIdentifierrecipientObjectId),
		string(SecurityImpactedUserAssetIdentifierrequestAccountDomain),
		string(SecurityImpactedUserAssetIdentifierrequestAccountName),
		string(SecurityImpactedUserAssetIdentifierrequestAccountSid),
		string(SecurityImpactedUserAssetIdentifierservicePrincipalId),
		string(SecurityImpactedUserAssetIdentifierservicePrincipalName),
		string(SecurityImpactedUserAssetIdentifiertargetAccountUpn),
	}
}

func parseSecurityImpactedUserAssetIdentifier(input string) (*SecurityImpactedUserAssetIdentifier, error) {
	vals := map[string]SecurityImpactedUserAssetIdentifier{
		"accountdomain":               SecurityImpactedUserAssetIdentifieraccountDomain,
		"accountid":                   SecurityImpactedUserAssetIdentifieraccountId,
		"accountname":                 SecurityImpactedUserAssetIdentifieraccountName,
		"accountobjectid":             SecurityImpactedUserAssetIdentifieraccountObjectId,
		"accountsid":                  SecurityImpactedUserAssetIdentifieraccountSid,
		"accountupn":                  SecurityImpactedUserAssetIdentifieraccountUpn,
		"initiatingaccountdomain":     SecurityImpactedUserAssetIdentifierinitiatingAccountDomain,
		"initiatingaccountname":       SecurityImpactedUserAssetIdentifierinitiatingAccountName,
		"initiatingaccountsid":        SecurityImpactedUserAssetIdentifierinitiatingAccountSid,
		"initiatingprocessaccountupn": SecurityImpactedUserAssetIdentifierinitiatingProcessAccountUpn,
		"processaccountobjectid":      SecurityImpactedUserAssetIdentifierprocessAccountObjectId,
		"recipientobjectid":           SecurityImpactedUserAssetIdentifierrecipientObjectId,
		"requestaccountdomain":        SecurityImpactedUserAssetIdentifierrequestAccountDomain,
		"requestaccountname":          SecurityImpactedUserAssetIdentifierrequestAccountName,
		"requestaccountsid":           SecurityImpactedUserAssetIdentifierrequestAccountSid,
		"serviceprincipalid":          SecurityImpactedUserAssetIdentifierservicePrincipalId,
		"serviceprincipalname":        SecurityImpactedUserAssetIdentifierservicePrincipalName,
		"targetaccountupn":            SecurityImpactedUserAssetIdentifiertargetAccountUpn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityImpactedUserAssetIdentifier(input)
	return &out, nil
}
