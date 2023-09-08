package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCServicePlanSupportedSolution string

const (
	CloudPCServicePlanSupportedSolutiondevBox     CloudPCServicePlanSupportedSolution = "DevBox"
	CloudPCServicePlanSupportedSolutionrpaBox     CloudPCServicePlanSupportedSolution = "RpaBox"
	CloudPCServicePlanSupportedSolutionwindows365 CloudPCServicePlanSupportedSolution = "Windows365"
)

func PossibleValuesForCloudPCServicePlanSupportedSolution() []string {
	return []string{
		string(CloudPCServicePlanSupportedSolutiondevBox),
		string(CloudPCServicePlanSupportedSolutionrpaBox),
		string(CloudPCServicePlanSupportedSolutionwindows365),
	}
}

func parseCloudPCServicePlanSupportedSolution(input string) (*CloudPCServicePlanSupportedSolution, error) {
	vals := map[string]CloudPCServicePlanSupportedSolution{
		"devbox":     CloudPCServicePlanSupportedSolutiondevBox,
		"rpabox":     CloudPCServicePlanSupportedSolutionrpaBox,
		"windows365": CloudPCServicePlanSupportedSolutionwindows365,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCServicePlanSupportedSolution(input)
	return &out, nil
}
