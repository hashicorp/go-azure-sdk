package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCProvisioningType string

const (
	CloudPCProvisioningTypededicated CloudPCProvisioningType = "Dedicated"
	CloudPCProvisioningTypeshared    CloudPCProvisioningType = "Shared"
)

func PossibleValuesForCloudPCProvisioningType() []string {
	return []string{
		string(CloudPCProvisioningTypededicated),
		string(CloudPCProvisioningTypeshared),
	}
}

func parseCloudPCProvisioningType(input string) (*CloudPCProvisioningType, error) {
	vals := map[string]CloudPCProvisioningType{
		"dedicated": CloudPCProvisioningTypededicated,
		"shared":    CloudPCProvisioningTypeshared,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCProvisioningType(input)
	return &out, nil
}
