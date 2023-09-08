package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOrganizationSettingsOsVersion string

const (
	CloudPCOrganizationSettingsOsVersionwindows10 CloudPCOrganizationSettingsOsVersion = "Windows10"
	CloudPCOrganizationSettingsOsVersionwindows11 CloudPCOrganizationSettingsOsVersion = "Windows11"
)

func PossibleValuesForCloudPCOrganizationSettingsOsVersion() []string {
	return []string{
		string(CloudPCOrganizationSettingsOsVersionwindows10),
		string(CloudPCOrganizationSettingsOsVersionwindows11),
	}
}

func parseCloudPCOrganizationSettingsOsVersion(input string) (*CloudPCOrganizationSettingsOsVersion, error) {
	vals := map[string]CloudPCOrganizationSettingsOsVersion{
		"windows10": CloudPCOrganizationSettingsOsVersionwindows10,
		"windows11": CloudPCOrganizationSettingsOsVersionwindows11,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOrganizationSettingsOsVersion(input)
	return &out, nil
}
