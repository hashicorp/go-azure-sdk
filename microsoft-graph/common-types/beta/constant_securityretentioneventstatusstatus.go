package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionEventStatusStatus string

const (
	SecurityRetentionEventStatusStatus_Error        SecurityRetentionEventStatusStatus = "error"
	SecurityRetentionEventStatusStatus_NotAvaliable SecurityRetentionEventStatusStatus = "notAvaliable"
	SecurityRetentionEventStatusStatus_Pending      SecurityRetentionEventStatusStatus = "pending"
	SecurityRetentionEventStatusStatus_Success      SecurityRetentionEventStatusStatus = "success"
)

func PossibleValuesForSecurityRetentionEventStatusStatus() []string {
	return []string{
		string(SecurityRetentionEventStatusStatus_Error),
		string(SecurityRetentionEventStatusStatus_NotAvaliable),
		string(SecurityRetentionEventStatusStatus_Pending),
		string(SecurityRetentionEventStatusStatus_Success),
	}
}

func (s *SecurityRetentionEventStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRetentionEventStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRetentionEventStatusStatus(input string) (*SecurityRetentionEventStatusStatus, error) {
	vals := map[string]SecurityRetentionEventStatusStatus{
		"error":        SecurityRetentionEventStatusStatus_Error,
		"notavaliable": SecurityRetentionEventStatusStatus_NotAvaliable,
		"pending":      SecurityRetentionEventStatusStatus_Pending,
		"success":      SecurityRetentionEventStatusStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRetentionEventStatusStatus(input)
	return &out, nil
}
