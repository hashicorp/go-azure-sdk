package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionClientSource string

const (
	SecurityEmailThreatSubmissionClientSource_Microsoft SecurityEmailThreatSubmissionClientSource = "microsoft"
	SecurityEmailThreatSubmissionClientSource_Other     SecurityEmailThreatSubmissionClientSource = "other"
)

func PossibleValuesForSecurityEmailThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityEmailThreatSubmissionClientSource_Microsoft),
		string(SecurityEmailThreatSubmissionClientSource_Other),
	}
}

func (s *SecurityEmailThreatSubmissionClientSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailThreatSubmissionClientSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailThreatSubmissionClientSource(input string) (*SecurityEmailThreatSubmissionClientSource, error) {
	vals := map[string]SecurityEmailThreatSubmissionClientSource{
		"microsoft": SecurityEmailThreatSubmissionClientSource_Microsoft,
		"other":     SecurityEmailThreatSubmissionClientSource_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailThreatSubmissionClientSource(input)
	return &out, nil
}
