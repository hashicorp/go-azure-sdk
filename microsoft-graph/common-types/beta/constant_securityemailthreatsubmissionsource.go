package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionSource string

const (
	SecurityEmailThreatSubmissionSource_Administrator SecurityEmailThreatSubmissionSource = "administrator"
	SecurityEmailThreatSubmissionSource_User          SecurityEmailThreatSubmissionSource = "user"
)

func PossibleValuesForSecurityEmailThreatSubmissionSource() []string {
	return []string{
		string(SecurityEmailThreatSubmissionSource_Administrator),
		string(SecurityEmailThreatSubmissionSource_User),
	}
}

func (s *SecurityEmailThreatSubmissionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailThreatSubmissionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailThreatSubmissionSource(input string) (*SecurityEmailThreatSubmissionSource, error) {
	vals := map[string]SecurityEmailThreatSubmissionSource{
		"administrator": SecurityEmailThreatSubmissionSource_Administrator,
		"user":          SecurityEmailThreatSubmissionSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailThreatSubmissionSource(input)
	return &out, nil
}
