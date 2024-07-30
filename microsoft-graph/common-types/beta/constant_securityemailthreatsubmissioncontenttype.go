package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionContentType string

const (
	SecurityEmailThreatSubmissionContentType_App   SecurityEmailThreatSubmissionContentType = "app"
	SecurityEmailThreatSubmissionContentType_Email SecurityEmailThreatSubmissionContentType = "email"
	SecurityEmailThreatSubmissionContentType_File  SecurityEmailThreatSubmissionContentType = "file"
	SecurityEmailThreatSubmissionContentType_Url   SecurityEmailThreatSubmissionContentType = "url"
)

func PossibleValuesForSecurityEmailThreatSubmissionContentType() []string {
	return []string{
		string(SecurityEmailThreatSubmissionContentType_App),
		string(SecurityEmailThreatSubmissionContentType_Email),
		string(SecurityEmailThreatSubmissionContentType_File),
		string(SecurityEmailThreatSubmissionContentType_Url),
	}
}

func (s *SecurityEmailThreatSubmissionContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailThreatSubmissionContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailThreatSubmissionContentType(input string) (*SecurityEmailThreatSubmissionContentType, error) {
	vals := map[string]SecurityEmailThreatSubmissionContentType{
		"app":   SecurityEmailThreatSubmissionContentType_App,
		"email": SecurityEmailThreatSubmissionContentType_Email,
		"file":  SecurityEmailThreatSubmissionContentType_File,
		"url":   SecurityEmailThreatSubmissionContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailThreatSubmissionContentType(input)
	return &out, nil
}
