package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionSource string

const (
	SecurityEmailContentThreatSubmissionSource_Administrator SecurityEmailContentThreatSubmissionSource = "administrator"
	SecurityEmailContentThreatSubmissionSource_User          SecurityEmailContentThreatSubmissionSource = "user"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionSource() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionSource_Administrator),
		string(SecurityEmailContentThreatSubmissionSource_User),
	}
}

func (s *SecurityEmailContentThreatSubmissionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailContentThreatSubmissionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailContentThreatSubmissionSource(input string) (*SecurityEmailContentThreatSubmissionSource, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionSource{
		"administrator": SecurityEmailContentThreatSubmissionSource_Administrator,
		"user":          SecurityEmailContentThreatSubmissionSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionSource(input)
	return &out, nil
}
