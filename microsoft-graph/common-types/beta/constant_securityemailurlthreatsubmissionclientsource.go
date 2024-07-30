package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionClientSource string

const (
	SecurityEmailUrlThreatSubmissionClientSource_Microsoft SecurityEmailUrlThreatSubmissionClientSource = "microsoft"
	SecurityEmailUrlThreatSubmissionClientSource_Other     SecurityEmailUrlThreatSubmissionClientSource = "other"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionClientSource_Microsoft),
		string(SecurityEmailUrlThreatSubmissionClientSource_Other),
	}
}

func (s *SecurityEmailUrlThreatSubmissionClientSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailUrlThreatSubmissionClientSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailUrlThreatSubmissionClientSource(input string) (*SecurityEmailUrlThreatSubmissionClientSource, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionClientSource{
		"microsoft": SecurityEmailUrlThreatSubmissionClientSource_Microsoft,
		"other":     SecurityEmailUrlThreatSubmissionClientSource_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionClientSource(input)
	return &out, nil
}
