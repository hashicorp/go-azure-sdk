package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSupportedRegionRegionStatus string

const (
	CloudPCSupportedRegionRegionStatusavailable   CloudPCSupportedRegionRegionStatus = "Available"
	CloudPCSupportedRegionRegionStatusrestricted  CloudPCSupportedRegionRegionStatus = "Restricted"
	CloudPCSupportedRegionRegionStatusunavailable CloudPCSupportedRegionRegionStatus = "Unavailable"
)

func PossibleValuesForCloudPCSupportedRegionRegionStatus() []string {
	return []string{
		string(CloudPCSupportedRegionRegionStatusavailable),
		string(CloudPCSupportedRegionRegionStatusrestricted),
		string(CloudPCSupportedRegionRegionStatusunavailable),
	}
}

func parseCloudPCSupportedRegionRegionStatus(input string) (*CloudPCSupportedRegionRegionStatus, error) {
	vals := map[string]CloudPCSupportedRegionRegionStatus{
		"available":   CloudPCSupportedRegionRegionStatusavailable,
		"restricted":  CloudPCSupportedRegionRegionStatusrestricted,
		"unavailable": CloudPCSupportedRegionRegionStatusunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSupportedRegionRegionStatus(input)
	return &out, nil
}
