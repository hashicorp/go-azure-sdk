package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionSource string

const (
	SecurityEmailUrlThreatSubmissionSource_Administrator SecurityEmailUrlThreatSubmissionSource = "administrator"
	SecurityEmailUrlThreatSubmissionSource_User          SecurityEmailUrlThreatSubmissionSource = "user"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionSource() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionSource_Administrator),
		string(SecurityEmailUrlThreatSubmissionSource_User),
	}
}

func (s *SecurityEmailUrlThreatSubmissionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailUrlThreatSubmissionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailUrlThreatSubmissionSource(input string) (*SecurityEmailUrlThreatSubmissionSource, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionSource{
		"administrator": SecurityEmailUrlThreatSubmissionSource_Administrator,
		"user":          SecurityEmailUrlThreatSubmissionSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionSource(input)
	return &out, nil
}
