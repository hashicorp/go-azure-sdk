package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadActionDeploymentStatusStatus string

const (
	ManagedTenantsWorkloadActionDeploymentStatusStatuscompleted  ManagedTenantsWorkloadActionDeploymentStatusStatus = "Completed"
	ManagedTenantsWorkloadActionDeploymentStatusStatuserror      ManagedTenantsWorkloadActionDeploymentStatusStatus = "Error"
	ManagedTenantsWorkloadActionDeploymentStatusStatusinProgress ManagedTenantsWorkloadActionDeploymentStatusStatus = "InProgress"
	ManagedTenantsWorkloadActionDeploymentStatusStatustimeOut    ManagedTenantsWorkloadActionDeploymentStatusStatus = "TimeOut"
	ManagedTenantsWorkloadActionDeploymentStatusStatustoAddress  ManagedTenantsWorkloadActionDeploymentStatusStatus = "ToAddress"
)

func PossibleValuesForManagedTenantsWorkloadActionDeploymentStatusStatus() []string {
	return []string{
		string(ManagedTenantsWorkloadActionDeploymentStatusStatuscompleted),
		string(ManagedTenantsWorkloadActionDeploymentStatusStatuserror),
		string(ManagedTenantsWorkloadActionDeploymentStatusStatusinProgress),
		string(ManagedTenantsWorkloadActionDeploymentStatusStatustimeOut),
		string(ManagedTenantsWorkloadActionDeploymentStatusStatustoAddress),
	}
}

func parseManagedTenantsWorkloadActionDeploymentStatusStatus(input string) (*ManagedTenantsWorkloadActionDeploymentStatusStatus, error) {
	vals := map[string]ManagedTenantsWorkloadActionDeploymentStatusStatus{
		"completed":  ManagedTenantsWorkloadActionDeploymentStatusStatuscompleted,
		"error":      ManagedTenantsWorkloadActionDeploymentStatusStatuserror,
		"inprogress": ManagedTenantsWorkloadActionDeploymentStatusStatusinProgress,
		"timeout":    ManagedTenantsWorkloadActionDeploymentStatusStatustimeOut,
		"toaddress":  ManagedTenantsWorkloadActionDeploymentStatusStatustoAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsWorkloadActionDeploymentStatusStatus(input)
	return &out, nil
}
