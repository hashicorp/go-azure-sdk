package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCProvisioningPolicyImageType string

const (
	CloudPCProvisioningPolicyImageTypecustom  CloudPCProvisioningPolicyImageType = "Custom"
	CloudPCProvisioningPolicyImageTypegallery CloudPCProvisioningPolicyImageType = "Gallery"
)

func PossibleValuesForCloudPCProvisioningPolicyImageType() []string {
	return []string{
		string(CloudPCProvisioningPolicyImageTypecustom),
		string(CloudPCProvisioningPolicyImageTypegallery),
	}
}

func parseCloudPCProvisioningPolicyImageType(input string) (*CloudPCProvisioningPolicyImageType, error) {
	vals := map[string]CloudPCProvisioningPolicyImageType{
		"custom":  CloudPCProvisioningPolicyImageTypecustom,
		"gallery": CloudPCProvisioningPolicyImageTypegallery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCProvisioningPolicyImageType(input)
	return &out, nil
}
