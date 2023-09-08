package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLogCollectionResponseStatus string

const (
	DeviceLogCollectionResponseStatuscompleted DeviceLogCollectionResponseStatus = "Completed"
	DeviceLogCollectionResponseStatusfailed    DeviceLogCollectionResponseStatus = "Failed"
	DeviceLogCollectionResponseStatuspending   DeviceLogCollectionResponseStatus = "Pending"
)

func PossibleValuesForDeviceLogCollectionResponseStatus() []string {
	return []string{
		string(DeviceLogCollectionResponseStatuscompleted),
		string(DeviceLogCollectionResponseStatusfailed),
		string(DeviceLogCollectionResponseStatuspending),
	}
}

func parseDeviceLogCollectionResponseStatus(input string) (*DeviceLogCollectionResponseStatus, error) {
	vals := map[string]DeviceLogCollectionResponseStatus{
		"completed": DeviceLogCollectionResponseStatuscompleted,
		"failed":    DeviceLogCollectionResponseStatusfailed,
		"pending":   DeviceLogCollectionResponseStatuspending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceLogCollectionResponseStatus(input)
	return &out, nil
}
