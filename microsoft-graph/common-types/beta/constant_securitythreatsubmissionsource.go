package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmissionSource string

const (
	SecurityThreatSubmissionSource_Administrator SecurityThreatSubmissionSource = "administrator"
	SecurityThreatSubmissionSource_User          SecurityThreatSubmissionSource = "user"
)

func PossibleValuesForSecurityThreatSubmissionSource() []string {
	return []string{
		string(SecurityThreatSubmissionSource_Administrator),
		string(SecurityThreatSubmissionSource_User),
	}
}

func (s *SecurityThreatSubmissionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityThreatSubmissionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityThreatSubmissionSource(input string) (*SecurityThreatSubmissionSource, error) {
	vals := map[string]SecurityThreatSubmissionSource{
		"administrator": SecurityThreatSubmissionSource_Administrator,
		"user":          SecurityThreatSubmissionSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityThreatSubmissionSource(input)
	return &out, nil
}
