package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSupportedRegionSupportedSolution string

const (
	CloudPCSupportedRegionSupportedSolutiondevBox     CloudPCSupportedRegionSupportedSolution = "DevBox"
	CloudPCSupportedRegionSupportedSolutionrpaBox     CloudPCSupportedRegionSupportedSolution = "RpaBox"
	CloudPCSupportedRegionSupportedSolutionwindows365 CloudPCSupportedRegionSupportedSolution = "Windows365"
)

func PossibleValuesForCloudPCSupportedRegionSupportedSolution() []string {
	return []string{
		string(CloudPCSupportedRegionSupportedSolutiondevBox),
		string(CloudPCSupportedRegionSupportedSolutionrpaBox),
		string(CloudPCSupportedRegionSupportedSolutionwindows365),
	}
}

func parseCloudPCSupportedRegionSupportedSolution(input string) (*CloudPCSupportedRegionSupportedSolution, error) {
	vals := map[string]CloudPCSupportedRegionSupportedSolution{
		"devbox":     CloudPCSupportedRegionSupportedSolutiondevBox,
		"rpabox":     CloudPCSupportedRegionSupportedSolutionrpaBox,
		"windows365": CloudPCSupportedRegionSupportedSolutionwindows365,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSupportedRegionSupportedSolution(input)
	return &out, nil
}
