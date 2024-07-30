package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileContentThreatSubmissionClientSource string

const (
	SecurityFileContentThreatSubmissionClientSource_Microsoft SecurityFileContentThreatSubmissionClientSource = "microsoft"
	SecurityFileContentThreatSubmissionClientSource_Other     SecurityFileContentThreatSubmissionClientSource = "other"
)

func PossibleValuesForSecurityFileContentThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityFileContentThreatSubmissionClientSource_Microsoft),
		string(SecurityFileContentThreatSubmissionClientSource_Other),
	}
}

func (s *SecurityFileContentThreatSubmissionClientSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileContentThreatSubmissionClientSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileContentThreatSubmissionClientSource(input string) (*SecurityFileContentThreatSubmissionClientSource, error) {
	vals := map[string]SecurityFileContentThreatSubmissionClientSource{
		"microsoft": SecurityFileContentThreatSubmissionClientSource_Microsoft,
		"other":     SecurityFileContentThreatSubmissionClientSource_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileContentThreatSubmissionClientSource(input)
	return &out, nil
}
