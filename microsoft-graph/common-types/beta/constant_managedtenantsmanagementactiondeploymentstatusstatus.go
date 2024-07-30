package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementActionDeploymentStatusStatus string

const (
	ManagedTenantsManagementActionDeploymentStatusStatus_Completed                          ManagedTenantsManagementActionDeploymentStatusStatus = "completed"
	ManagedTenantsManagementActionDeploymentStatusStatus_Error                              ManagedTenantsManagementActionDeploymentStatusStatus = "error"
	ManagedTenantsManagementActionDeploymentStatusStatus_InProgress                         ManagedTenantsManagementActionDeploymentStatusStatus = "inProgress"
	ManagedTenantsManagementActionDeploymentStatusStatus_Planned                            ManagedTenantsManagementActionDeploymentStatusStatus = "planned"
	ManagedTenantsManagementActionDeploymentStatusStatus_ResolvedBy3rdParty                 ManagedTenantsManagementActionDeploymentStatusStatus = "resolvedBy3rdParty"
	ManagedTenantsManagementActionDeploymentStatusStatus_ResolvedThroughAlternateMitigation ManagedTenantsManagementActionDeploymentStatusStatus = "resolvedThroughAlternateMitigation"
	ManagedTenantsManagementActionDeploymentStatusStatus_RiskAccepted                       ManagedTenantsManagementActionDeploymentStatusStatus = "riskAccepted"
	ManagedTenantsManagementActionDeploymentStatusStatus_TimeOut                            ManagedTenantsManagementActionDeploymentStatusStatus = "timeOut"
	ManagedTenantsManagementActionDeploymentStatusStatus_ToAddress                          ManagedTenantsManagementActionDeploymentStatusStatus = "toAddress"
)

func PossibleValuesForManagedTenantsManagementActionDeploymentStatusStatus() []string {
	return []string{
		string(ManagedTenantsManagementActionDeploymentStatusStatus_Completed),
		string(ManagedTenantsManagementActionDeploymentStatusStatus_Error),
		string(ManagedTenantsManagementActionDeploymentStatusStatus_InProgress),
		string(ManagedTenantsManagementActionDeploymentStatusStatus_Planned),
		string(ManagedTenantsManagementActionDeploymentStatusStatus_ResolvedBy3rdParty),
		string(ManagedTenantsManagementActionDeploymentStatusStatus_ResolvedThroughAlternateMitigation),
		string(ManagedTenantsManagementActionDeploymentStatusStatus_RiskAccepted),
		string(ManagedTenantsManagementActionDeploymentStatusStatus_TimeOut),
		string(ManagedTenantsManagementActionDeploymentStatusStatus_ToAddress),
	}
}

func (s *ManagedTenantsManagementActionDeploymentStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementActionDeploymentStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementActionDeploymentStatusStatus(input string) (*ManagedTenantsManagementActionDeploymentStatusStatus, error) {
	vals := map[string]ManagedTenantsManagementActionDeploymentStatusStatus{
		"completed":                          ManagedTenantsManagementActionDeploymentStatusStatus_Completed,
		"error":                              ManagedTenantsManagementActionDeploymentStatusStatus_Error,
		"inprogress":                         ManagedTenantsManagementActionDeploymentStatusStatus_InProgress,
		"planned":                            ManagedTenantsManagementActionDeploymentStatusStatus_Planned,
		"resolvedby3rdparty":                 ManagedTenantsManagementActionDeploymentStatusStatus_ResolvedBy3rdParty,
		"resolvedthroughalternatemitigation": ManagedTenantsManagementActionDeploymentStatusStatus_ResolvedThroughAlternateMitigation,
		"riskaccepted":                       ManagedTenantsManagementActionDeploymentStatusStatus_RiskAccepted,
		"timeout":                            ManagedTenantsManagementActionDeploymentStatusStatus_TimeOut,
		"toaddress":                          ManagedTenantsManagementActionDeploymentStatusStatus_ToAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementActionDeploymentStatusStatus(input)
	return &out, nil
}
