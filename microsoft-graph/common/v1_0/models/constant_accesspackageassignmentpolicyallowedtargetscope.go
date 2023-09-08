package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentPolicyAllowedTargetScope string

const (
	AccessPackageAssignmentPolicyAllowedTargetScopeallConfiguredConnectedOrganizationUsers AccessPackageAssignmentPolicyAllowedTargetScope = "AllConfiguredConnectedOrganizationUsers"
	AccessPackageAssignmentPolicyAllowedTargetScopeallDirectoryServicePrincipals           AccessPackageAssignmentPolicyAllowedTargetScope = "AllDirectoryServicePrincipals"
	AccessPackageAssignmentPolicyAllowedTargetScopeallDirectoryUsers                       AccessPackageAssignmentPolicyAllowedTargetScope = "AllDirectoryUsers"
	AccessPackageAssignmentPolicyAllowedTargetScopeallExternalUsers                        AccessPackageAssignmentPolicyAllowedTargetScope = "AllExternalUsers"
	AccessPackageAssignmentPolicyAllowedTargetScopeallMemberUsers                          AccessPackageAssignmentPolicyAllowedTargetScope = "AllMemberUsers"
	AccessPackageAssignmentPolicyAllowedTargetScopenotSpecified                            AccessPackageAssignmentPolicyAllowedTargetScope = "NotSpecified"
	AccessPackageAssignmentPolicyAllowedTargetScopespecificConnectedOrganizationUsers      AccessPackageAssignmentPolicyAllowedTargetScope = "SpecificConnectedOrganizationUsers"
	AccessPackageAssignmentPolicyAllowedTargetScopespecificDirectoryServicePrincipals      AccessPackageAssignmentPolicyAllowedTargetScope = "SpecificDirectoryServicePrincipals"
	AccessPackageAssignmentPolicyAllowedTargetScopespecificDirectoryUsers                  AccessPackageAssignmentPolicyAllowedTargetScope = "SpecificDirectoryUsers"
)

func PossibleValuesForAccessPackageAssignmentPolicyAllowedTargetScope() []string {
	return []string{
		string(AccessPackageAssignmentPolicyAllowedTargetScopeallConfiguredConnectedOrganizationUsers),
		string(AccessPackageAssignmentPolicyAllowedTargetScopeallDirectoryServicePrincipals),
		string(AccessPackageAssignmentPolicyAllowedTargetScopeallDirectoryUsers),
		string(AccessPackageAssignmentPolicyAllowedTargetScopeallExternalUsers),
		string(AccessPackageAssignmentPolicyAllowedTargetScopeallMemberUsers),
		string(AccessPackageAssignmentPolicyAllowedTargetScopenotSpecified),
		string(AccessPackageAssignmentPolicyAllowedTargetScopespecificConnectedOrganizationUsers),
		string(AccessPackageAssignmentPolicyAllowedTargetScopespecificDirectoryServicePrincipals),
		string(AccessPackageAssignmentPolicyAllowedTargetScopespecificDirectoryUsers),
	}
}

func parseAccessPackageAssignmentPolicyAllowedTargetScope(input string) (*AccessPackageAssignmentPolicyAllowedTargetScope, error) {
	vals := map[string]AccessPackageAssignmentPolicyAllowedTargetScope{
		"allconfiguredconnectedorganizationusers": AccessPackageAssignmentPolicyAllowedTargetScopeallConfiguredConnectedOrganizationUsers,
		"alldirectoryserviceprincipals":           AccessPackageAssignmentPolicyAllowedTargetScopeallDirectoryServicePrincipals,
		"alldirectoryusers":                       AccessPackageAssignmentPolicyAllowedTargetScopeallDirectoryUsers,
		"allexternalusers":                        AccessPackageAssignmentPolicyAllowedTargetScopeallExternalUsers,
		"allmemberusers":                          AccessPackageAssignmentPolicyAllowedTargetScopeallMemberUsers,
		"notspecified":                            AccessPackageAssignmentPolicyAllowedTargetScopenotSpecified,
		"specificconnectedorganizationusers":      AccessPackageAssignmentPolicyAllowedTargetScopespecificConnectedOrganizationUsers,
		"specificdirectoryserviceprincipals":      AccessPackageAssignmentPolicyAllowedTargetScopespecificDirectoryServicePrincipals,
		"specificdirectoryusers":                  AccessPackageAssignmentPolicyAllowedTargetScopespecificDirectoryUsers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentPolicyAllowedTargetScope(input)
	return &out, nil
}
