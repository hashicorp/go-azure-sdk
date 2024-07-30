package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateStepDeploymentStatus string

const (
	ManagedTenantsManagementTemplateStepDeploymentStatus_Completed  ManagedTenantsManagementTemplateStepDeploymentStatus = "completed"
	ManagedTenantsManagementTemplateStepDeploymentStatus_Failed     ManagedTenantsManagementTemplateStepDeploymentStatus = "failed"
	ManagedTenantsManagementTemplateStepDeploymentStatus_InProgress ManagedTenantsManagementTemplateStepDeploymentStatus = "inProgress"
	ManagedTenantsManagementTemplateStepDeploymentStatus_Ineligible ManagedTenantsManagementTemplateStepDeploymentStatus = "ineligible"
	ManagedTenantsManagementTemplateStepDeploymentStatus_Unknown    ManagedTenantsManagementTemplateStepDeploymentStatus = "unknown"
)

func PossibleValuesForManagedTenantsManagementTemplateStepDeploymentStatus() []string {
	return []string{
		string(ManagedTenantsManagementTemplateStepDeploymentStatus_Completed),
		string(ManagedTenantsManagementTemplateStepDeploymentStatus_Failed),
		string(ManagedTenantsManagementTemplateStepDeploymentStatus_InProgress),
		string(ManagedTenantsManagementTemplateStepDeploymentStatus_Ineligible),
		string(ManagedTenantsManagementTemplateStepDeploymentStatus_Unknown),
	}
}

func (s *ManagedTenantsManagementTemplateStepDeploymentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementTemplateStepDeploymentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementTemplateStepDeploymentStatus(input string) (*ManagedTenantsManagementTemplateStepDeploymentStatus, error) {
	vals := map[string]ManagedTenantsManagementTemplateStepDeploymentStatus{
		"completed":  ManagedTenantsManagementTemplateStepDeploymentStatus_Completed,
		"failed":     ManagedTenantsManagementTemplateStepDeploymentStatus_Failed,
		"inprogress": ManagedTenantsManagementTemplateStepDeploymentStatus_InProgress,
		"ineligible": ManagedTenantsManagementTemplateStepDeploymentStatus_Ineligible,
		"unknown":    ManagedTenantsManagementTemplateStepDeploymentStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementTemplateStepDeploymentStatus(input)
	return &out, nil
}
