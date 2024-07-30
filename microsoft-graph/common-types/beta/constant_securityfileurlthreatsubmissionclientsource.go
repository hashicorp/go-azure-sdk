package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileUrlThreatSubmissionClientSource string

const (
	SecurityFileUrlThreatSubmissionClientSource_Microsoft SecurityFileUrlThreatSubmissionClientSource = "microsoft"
	SecurityFileUrlThreatSubmissionClientSource_Other     SecurityFileUrlThreatSubmissionClientSource = "other"
)

func PossibleValuesForSecurityFileUrlThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityFileUrlThreatSubmissionClientSource_Microsoft),
		string(SecurityFileUrlThreatSubmissionClientSource_Other),
	}
}

func (s *SecurityFileUrlThreatSubmissionClientSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileUrlThreatSubmissionClientSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileUrlThreatSubmissionClientSource(input string) (*SecurityFileUrlThreatSubmissionClientSource, error) {
	vals := map[string]SecurityFileUrlThreatSubmissionClientSource{
		"microsoft": SecurityFileUrlThreatSubmissionClientSource_Microsoft,
		"other":     SecurityFileUrlThreatSubmissionClientSource_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileUrlThreatSubmissionClientSource(input)
	return &out, nil
}
