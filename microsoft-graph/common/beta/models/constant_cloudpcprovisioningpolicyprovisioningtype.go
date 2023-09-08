package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCProvisioningPolicyProvisioningType string

const (
	CloudPCProvisioningPolicyProvisioningTypededicated CloudPCProvisioningPolicyProvisioningType = "Dedicated"
	CloudPCProvisioningPolicyProvisioningTypeshared    CloudPCProvisioningPolicyProvisioningType = "Shared"
)

func PossibleValuesForCloudPCProvisioningPolicyProvisioningType() []string {
	return []string{
		string(CloudPCProvisioningPolicyProvisioningTypededicated),
		string(CloudPCProvisioningPolicyProvisioningTypeshared),
	}
}

func parseCloudPCProvisioningPolicyProvisioningType(input string) (*CloudPCProvisioningPolicyProvisioningType, error) {
	vals := map[string]CloudPCProvisioningPolicyProvisioningType{
		"dedicated": CloudPCProvisioningPolicyProvisioningTypededicated,
		"shared":    CloudPCProvisioningPolicyProvisioningTypeshared,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCProvisioningPolicyProvisioningType(input)
	return &out, nil
}
