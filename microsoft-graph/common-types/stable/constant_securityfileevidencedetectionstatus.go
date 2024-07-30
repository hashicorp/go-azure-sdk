package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileEvidenceDetectionStatus string

const (
	SecurityFileEvidenceDetectionStatus_Blocked   SecurityFileEvidenceDetectionStatus = "blocked"
	SecurityFileEvidenceDetectionStatus_Detected  SecurityFileEvidenceDetectionStatus = "detected"
	SecurityFileEvidenceDetectionStatus_Prevented SecurityFileEvidenceDetectionStatus = "prevented"
)

func PossibleValuesForSecurityFileEvidenceDetectionStatus() []string {
	return []string{
		string(SecurityFileEvidenceDetectionStatus_Blocked),
		string(SecurityFileEvidenceDetectionStatus_Detected),
		string(SecurityFileEvidenceDetectionStatus_Prevented),
	}
}

func (s *SecurityFileEvidenceDetectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileEvidenceDetectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileEvidenceDetectionStatus(input string) (*SecurityFileEvidenceDetectionStatus, error) {
	vals := map[string]SecurityFileEvidenceDetectionStatus{
		"blocked":   SecurityFileEvidenceDetectionStatus_Blocked,
		"detected":  SecurityFileEvidenceDetectionStatus_Detected,
		"prevented": SecurityFileEvidenceDetectionStatus_Prevented,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileEvidenceDetectionStatus(input)
	return &out, nil
}
