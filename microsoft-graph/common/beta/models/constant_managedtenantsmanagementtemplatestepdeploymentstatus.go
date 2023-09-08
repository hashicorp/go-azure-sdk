package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateStepDeploymentStatus string

const (
	ManagedTenantsManagementTemplateStepDeploymentStatuscompleted  ManagedTenantsManagementTemplateStepDeploymentStatus = "Completed"
	ManagedTenantsManagementTemplateStepDeploymentStatusfailed     ManagedTenantsManagementTemplateStepDeploymentStatus = "Failed"
	ManagedTenantsManagementTemplateStepDeploymentStatusinProgress ManagedTenantsManagementTemplateStepDeploymentStatus = "InProgress"
	ManagedTenantsManagementTemplateStepDeploymentStatusineligible ManagedTenantsManagementTemplateStepDeploymentStatus = "Ineligible"
	ManagedTenantsManagementTemplateStepDeploymentStatusunknown    ManagedTenantsManagementTemplateStepDeploymentStatus = "Unknown"
)

func PossibleValuesForManagedTenantsManagementTemplateStepDeploymentStatus() []string {
	return []string{
		string(ManagedTenantsManagementTemplateStepDeploymentStatuscompleted),
		string(ManagedTenantsManagementTemplateStepDeploymentStatusfailed),
		string(ManagedTenantsManagementTemplateStepDeploymentStatusinProgress),
		string(ManagedTenantsManagementTemplateStepDeploymentStatusineligible),
		string(ManagedTenantsManagementTemplateStepDeploymentStatusunknown),
	}
}

func parseManagedTenantsManagementTemplateStepDeploymentStatus(input string) (*ManagedTenantsManagementTemplateStepDeploymentStatus, error) {
	vals := map[string]ManagedTenantsManagementTemplateStepDeploymentStatus{
		"completed":  ManagedTenantsManagementTemplateStepDeploymentStatuscompleted,
		"failed":     ManagedTenantsManagementTemplateStepDeploymentStatusfailed,
		"inprogress": ManagedTenantsManagementTemplateStepDeploymentStatusinProgress,
		"ineligible": ManagedTenantsManagementTemplateStepDeploymentStatusineligible,
		"unknown":    ManagedTenantsManagementTemplateStepDeploymentStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementTemplateStepDeploymentStatus(input)
	return &out, nil
}
