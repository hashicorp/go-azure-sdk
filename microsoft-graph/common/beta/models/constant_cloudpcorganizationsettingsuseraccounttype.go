package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOrganizationSettingsUserAccountType string

const (
	CloudPCOrganizationSettingsUserAccountTypeadministrator CloudPCOrganizationSettingsUserAccountType = "Administrator"
	CloudPCOrganizationSettingsUserAccountTypestandardUser  CloudPCOrganizationSettingsUserAccountType = "StandardUser"
)

func PossibleValuesForCloudPCOrganizationSettingsUserAccountType() []string {
	return []string{
		string(CloudPCOrganizationSettingsUserAccountTypeadministrator),
		string(CloudPCOrganizationSettingsUserAccountTypestandardUser),
	}
}

func parseCloudPCOrganizationSettingsUserAccountType(input string) (*CloudPCOrganizationSettingsUserAccountType, error) {
	vals := map[string]CloudPCOrganizationSettingsUserAccountType{
		"administrator": CloudPCOrganizationSettingsUserAccountTypeadministrator,
		"standarduser":  CloudPCOrganizationSettingsUserAccountTypestandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOrganizationSettingsUserAccountType(input)
	return &out, nil
}
