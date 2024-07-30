package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionHealthCheckStatus string

const (
	CloudPCOnPremisesConnectionHealthCheckStatus_Failed        CloudPCOnPremisesConnectionHealthCheckStatus = "failed"
	CloudPCOnPremisesConnectionHealthCheckStatus_Informational CloudPCOnPremisesConnectionHealthCheckStatus = "informational"
	CloudPCOnPremisesConnectionHealthCheckStatus_Passed        CloudPCOnPremisesConnectionHealthCheckStatus = "passed"
	CloudPCOnPremisesConnectionHealthCheckStatus_Pending       CloudPCOnPremisesConnectionHealthCheckStatus = "pending"
	CloudPCOnPremisesConnectionHealthCheckStatus_Running       CloudPCOnPremisesConnectionHealthCheckStatus = "running"
	CloudPCOnPremisesConnectionHealthCheckStatus_Warning       CloudPCOnPremisesConnectionHealthCheckStatus = "warning"
)

func PossibleValuesForCloudPCOnPremisesConnectionHealthCheckStatus() []string {
	return []string{
		string(CloudPCOnPremisesConnectionHealthCheckStatus_Failed),
		string(CloudPCOnPremisesConnectionHealthCheckStatus_Informational),
		string(CloudPCOnPremisesConnectionHealthCheckStatus_Passed),
		string(CloudPCOnPremisesConnectionHealthCheckStatus_Pending),
		string(CloudPCOnPremisesConnectionHealthCheckStatus_Running),
		string(CloudPCOnPremisesConnectionHealthCheckStatus_Warning),
	}
}

func (s *CloudPCOnPremisesConnectionHealthCheckStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCOnPremisesConnectionHealthCheckStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCOnPremisesConnectionHealthCheckStatus(input string) (*CloudPCOnPremisesConnectionHealthCheckStatus, error) {
	vals := map[string]CloudPCOnPremisesConnectionHealthCheckStatus{
		"failed":        CloudPCOnPremisesConnectionHealthCheckStatus_Failed,
		"informational": CloudPCOnPremisesConnectionHealthCheckStatus_Informational,
		"passed":        CloudPCOnPremisesConnectionHealthCheckStatus_Passed,
		"pending":       CloudPCOnPremisesConnectionHealthCheckStatus_Pending,
		"running":       CloudPCOnPremisesConnectionHealthCheckStatus_Running,
		"warning":       CloudPCOnPremisesConnectionHealthCheckStatus_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOnPremisesConnectionHealthCheckStatus(input)
	return &out, nil
}
