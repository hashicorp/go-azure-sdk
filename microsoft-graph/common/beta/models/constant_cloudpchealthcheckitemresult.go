package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCHealthCheckItemResult string

const (
	CloudPCHealthCheckItemResultfailure CloudPCHealthCheckItemResult = "Failure"
	CloudPCHealthCheckItemResultsuccess CloudPCHealthCheckItemResult = "Success"
	CloudPCHealthCheckItemResultunknown CloudPCHealthCheckItemResult = "Unknown"
)

func PossibleValuesForCloudPCHealthCheckItemResult() []string {
	return []string{
		string(CloudPCHealthCheckItemResultfailure),
		string(CloudPCHealthCheckItemResultsuccess),
		string(CloudPCHealthCheckItemResultunknown),
	}
}

func parseCloudPCHealthCheckItemResult(input string) (*CloudPCHealthCheckItemResult, error) {
	vals := map[string]CloudPCHealthCheckItemResult{
		"failure": CloudPCHealthCheckItemResultfailure,
		"success": CloudPCHealthCheckItemResultsuccess,
		"unknown": CloudPCHealthCheckItemResultunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCHealthCheckItemResult(input)
	return &out, nil
}
