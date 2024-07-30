package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProcessEvidenceDetectionStatus string

const (
	SecurityProcessEvidenceDetectionStatus_Blocked   SecurityProcessEvidenceDetectionStatus = "blocked"
	SecurityProcessEvidenceDetectionStatus_Detected  SecurityProcessEvidenceDetectionStatus = "detected"
	SecurityProcessEvidenceDetectionStatus_Prevented SecurityProcessEvidenceDetectionStatus = "prevented"
)

func PossibleValuesForSecurityProcessEvidenceDetectionStatus() []string {
	return []string{
		string(SecurityProcessEvidenceDetectionStatus_Blocked),
		string(SecurityProcessEvidenceDetectionStatus_Detected),
		string(SecurityProcessEvidenceDetectionStatus_Prevented),
	}
}

func (s *SecurityProcessEvidenceDetectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityProcessEvidenceDetectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityProcessEvidenceDetectionStatus(input string) (*SecurityProcessEvidenceDetectionStatus, error) {
	vals := map[string]SecurityProcessEvidenceDetectionStatus{
		"blocked":   SecurityProcessEvidenceDetectionStatus_Blocked,
		"detected":  SecurityProcessEvidenceDetectionStatus_Detected,
		"prevented": SecurityProcessEvidenceDetectionStatus_Prevented,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityProcessEvidenceDetectionStatus(input)
	return &out, nil
}
