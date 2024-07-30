package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileContentThreatSubmissionSource string

const (
	SecurityFileContentThreatSubmissionSource_Administrator SecurityFileContentThreatSubmissionSource = "administrator"
	SecurityFileContentThreatSubmissionSource_User          SecurityFileContentThreatSubmissionSource = "user"
)

func PossibleValuesForSecurityFileContentThreatSubmissionSource() []string {
	return []string{
		string(SecurityFileContentThreatSubmissionSource_Administrator),
		string(SecurityFileContentThreatSubmissionSource_User),
	}
}

func (s *SecurityFileContentThreatSubmissionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileContentThreatSubmissionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileContentThreatSubmissionSource(input string) (*SecurityFileContentThreatSubmissionSource, error) {
	vals := map[string]SecurityFileContentThreatSubmissionSource{
		"administrator": SecurityFileContentThreatSubmissionSource_Administrator,
		"user":          SecurityFileContentThreatSubmissionSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileContentThreatSubmissionSource(input)
	return &out, nil
}
