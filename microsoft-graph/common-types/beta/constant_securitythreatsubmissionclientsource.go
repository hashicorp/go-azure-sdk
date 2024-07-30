package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmissionClientSource string

const (
	SecurityThreatSubmissionClientSource_Microsoft SecurityThreatSubmissionClientSource = "microsoft"
	SecurityThreatSubmissionClientSource_Other     SecurityThreatSubmissionClientSource = "other"
)

func PossibleValuesForSecurityThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityThreatSubmissionClientSource_Microsoft),
		string(SecurityThreatSubmissionClientSource_Other),
	}
}

func (s *SecurityThreatSubmissionClientSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityThreatSubmissionClientSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityThreatSubmissionClientSource(input string) (*SecurityThreatSubmissionClientSource, error) {
	vals := map[string]SecurityThreatSubmissionClientSource{
		"microsoft": SecurityThreatSubmissionClientSource_Microsoft,
		"other":     SecurityThreatSubmissionClientSource_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityThreatSubmissionClientSource(input)
	return &out, nil
}
