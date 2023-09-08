package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCServicePlanType string

const (
	CloudPCServicePlanTypebusiness   CloudPCServicePlanType = "Business"
	CloudPCServicePlanTypeenterprise CloudPCServicePlanType = "Enterprise"
)

func PossibleValuesForCloudPCServicePlanType() []string {
	return []string{
		string(CloudPCServicePlanTypebusiness),
		string(CloudPCServicePlanTypeenterprise),
	}
}

func parseCloudPCServicePlanType(input string) (*CloudPCServicePlanType, error) {
	vals := map[string]CloudPCServicePlanType{
		"business":   CloudPCServicePlanTypebusiness,
		"enterprise": CloudPCServicePlanTypeenterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCServicePlanType(input)
	return &out, nil
}
