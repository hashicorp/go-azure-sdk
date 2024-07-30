package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None    AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForAndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude),
		string(AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include),
		string(AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None),
	}
}

func (s *AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude,
		"include": AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include,
		"none":    AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
