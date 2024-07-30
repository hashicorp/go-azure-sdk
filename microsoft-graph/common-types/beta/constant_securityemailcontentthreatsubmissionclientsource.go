package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionClientSource string

const (
	SecurityEmailContentThreatSubmissionClientSource_Microsoft SecurityEmailContentThreatSubmissionClientSource = "microsoft"
	SecurityEmailContentThreatSubmissionClientSource_Other     SecurityEmailContentThreatSubmissionClientSource = "other"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionClientSource_Microsoft),
		string(SecurityEmailContentThreatSubmissionClientSource_Other),
	}
}

func (s *SecurityEmailContentThreatSubmissionClientSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailContentThreatSubmissionClientSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailContentThreatSubmissionClientSource(input string) (*SecurityEmailContentThreatSubmissionClientSource, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionClientSource{
		"microsoft": SecurityEmailContentThreatSubmissionClientSource_Microsoft,
		"other":     SecurityEmailContentThreatSubmissionClientSource_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionClientSource(input)
	return &out, nil
}
