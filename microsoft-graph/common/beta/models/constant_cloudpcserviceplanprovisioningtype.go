package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCServicePlanProvisioningType string

const (
	CloudPCServicePlanProvisioningTypededicated CloudPCServicePlanProvisioningType = "Dedicated"
	CloudPCServicePlanProvisioningTypeshared    CloudPCServicePlanProvisioningType = "Shared"
)

func PossibleValuesForCloudPCServicePlanProvisioningType() []string {
	return []string{
		string(CloudPCServicePlanProvisioningTypededicated),
		string(CloudPCServicePlanProvisioningTypeshared),
	}
}

func parseCloudPCServicePlanProvisioningType(input string) (*CloudPCServicePlanProvisioningType, error) {
	vals := map[string]CloudPCServicePlanProvisioningType{
		"dedicated": CloudPCServicePlanProvisioningTypededicated,
		"shared":    CloudPCServicePlanProvisioningTypeshared,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCServicePlanProvisioningType(input)
	return &out, nil
}
