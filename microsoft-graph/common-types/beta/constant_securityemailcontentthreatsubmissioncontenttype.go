package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionContentType string

const (
	SecurityEmailContentThreatSubmissionContentType_App   SecurityEmailContentThreatSubmissionContentType = "app"
	SecurityEmailContentThreatSubmissionContentType_Email SecurityEmailContentThreatSubmissionContentType = "email"
	SecurityEmailContentThreatSubmissionContentType_File  SecurityEmailContentThreatSubmissionContentType = "file"
	SecurityEmailContentThreatSubmissionContentType_Url   SecurityEmailContentThreatSubmissionContentType = "url"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionContentType() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionContentType_App),
		string(SecurityEmailContentThreatSubmissionContentType_Email),
		string(SecurityEmailContentThreatSubmissionContentType_File),
		string(SecurityEmailContentThreatSubmissionContentType_Url),
	}
}

func (s *SecurityEmailContentThreatSubmissionContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailContentThreatSubmissionContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailContentThreatSubmissionContentType(input string) (*SecurityEmailContentThreatSubmissionContentType, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionContentType{
		"app":   SecurityEmailContentThreatSubmissionContentType_App,
		"email": SecurityEmailContentThreatSubmissionContentType_Email,
		"file":  SecurityEmailContentThreatSubmissionContentType_File,
		"url":   SecurityEmailContentThreatSubmissionContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionContentType(input)
	return &out, nil
}
