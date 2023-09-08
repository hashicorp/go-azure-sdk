package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementActionDeploymentStatusStatus string

const (
	ManagedTenantsManagementActionDeploymentStatusStatuscompleted                          ManagedTenantsManagementActionDeploymentStatusStatus = "Completed"
	ManagedTenantsManagementActionDeploymentStatusStatuserror                              ManagedTenantsManagementActionDeploymentStatusStatus = "Error"
	ManagedTenantsManagementActionDeploymentStatusStatusinProgress                         ManagedTenantsManagementActionDeploymentStatusStatus = "InProgress"
	ManagedTenantsManagementActionDeploymentStatusStatusplanned                            ManagedTenantsManagementActionDeploymentStatusStatus = "Planned"
	ManagedTenantsManagementActionDeploymentStatusStatusresolvedBy3rdParty                 ManagedTenantsManagementActionDeploymentStatusStatus = "ResolvedBy3rdParty"
	ManagedTenantsManagementActionDeploymentStatusStatusresolvedThroughAlternateMitigation ManagedTenantsManagementActionDeploymentStatusStatus = "ResolvedThroughAlternateMitigation"
	ManagedTenantsManagementActionDeploymentStatusStatusriskAccepted                       ManagedTenantsManagementActionDeploymentStatusStatus = "RiskAccepted"
	ManagedTenantsManagementActionDeploymentStatusStatustimeOut                            ManagedTenantsManagementActionDeploymentStatusStatus = "TimeOut"
	ManagedTenantsManagementActionDeploymentStatusStatustoAddress                          ManagedTenantsManagementActionDeploymentStatusStatus = "ToAddress"
)

func PossibleValuesForManagedTenantsManagementActionDeploymentStatusStatus() []string {
	return []string{
		string(ManagedTenantsManagementActionDeploymentStatusStatuscompleted),
		string(ManagedTenantsManagementActionDeploymentStatusStatuserror),
		string(ManagedTenantsManagementActionDeploymentStatusStatusinProgress),
		string(ManagedTenantsManagementActionDeploymentStatusStatusplanned),
		string(ManagedTenantsManagementActionDeploymentStatusStatusresolvedBy3rdParty),
		string(ManagedTenantsManagementActionDeploymentStatusStatusresolvedThroughAlternateMitigation),
		string(ManagedTenantsManagementActionDeploymentStatusStatusriskAccepted),
		string(ManagedTenantsManagementActionDeploymentStatusStatustimeOut),
		string(ManagedTenantsManagementActionDeploymentStatusStatustoAddress),
	}
}

func parseManagedTenantsManagementActionDeploymentStatusStatus(input string) (*ManagedTenantsManagementActionDeploymentStatusStatus, error) {
	vals := map[string]ManagedTenantsManagementActionDeploymentStatusStatus{
		"completed":                          ManagedTenantsManagementActionDeploymentStatusStatuscompleted,
		"error":                              ManagedTenantsManagementActionDeploymentStatusStatuserror,
		"inprogress":                         ManagedTenantsManagementActionDeploymentStatusStatusinProgress,
		"planned":                            ManagedTenantsManagementActionDeploymentStatusStatusplanned,
		"resolvedby3rdparty":                 ManagedTenantsManagementActionDeploymentStatusStatusresolvedBy3rdParty,
		"resolvedthroughalternatemitigation": ManagedTenantsManagementActionDeploymentStatusStatusresolvedThroughAlternateMitigation,
		"riskaccepted":                       ManagedTenantsManagementActionDeploymentStatusStatusriskAccepted,
		"timeout":                            ManagedTenantsManagementActionDeploymentStatusStatustimeOut,
		"toaddress":                          ManagedTenantsManagementActionDeploymentStatusStatustoAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementActionDeploymentStatusStatus(input)
	return &out, nil
}
