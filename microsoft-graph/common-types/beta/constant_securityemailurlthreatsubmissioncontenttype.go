package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionContentType string

const (
	SecurityEmailUrlThreatSubmissionContentType_App   SecurityEmailUrlThreatSubmissionContentType = "app"
	SecurityEmailUrlThreatSubmissionContentType_Email SecurityEmailUrlThreatSubmissionContentType = "email"
	SecurityEmailUrlThreatSubmissionContentType_File  SecurityEmailUrlThreatSubmissionContentType = "file"
	SecurityEmailUrlThreatSubmissionContentType_Url   SecurityEmailUrlThreatSubmissionContentType = "url"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionContentType() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionContentType_App),
		string(SecurityEmailUrlThreatSubmissionContentType_Email),
		string(SecurityEmailUrlThreatSubmissionContentType_File),
		string(SecurityEmailUrlThreatSubmissionContentType_Url),
	}
}

func (s *SecurityEmailUrlThreatSubmissionContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailUrlThreatSubmissionContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailUrlThreatSubmissionContentType(input string) (*SecurityEmailUrlThreatSubmissionContentType, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionContentType{
		"app":   SecurityEmailUrlThreatSubmissionContentType_App,
		"email": SecurityEmailUrlThreatSubmissionContentType_Email,
		"file":  SecurityEmailUrlThreatSubmissionContentType_File,
		"url":   SecurityEmailUrlThreatSubmissionContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionContentType(input)
	return &out, nil
}
