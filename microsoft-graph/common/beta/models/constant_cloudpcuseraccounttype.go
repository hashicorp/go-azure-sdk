package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCUserAccountType string

const (
	CloudPCUserAccountTypeadministrator CloudPCUserAccountType = "Administrator"
	CloudPCUserAccountTypestandardUser  CloudPCUserAccountType = "StandardUser"
)

func PossibleValuesForCloudPCUserAccountType() []string {
	return []string{
		string(CloudPCUserAccountTypeadministrator),
		string(CloudPCUserAccountTypestandardUser),
	}
}

func parseCloudPCUserAccountType(input string) (*CloudPCUserAccountType, error) {
	vals := map[string]CloudPCUserAccountType{
		"administrator": CloudPCUserAccountTypeadministrator,
		"standarduser":  CloudPCUserAccountTypestandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCUserAccountType(input)
	return &out, nil
}
