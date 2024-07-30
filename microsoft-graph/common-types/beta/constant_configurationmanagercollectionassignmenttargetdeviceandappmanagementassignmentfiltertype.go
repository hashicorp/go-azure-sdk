package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None    ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude),
		string(ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include),
		string(ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None),
	}
}

func (s *ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude,
		"include": ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include,
		"none":    ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
