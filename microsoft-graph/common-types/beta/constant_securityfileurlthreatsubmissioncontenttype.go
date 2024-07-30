package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileUrlThreatSubmissionContentType string

const (
	SecurityFileUrlThreatSubmissionContentType_App   SecurityFileUrlThreatSubmissionContentType = "app"
	SecurityFileUrlThreatSubmissionContentType_Email SecurityFileUrlThreatSubmissionContentType = "email"
	SecurityFileUrlThreatSubmissionContentType_File  SecurityFileUrlThreatSubmissionContentType = "file"
	SecurityFileUrlThreatSubmissionContentType_Url   SecurityFileUrlThreatSubmissionContentType = "url"
)

func PossibleValuesForSecurityFileUrlThreatSubmissionContentType() []string {
	return []string{
		string(SecurityFileUrlThreatSubmissionContentType_App),
		string(SecurityFileUrlThreatSubmissionContentType_Email),
		string(SecurityFileUrlThreatSubmissionContentType_File),
		string(SecurityFileUrlThreatSubmissionContentType_Url),
	}
}

func (s *SecurityFileUrlThreatSubmissionContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileUrlThreatSubmissionContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileUrlThreatSubmissionContentType(input string) (*SecurityFileUrlThreatSubmissionContentType, error) {
	vals := map[string]SecurityFileUrlThreatSubmissionContentType{
		"app":   SecurityFileUrlThreatSubmissionContentType_App,
		"email": SecurityFileUrlThreatSubmissionContentType_Email,
		"file":  SecurityFileUrlThreatSubmissionContentType_File,
		"url":   SecurityFileUrlThreatSubmissionContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileUrlThreatSubmissionContentType(input)
	return &out, nil
}
