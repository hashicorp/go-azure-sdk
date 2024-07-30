package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None    AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForAllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude),
		string(AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include),
		string(AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None),
	}
}

func (s *AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude,
		"include": AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include,
		"none":    AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
