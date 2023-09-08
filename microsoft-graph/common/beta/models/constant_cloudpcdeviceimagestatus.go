package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDeviceImageStatus string

const (
	CloudPCDeviceImageStatusfailed  CloudPCDeviceImageStatus = "Failed"
	CloudPCDeviceImageStatuspending CloudPCDeviceImageStatus = "Pending"
	CloudPCDeviceImageStatusready   CloudPCDeviceImageStatus = "Ready"
)

func PossibleValuesForCloudPCDeviceImageStatus() []string {
	return []string{
		string(CloudPCDeviceImageStatusfailed),
		string(CloudPCDeviceImageStatuspending),
		string(CloudPCDeviceImageStatusready),
	}
}

func parseCloudPCDeviceImageStatus(input string) (*CloudPCDeviceImageStatus, error) {
	vals := map[string]CloudPCDeviceImageStatus{
		"failed":  CloudPCDeviceImageStatusfailed,
		"pending": CloudPCDeviceImageStatuspending,
		"ready":   CloudPCDeviceImageStatusready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDeviceImageStatus(input)
	return &out, nil
}
