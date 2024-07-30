package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileThreatSubmissionClientSource string

const (
	SecurityFileThreatSubmissionClientSource_Microsoft SecurityFileThreatSubmissionClientSource = "microsoft"
	SecurityFileThreatSubmissionClientSource_Other     SecurityFileThreatSubmissionClientSource = "other"
)

func PossibleValuesForSecurityFileThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityFileThreatSubmissionClientSource_Microsoft),
		string(SecurityFileThreatSubmissionClientSource_Other),
	}
}

func (s *SecurityFileThreatSubmissionClientSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileThreatSubmissionClientSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileThreatSubmissionClientSource(input string) (*SecurityFileThreatSubmissionClientSource, error) {
	vals := map[string]SecurityFileThreatSubmissionClientSource{
		"microsoft": SecurityFileThreatSubmissionClientSource_Microsoft,
		"other":     SecurityFileThreatSubmissionClientSource_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileThreatSubmissionClientSource(input)
	return &out, nil
}
