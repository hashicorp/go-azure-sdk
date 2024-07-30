package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadActionDeploymentStatusStatus string

const (
	ManagedTenantsWorkloadActionDeploymentStatusStatus_Completed  ManagedTenantsWorkloadActionDeploymentStatusStatus = "completed"
	ManagedTenantsWorkloadActionDeploymentStatusStatus_Error      ManagedTenantsWorkloadActionDeploymentStatusStatus = "error"
	ManagedTenantsWorkloadActionDeploymentStatusStatus_InProgress ManagedTenantsWorkloadActionDeploymentStatusStatus = "inProgress"
	ManagedTenantsWorkloadActionDeploymentStatusStatus_TimeOut    ManagedTenantsWorkloadActionDeploymentStatusStatus = "timeOut"
	ManagedTenantsWorkloadActionDeploymentStatusStatus_ToAddress  ManagedTenantsWorkloadActionDeploymentStatusStatus = "toAddress"
)

func PossibleValuesForManagedTenantsWorkloadActionDeploymentStatusStatus() []string {
	return []string{
		string(ManagedTenantsWorkloadActionDeploymentStatusStatus_Completed),
		string(ManagedTenantsWorkloadActionDeploymentStatusStatus_Error),
		string(ManagedTenantsWorkloadActionDeploymentStatusStatus_InProgress),
		string(ManagedTenantsWorkloadActionDeploymentStatusStatus_TimeOut),
		string(ManagedTenantsWorkloadActionDeploymentStatusStatus_ToAddress),
	}
}

func (s *ManagedTenantsWorkloadActionDeploymentStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsWorkloadActionDeploymentStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsWorkloadActionDeploymentStatusStatus(input string) (*ManagedTenantsWorkloadActionDeploymentStatusStatus, error) {
	vals := map[string]ManagedTenantsWorkloadActionDeploymentStatusStatus{
		"completed":  ManagedTenantsWorkloadActionDeploymentStatusStatus_Completed,
		"error":      ManagedTenantsWorkloadActionDeploymentStatusStatus_Error,
		"inprogress": ManagedTenantsWorkloadActionDeploymentStatusStatus_InProgress,
		"timeout":    ManagedTenantsWorkloadActionDeploymentStatusStatus_TimeOut,
		"toaddress":  ManagedTenantsWorkloadActionDeploymentStatusStatus_ToAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsWorkloadActionDeploymentStatusStatus(input)
	return &out, nil
}
