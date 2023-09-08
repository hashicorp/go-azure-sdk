package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionManagedBy string

const (
	CloudPCOnPremisesConnectionManagedBydevBox     CloudPCOnPremisesConnectionManagedBy = "DevBox"
	CloudPCOnPremisesConnectionManagedByrpaBox     CloudPCOnPremisesConnectionManagedBy = "RpaBox"
	CloudPCOnPremisesConnectionManagedBywindows365 CloudPCOnPremisesConnectionManagedBy = "Windows365"
)

func PossibleValuesForCloudPCOnPremisesConnectionManagedBy() []string {
	return []string{
		string(CloudPCOnPremisesConnectionManagedBydevBox),
		string(CloudPCOnPremisesConnectionManagedByrpaBox),
		string(CloudPCOnPremisesConnectionManagedBywindows365),
	}
}

func parseCloudPCOnPremisesConnectionManagedBy(input string) (*CloudPCOnPremisesConnectionManagedBy, error) {
	vals := map[string]CloudPCOnPremisesConnectionManagedBy{
		"devbox":     CloudPCOnPremisesConnectionManagedBydevBox,
		"rpabox":     CloudPCOnPremisesConnectionManagedByrpaBox,
		"windows365": CloudPCOnPremisesConnectionManagedBywindows365,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOnPremisesConnectionManagedBy(input)
	return &out, nil
}
