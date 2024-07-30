package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmissionContentType string

const (
	SecurityThreatSubmissionContentType_App   SecurityThreatSubmissionContentType = "app"
	SecurityThreatSubmissionContentType_Email SecurityThreatSubmissionContentType = "email"
	SecurityThreatSubmissionContentType_File  SecurityThreatSubmissionContentType = "file"
	SecurityThreatSubmissionContentType_Url   SecurityThreatSubmissionContentType = "url"
)

func PossibleValuesForSecurityThreatSubmissionContentType() []string {
	return []string{
		string(SecurityThreatSubmissionContentType_App),
		string(SecurityThreatSubmissionContentType_Email),
		string(SecurityThreatSubmissionContentType_File),
		string(SecurityThreatSubmissionContentType_Url),
	}
}

func (s *SecurityThreatSubmissionContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityThreatSubmissionContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityThreatSubmissionContentType(input string) (*SecurityThreatSubmissionContentType, error) {
	vals := map[string]SecurityThreatSubmissionContentType{
		"app":   SecurityThreatSubmissionContentType_App,
		"email": SecurityThreatSubmissionContentType_Email,
		"file":  SecurityThreatSubmissionContentType_File,
		"url":   SecurityThreatSubmissionContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityThreatSubmissionContentType(input)
	return &out, nil
}
