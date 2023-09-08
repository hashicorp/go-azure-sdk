package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityEventEventResult string

const (
	CloudPCConnectivityEventEventResultfailure CloudPCConnectivityEventEventResult = "Failure"
	CloudPCConnectivityEventEventResultsuccess CloudPCConnectivityEventEventResult = "Success"
	CloudPCConnectivityEventEventResultunknown CloudPCConnectivityEventEventResult = "Unknown"
)

func PossibleValuesForCloudPCConnectivityEventEventResult() []string {
	return []string{
		string(CloudPCConnectivityEventEventResultfailure),
		string(CloudPCConnectivityEventEventResultsuccess),
		string(CloudPCConnectivityEventEventResultunknown),
	}
}

func parseCloudPCConnectivityEventEventResult(input string) (*CloudPCConnectivityEventEventResult, error) {
	vals := map[string]CloudPCConnectivityEventEventResult{
		"failure": CloudPCConnectivityEventEventResultfailure,
		"success": CloudPCConnectivityEventEventResultsuccess,
		"unknown": CloudPCConnectivityEventEventResultunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCConnectivityEventEventResult(input)
	return &out, nil
}
