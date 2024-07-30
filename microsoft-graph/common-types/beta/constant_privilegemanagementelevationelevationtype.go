package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementElevationElevationType string

const (
	PrivilegeManagementElevationElevationType_SupportApprovedElevation PrivilegeManagementElevationElevationType = "supportApprovedElevation"
	PrivilegeManagementElevationElevationType_Undetermined             PrivilegeManagementElevationElevationType = "undetermined"
	PrivilegeManagementElevationElevationType_UnmanagedElevation       PrivilegeManagementElevationElevationType = "unmanagedElevation"
	PrivilegeManagementElevationElevationType_UserConfirmedElevation   PrivilegeManagementElevationElevationType = "userConfirmedElevation"
	PrivilegeManagementElevationElevationType_ZeroTouchElevation       PrivilegeManagementElevationElevationType = "zeroTouchElevation"
)

func PossibleValuesForPrivilegeManagementElevationElevationType() []string {
	return []string{
		string(PrivilegeManagementElevationElevationType_SupportApprovedElevation),
		string(PrivilegeManagementElevationElevationType_Undetermined),
		string(PrivilegeManagementElevationElevationType_UnmanagedElevation),
		string(PrivilegeManagementElevationElevationType_UserConfirmedElevation),
		string(PrivilegeManagementElevationElevationType_ZeroTouchElevation),
	}
}

func (s *PrivilegeManagementElevationElevationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegeManagementElevationElevationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegeManagementElevationElevationType(input string) (*PrivilegeManagementElevationElevationType, error) {
	vals := map[string]PrivilegeManagementElevationElevationType{
		"supportapprovedelevation": PrivilegeManagementElevationElevationType_SupportApprovedElevation,
		"undetermined":             PrivilegeManagementElevationElevationType_Undetermined,
		"unmanagedelevation":       PrivilegeManagementElevationElevationType_UnmanagedElevation,
		"userconfirmedelevation":   PrivilegeManagementElevationElevationType_UserConfirmedElevation,
		"zerotouchelevation":       PrivilegeManagementElevationElevationType_ZeroTouchElevation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegeManagementElevationElevationType(input)
	return &out, nil
}
