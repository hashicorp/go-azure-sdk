package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCProvisioningPolicyManagedBy string

const (
	CloudPCProvisioningPolicyManagedBydevBox     CloudPCProvisioningPolicyManagedBy = "DevBox"
	CloudPCProvisioningPolicyManagedByrpaBox     CloudPCProvisioningPolicyManagedBy = "RpaBox"
	CloudPCProvisioningPolicyManagedBywindows365 CloudPCProvisioningPolicyManagedBy = "Windows365"
)

func PossibleValuesForCloudPCProvisioningPolicyManagedBy() []string {
	return []string{
		string(CloudPCProvisioningPolicyManagedBydevBox),
		string(CloudPCProvisioningPolicyManagedByrpaBox),
		string(CloudPCProvisioningPolicyManagedBywindows365),
	}
}

func parseCloudPCProvisioningPolicyManagedBy(input string) (*CloudPCProvisioningPolicyManagedBy, error) {
	vals := map[string]CloudPCProvisioningPolicyManagedBy{
		"devbox":     CloudPCProvisioningPolicyManagedBydevBox,
		"rpabox":     CloudPCProvisioningPolicyManagedByrpaBox,
		"windows365": CloudPCProvisioningPolicyManagedBywindows365,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCProvisioningPolicyManagedBy(input)
	return &out, nil
}
