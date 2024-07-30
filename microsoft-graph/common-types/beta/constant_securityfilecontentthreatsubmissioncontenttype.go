package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileContentThreatSubmissionContentType string

const (
	SecurityFileContentThreatSubmissionContentType_App   SecurityFileContentThreatSubmissionContentType = "app"
	SecurityFileContentThreatSubmissionContentType_Email SecurityFileContentThreatSubmissionContentType = "email"
	SecurityFileContentThreatSubmissionContentType_File  SecurityFileContentThreatSubmissionContentType = "file"
	SecurityFileContentThreatSubmissionContentType_Url   SecurityFileContentThreatSubmissionContentType = "url"
)

func PossibleValuesForSecurityFileContentThreatSubmissionContentType() []string {
	return []string{
		string(SecurityFileContentThreatSubmissionContentType_App),
		string(SecurityFileContentThreatSubmissionContentType_Email),
		string(SecurityFileContentThreatSubmissionContentType_File),
		string(SecurityFileContentThreatSubmissionContentType_Url),
	}
}

func (s *SecurityFileContentThreatSubmissionContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileContentThreatSubmissionContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileContentThreatSubmissionContentType(input string) (*SecurityFileContentThreatSubmissionContentType, error) {
	vals := map[string]SecurityFileContentThreatSubmissionContentType{
		"app":   SecurityFileContentThreatSubmissionContentType_App,
		"email": SecurityFileContentThreatSubmissionContentType_Email,
		"file":  SecurityFileContentThreatSubmissionContentType_File,
		"url":   SecurityFileContentThreatSubmissionContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileContentThreatSubmissionContentType(input)
	return &out, nil
}
