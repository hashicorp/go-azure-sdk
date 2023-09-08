package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDeviceImageOsStatus string

const (
	CloudPCDeviceImageOsStatussupported            CloudPCDeviceImageOsStatus = "Supported"
	CloudPCDeviceImageOsStatussupportedWithWarning CloudPCDeviceImageOsStatus = "SupportedWithWarning"
)

func PossibleValuesForCloudPCDeviceImageOsStatus() []string {
	return []string{
		string(CloudPCDeviceImageOsStatussupported),
		string(CloudPCDeviceImageOsStatussupportedWithWarning),
	}
}

func parseCloudPCDeviceImageOsStatus(input string) (*CloudPCDeviceImageOsStatus, error) {
	vals := map[string]CloudPCDeviceImageOsStatus{
		"supported":            CloudPCDeviceImageOsStatussupported,
		"supportedwithwarning": CloudPCDeviceImageOsStatussupportedWithWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDeviceImageOsStatus(input)
	return &out, nil
}
