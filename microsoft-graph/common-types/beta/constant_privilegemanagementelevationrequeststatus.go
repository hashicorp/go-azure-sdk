package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementElevationRequestStatus string

const (
	PrivilegeManagementElevationRequestStatus_Approved PrivilegeManagementElevationRequestStatus = "approved"
	PrivilegeManagementElevationRequestStatus_Denied   PrivilegeManagementElevationRequestStatus = "denied"
	PrivilegeManagementElevationRequestStatus_Expired  PrivilegeManagementElevationRequestStatus = "expired"
	PrivilegeManagementElevationRequestStatus_None     PrivilegeManagementElevationRequestStatus = "none"
	PrivilegeManagementElevationRequestStatus_Pending  PrivilegeManagementElevationRequestStatus = "pending"
)

func PossibleValuesForPrivilegeManagementElevationRequestStatus() []string {
	return []string{
		string(PrivilegeManagementElevationRequestStatus_Approved),
		string(PrivilegeManagementElevationRequestStatus_Denied),
		string(PrivilegeManagementElevationRequestStatus_Expired),
		string(PrivilegeManagementElevationRequestStatus_None),
		string(PrivilegeManagementElevationRequestStatus_Pending),
	}
}

func (s *PrivilegeManagementElevationRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegeManagementElevationRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegeManagementElevationRequestStatus(input string) (*PrivilegeManagementElevationRequestStatus, error) {
	vals := map[string]PrivilegeManagementElevationRequestStatus{
		"approved": PrivilegeManagementElevationRequestStatus_Approved,
		"denied":   PrivilegeManagementElevationRequestStatus_Denied,
		"expired":  PrivilegeManagementElevationRequestStatus_Expired,
		"none":     PrivilegeManagementElevationRequestStatus_None,
		"pending":  PrivilegeManagementElevationRequestStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegeManagementElevationRequestStatus(input)
	return &out, nil
}
