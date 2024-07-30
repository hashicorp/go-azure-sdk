package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlThreatSubmissionClientSource string

const (
	SecurityUrlThreatSubmissionClientSource_Microsoft SecurityUrlThreatSubmissionClientSource = "microsoft"
	SecurityUrlThreatSubmissionClientSource_Other     SecurityUrlThreatSubmissionClientSource = "other"
)

func PossibleValuesForSecurityUrlThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityUrlThreatSubmissionClientSource_Microsoft),
		string(SecurityUrlThreatSubmissionClientSource_Other),
	}
}

func (s *SecurityUrlThreatSubmissionClientSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUrlThreatSubmissionClientSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUrlThreatSubmissionClientSource(input string) (*SecurityUrlThreatSubmissionClientSource, error) {
	vals := map[string]SecurityUrlThreatSubmissionClientSource{
		"microsoft": SecurityUrlThreatSubmissionClientSource_Microsoft,
		"other":     SecurityUrlThreatSubmissionClientSource_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlThreatSubmissionClientSource(input)
	return &out, nil
}
