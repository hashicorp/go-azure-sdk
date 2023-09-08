package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KerberosSignOnSettingsKerberosSignOnMappingAttributeType string

const (
	KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeonPremisesSAMAccountName        KerberosSignOnSettingsKerberosSignOnMappingAttributeType = "OnPremisesSAMAccountName"
	KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeonPremisesUserPrincipalName     KerberosSignOnSettingsKerberosSignOnMappingAttributeType = "OnPremisesUserPrincipalName"
	KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeonPremisesUserPrincipalUsername KerberosSignOnSettingsKerberosSignOnMappingAttributeType = "OnPremisesUserPrincipalUsername"
	KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeuserPrincipalName               KerberosSignOnSettingsKerberosSignOnMappingAttributeType = "UserPrincipalName"
	KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeuserPrincipalUsername           KerberosSignOnSettingsKerberosSignOnMappingAttributeType = "UserPrincipalUsername"
)

func PossibleValuesForKerberosSignOnSettingsKerberosSignOnMappingAttributeType() []string {
	return []string{
		string(KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeonPremisesSAMAccountName),
		string(KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeonPremisesUserPrincipalName),
		string(KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeonPremisesUserPrincipalUsername),
		string(KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeuserPrincipalName),
		string(KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeuserPrincipalUsername),
	}
}

func parseKerberosSignOnSettingsKerberosSignOnMappingAttributeType(input string) (*KerberosSignOnSettingsKerberosSignOnMappingAttributeType, error) {
	vals := map[string]KerberosSignOnSettingsKerberosSignOnMappingAttributeType{
		"onpremisessamaccountname":        KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeonPremisesSAMAccountName,
		"onpremisesuserprincipalname":     KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeonPremisesUserPrincipalName,
		"onpremisesuserprincipalusername": KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeonPremisesUserPrincipalUsername,
		"userprincipalname":               KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeuserPrincipalName,
		"userprincipalusername":           KerberosSignOnSettingsKerberosSignOnMappingAttributeTypeuserPrincipalUsername,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KerberosSignOnSettingsKerberosSignOnMappingAttributeType(input)
	return &out, nil
}
