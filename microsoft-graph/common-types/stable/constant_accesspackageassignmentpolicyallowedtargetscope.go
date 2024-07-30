package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentPolicyAllowedTargetScope string

const (
	AccessPackageAssignmentPolicyAllowedTargetScope_AllConfiguredConnectedOrganizationUsers AccessPackageAssignmentPolicyAllowedTargetScope = "allConfiguredConnectedOrganizationUsers"
	AccessPackageAssignmentPolicyAllowedTargetScope_AllDirectoryServicePrincipals           AccessPackageAssignmentPolicyAllowedTargetScope = "allDirectoryServicePrincipals"
	AccessPackageAssignmentPolicyAllowedTargetScope_AllDirectoryUsers                       AccessPackageAssignmentPolicyAllowedTargetScope = "allDirectoryUsers"
	AccessPackageAssignmentPolicyAllowedTargetScope_AllExternalUsers                        AccessPackageAssignmentPolicyAllowedTargetScope = "allExternalUsers"
	AccessPackageAssignmentPolicyAllowedTargetScope_AllMemberUsers                          AccessPackageAssignmentPolicyAllowedTargetScope = "allMemberUsers"
	AccessPackageAssignmentPolicyAllowedTargetScope_NotSpecified                            AccessPackageAssignmentPolicyAllowedTargetScope = "notSpecified"
	AccessPackageAssignmentPolicyAllowedTargetScope_SpecificConnectedOrganizationUsers      AccessPackageAssignmentPolicyAllowedTargetScope = "specificConnectedOrganizationUsers"
	AccessPackageAssignmentPolicyAllowedTargetScope_SpecificDirectoryServicePrincipals      AccessPackageAssignmentPolicyAllowedTargetScope = "specificDirectoryServicePrincipals"
	AccessPackageAssignmentPolicyAllowedTargetScope_SpecificDirectoryUsers                  AccessPackageAssignmentPolicyAllowedTargetScope = "specificDirectoryUsers"
)

func PossibleValuesForAccessPackageAssignmentPolicyAllowedTargetScope() []string {
	return []string{
		string(AccessPackageAssignmentPolicyAllowedTargetScope_AllConfiguredConnectedOrganizationUsers),
		string(AccessPackageAssignmentPolicyAllowedTargetScope_AllDirectoryServicePrincipals),
		string(AccessPackageAssignmentPolicyAllowedTargetScope_AllDirectoryUsers),
		string(AccessPackageAssignmentPolicyAllowedTargetScope_AllExternalUsers),
		string(AccessPackageAssignmentPolicyAllowedTargetScope_AllMemberUsers),
		string(AccessPackageAssignmentPolicyAllowedTargetScope_NotSpecified),
		string(AccessPackageAssignmentPolicyAllowedTargetScope_SpecificConnectedOrganizationUsers),
		string(AccessPackageAssignmentPolicyAllowedTargetScope_SpecificDirectoryServicePrincipals),
		string(AccessPackageAssignmentPolicyAllowedTargetScope_SpecificDirectoryUsers),
	}
}

func (s *AccessPackageAssignmentPolicyAllowedTargetScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageAssignmentPolicyAllowedTargetScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageAssignmentPolicyAllowedTargetScope(input string) (*AccessPackageAssignmentPolicyAllowedTargetScope, error) {
	vals := map[string]AccessPackageAssignmentPolicyAllowedTargetScope{
		"allconfiguredconnectedorganizationusers": AccessPackageAssignmentPolicyAllowedTargetScope_AllConfiguredConnectedOrganizationUsers,
		"alldirectoryserviceprincipals":           AccessPackageAssignmentPolicyAllowedTargetScope_AllDirectoryServicePrincipals,
		"alldirectoryusers":                       AccessPackageAssignmentPolicyAllowedTargetScope_AllDirectoryUsers,
		"allexternalusers":                        AccessPackageAssignmentPolicyAllowedTargetScope_AllExternalUsers,
		"allmemberusers":                          AccessPackageAssignmentPolicyAllowedTargetScope_AllMemberUsers,
		"notspecified":                            AccessPackageAssignmentPolicyAllowedTargetScope_NotSpecified,
		"specificconnectedorganizationusers":      AccessPackageAssignmentPolicyAllowedTargetScope_SpecificConnectedOrganizationUsers,
		"specificdirectoryserviceprincipals":      AccessPackageAssignmentPolicyAllowedTargetScope_SpecificDirectoryServicePrincipals,
		"specificdirectoryusers":                  AccessPackageAssignmentPolicyAllowedTargetScope_SpecificDirectoryUsers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentPolicyAllowedTargetScope(input)
	return &out, nil
}
