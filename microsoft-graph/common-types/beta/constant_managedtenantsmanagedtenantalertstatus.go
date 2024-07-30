package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlertStatus string

const (
	ManagedTenantsManagedTenantAlertStatus_Dismissed  ManagedTenantsManagedTenantAlertStatus = "dismissed"
	ManagedTenantsManagedTenantAlertStatus_InProgress ManagedTenantsManagedTenantAlertStatus = "inProgress"
	ManagedTenantsManagedTenantAlertStatus_NewAlert   ManagedTenantsManagedTenantAlertStatus = "newAlert"
	ManagedTenantsManagedTenantAlertStatus_Resolved   ManagedTenantsManagedTenantAlertStatus = "resolved"
	ManagedTenantsManagedTenantAlertStatus_Unknown    ManagedTenantsManagedTenantAlertStatus = "unknown"
)

func PossibleValuesForManagedTenantsManagedTenantAlertStatus() []string {
	return []string{
		string(ManagedTenantsManagedTenantAlertStatus_Dismissed),
		string(ManagedTenantsManagedTenantAlertStatus_InProgress),
		string(ManagedTenantsManagedTenantAlertStatus_NewAlert),
		string(ManagedTenantsManagedTenantAlertStatus_Resolved),
		string(ManagedTenantsManagedTenantAlertStatus_Unknown),
	}
}

func (s *ManagedTenantsManagedTenantAlertStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagedTenantAlertStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagedTenantAlertStatus(input string) (*ManagedTenantsManagedTenantAlertStatus, error) {
	vals := map[string]ManagedTenantsManagedTenantAlertStatus{
		"dismissed":  ManagedTenantsManagedTenantAlertStatus_Dismissed,
		"inprogress": ManagedTenantsManagedTenantAlertStatus_InProgress,
		"newalert":   ManagedTenantsManagedTenantAlertStatus_NewAlert,
		"resolved":   ManagedTenantsManagedTenantAlertStatus_Resolved,
		"unknown":    ManagedTenantsManagedTenantAlertStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagedTenantAlertStatus(input)
	return &out, nil
}
