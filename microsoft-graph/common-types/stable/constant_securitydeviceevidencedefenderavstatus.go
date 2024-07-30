package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceDefenderAvStatus string

const (
	SecurityDeviceEvidenceDefenderAvStatus_Disabled     SecurityDeviceEvidenceDefenderAvStatus = "disabled"
	SecurityDeviceEvidenceDefenderAvStatus_NotReporting SecurityDeviceEvidenceDefenderAvStatus = "notReporting"
	SecurityDeviceEvidenceDefenderAvStatus_NotSupported SecurityDeviceEvidenceDefenderAvStatus = "notSupported"
	SecurityDeviceEvidenceDefenderAvStatus_NotUpdated   SecurityDeviceEvidenceDefenderAvStatus = "notUpdated"
	SecurityDeviceEvidenceDefenderAvStatus_Unknown      SecurityDeviceEvidenceDefenderAvStatus = "unknown"
	SecurityDeviceEvidenceDefenderAvStatus_Updated      SecurityDeviceEvidenceDefenderAvStatus = "updated"
)

func PossibleValuesForSecurityDeviceEvidenceDefenderAvStatus() []string {
	return []string{
		string(SecurityDeviceEvidenceDefenderAvStatus_Disabled),
		string(SecurityDeviceEvidenceDefenderAvStatus_NotReporting),
		string(SecurityDeviceEvidenceDefenderAvStatus_NotSupported),
		string(SecurityDeviceEvidenceDefenderAvStatus_NotUpdated),
		string(SecurityDeviceEvidenceDefenderAvStatus_Unknown),
		string(SecurityDeviceEvidenceDefenderAvStatus_Updated),
	}
}

func (s *SecurityDeviceEvidenceDefenderAvStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeviceEvidenceDefenderAvStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeviceEvidenceDefenderAvStatus(input string) (*SecurityDeviceEvidenceDefenderAvStatus, error) {
	vals := map[string]SecurityDeviceEvidenceDefenderAvStatus{
		"disabled":     SecurityDeviceEvidenceDefenderAvStatus_Disabled,
		"notreporting": SecurityDeviceEvidenceDefenderAvStatus_NotReporting,
		"notsupported": SecurityDeviceEvidenceDefenderAvStatus_NotSupported,
		"notupdated":   SecurityDeviceEvidenceDefenderAvStatus_NotUpdated,
		"unknown":      SecurityDeviceEvidenceDefenderAvStatus_Unknown,
		"updated":      SecurityDeviceEvidenceDefenderAvStatus_Updated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceDefenderAvStatus(input)
	return &out, nil
}
