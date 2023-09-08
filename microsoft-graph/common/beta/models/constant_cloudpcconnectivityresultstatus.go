package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityResultStatus string

const (
	CloudPCConnectivityResultStatusavailable            CloudPCConnectivityResultStatus = "Available"
	CloudPCConnectivityResultStatusavailableWithWarning CloudPCConnectivityResultStatus = "AvailableWithWarning"
	CloudPCConnectivityResultStatusunavailable          CloudPCConnectivityResultStatus = "Unavailable"
	CloudPCConnectivityResultStatusunknown              CloudPCConnectivityResultStatus = "Unknown"
)

func PossibleValuesForCloudPCConnectivityResultStatus() []string {
	return []string{
		string(CloudPCConnectivityResultStatusavailable),
		string(CloudPCConnectivityResultStatusavailableWithWarning),
		string(CloudPCConnectivityResultStatusunavailable),
		string(CloudPCConnectivityResultStatusunknown),
	}
}

func parseCloudPCConnectivityResultStatus(input string) (*CloudPCConnectivityResultStatus, error) {
	vals := map[string]CloudPCConnectivityResultStatus{
		"available":            CloudPCConnectivityResultStatusavailable,
		"availablewithwarning": CloudPCConnectivityResultStatusavailableWithWarning,
		"unavailable":          CloudPCConnectivityResultStatusunavailable,
		"unknown":              CloudPCConnectivityResultStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCConnectivityResultStatus(input)
	return &out, nil
}
