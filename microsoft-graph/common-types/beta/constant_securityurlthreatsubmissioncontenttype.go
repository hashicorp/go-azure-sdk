package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlThreatSubmissionContentType string

const (
	SecurityUrlThreatSubmissionContentType_App   SecurityUrlThreatSubmissionContentType = "app"
	SecurityUrlThreatSubmissionContentType_Email SecurityUrlThreatSubmissionContentType = "email"
	SecurityUrlThreatSubmissionContentType_File  SecurityUrlThreatSubmissionContentType = "file"
	SecurityUrlThreatSubmissionContentType_Url   SecurityUrlThreatSubmissionContentType = "url"
)

func PossibleValuesForSecurityUrlThreatSubmissionContentType() []string {
	return []string{
		string(SecurityUrlThreatSubmissionContentType_App),
		string(SecurityUrlThreatSubmissionContentType_Email),
		string(SecurityUrlThreatSubmissionContentType_File),
		string(SecurityUrlThreatSubmissionContentType_Url),
	}
}

func (s *SecurityUrlThreatSubmissionContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUrlThreatSubmissionContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUrlThreatSubmissionContentType(input string) (*SecurityUrlThreatSubmissionContentType, error) {
	vals := map[string]SecurityUrlThreatSubmissionContentType{
		"app":   SecurityUrlThreatSubmissionContentType_App,
		"email": SecurityUrlThreatSubmissionContentType_Email,
		"file":  SecurityUrlThreatSubmissionContentType_File,
		"url":   SecurityUrlThreatSubmissionContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlThreatSubmissionContentType(input)
	return &out, nil
}
