package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionHealthCheckStatus string

const (
	CloudPCOnPremisesConnectionHealthCheckStatusfailed        CloudPCOnPremisesConnectionHealthCheckStatus = "Failed"
	CloudPCOnPremisesConnectionHealthCheckStatusinformational CloudPCOnPremisesConnectionHealthCheckStatus = "Informational"
	CloudPCOnPremisesConnectionHealthCheckStatuspassed        CloudPCOnPremisesConnectionHealthCheckStatus = "Passed"
	CloudPCOnPremisesConnectionHealthCheckStatuspending       CloudPCOnPremisesConnectionHealthCheckStatus = "Pending"
	CloudPCOnPremisesConnectionHealthCheckStatusrunning       CloudPCOnPremisesConnectionHealthCheckStatus = "Running"
	CloudPCOnPremisesConnectionHealthCheckStatuswarning       CloudPCOnPremisesConnectionHealthCheckStatus = "Warning"
)

func PossibleValuesForCloudPCOnPremisesConnectionHealthCheckStatus() []string {
	return []string{
		string(CloudPCOnPremisesConnectionHealthCheckStatusfailed),
		string(CloudPCOnPremisesConnectionHealthCheckStatusinformational),
		string(CloudPCOnPremisesConnectionHealthCheckStatuspassed),
		string(CloudPCOnPremisesConnectionHealthCheckStatuspending),
		string(CloudPCOnPremisesConnectionHealthCheckStatusrunning),
		string(CloudPCOnPremisesConnectionHealthCheckStatuswarning),
	}
}

func parseCloudPCOnPremisesConnectionHealthCheckStatus(input string) (*CloudPCOnPremisesConnectionHealthCheckStatus, error) {
	vals := map[string]CloudPCOnPremisesConnectionHealthCheckStatus{
		"failed":        CloudPCOnPremisesConnectionHealthCheckStatusfailed,
		"informational": CloudPCOnPremisesConnectionHealthCheckStatusinformational,
		"passed":        CloudPCOnPremisesConnectionHealthCheckStatuspassed,
		"pending":       CloudPCOnPremisesConnectionHealthCheckStatuspending,
		"running":       CloudPCOnPremisesConnectionHealthCheckStatusrunning,
		"warning":       CloudPCOnPremisesConnectionHealthCheckStatuswarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOnPremisesConnectionHealthCheckStatus(input)
	return &out, nil
}
