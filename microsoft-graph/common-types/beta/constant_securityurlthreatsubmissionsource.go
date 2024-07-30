package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlThreatSubmissionSource string

const (
	SecurityUrlThreatSubmissionSource_Administrator SecurityUrlThreatSubmissionSource = "administrator"
	SecurityUrlThreatSubmissionSource_User          SecurityUrlThreatSubmissionSource = "user"
)

func PossibleValuesForSecurityUrlThreatSubmissionSource() []string {
	return []string{
		string(SecurityUrlThreatSubmissionSource_Administrator),
		string(SecurityUrlThreatSubmissionSource_User),
	}
}

func (s *SecurityUrlThreatSubmissionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUrlThreatSubmissionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUrlThreatSubmissionSource(input string) (*SecurityUrlThreatSubmissionSource, error) {
	vals := map[string]SecurityUrlThreatSubmissionSource{
		"administrator": SecurityUrlThreatSubmissionSource_Administrator,
		"user":          SecurityUrlThreatSubmissionSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlThreatSubmissionSource(input)
	return &out, nil
}
