package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationMobileDeviceManagementAuthority string

const (
	OrganizationMobileDeviceManagementAuthorityintune    OrganizationMobileDeviceManagementAuthority = "Intune"
	OrganizationMobileDeviceManagementAuthorityoffice365 OrganizationMobileDeviceManagementAuthority = "Office365"
	OrganizationMobileDeviceManagementAuthoritysccm      OrganizationMobileDeviceManagementAuthority = "Sccm"
	OrganizationMobileDeviceManagementAuthorityunknown   OrganizationMobileDeviceManagementAuthority = "Unknown"
)

func PossibleValuesForOrganizationMobileDeviceManagementAuthority() []string {
	return []string{
		string(OrganizationMobileDeviceManagementAuthorityintune),
		string(OrganizationMobileDeviceManagementAuthorityoffice365),
		string(OrganizationMobileDeviceManagementAuthoritysccm),
		string(OrganizationMobileDeviceManagementAuthorityunknown),
	}
}

func parseOrganizationMobileDeviceManagementAuthority(input string) (*OrganizationMobileDeviceManagementAuthority, error) {
	vals := map[string]OrganizationMobileDeviceManagementAuthority{
		"intune":    OrganizationMobileDeviceManagementAuthorityintune,
		"office365": OrganizationMobileDeviceManagementAuthorityoffice365,
		"sccm":      OrganizationMobileDeviceManagementAuthoritysccm,
		"unknown":   OrganizationMobileDeviceManagementAuthorityunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrganizationMobileDeviceManagementAuthority(input)
	return &out, nil
}
