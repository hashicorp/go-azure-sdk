package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileThreatSubmissionContentType string

const (
	SecurityFileThreatSubmissionContentType_App   SecurityFileThreatSubmissionContentType = "app"
	SecurityFileThreatSubmissionContentType_Email SecurityFileThreatSubmissionContentType = "email"
	SecurityFileThreatSubmissionContentType_File  SecurityFileThreatSubmissionContentType = "file"
	SecurityFileThreatSubmissionContentType_Url   SecurityFileThreatSubmissionContentType = "url"
)

func PossibleValuesForSecurityFileThreatSubmissionContentType() []string {
	return []string{
		string(SecurityFileThreatSubmissionContentType_App),
		string(SecurityFileThreatSubmissionContentType_Email),
		string(SecurityFileThreatSubmissionContentType_File),
		string(SecurityFileThreatSubmissionContentType_Url),
	}
}

func (s *SecurityFileThreatSubmissionContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileThreatSubmissionContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileThreatSubmissionContentType(input string) (*SecurityFileThreatSubmissionContentType, error) {
	vals := map[string]SecurityFileThreatSubmissionContentType{
		"app":   SecurityFileThreatSubmissionContentType_App,
		"email": SecurityFileThreatSubmissionContentType_Email,
		"file":  SecurityFileThreatSubmissionContentType_File,
		"url":   SecurityFileThreatSubmissionContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileThreatSubmissionContentType(input)
	return &out, nil
}
