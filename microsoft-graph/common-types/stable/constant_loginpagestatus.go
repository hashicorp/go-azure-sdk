package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoginPageStatus string

const (
	LoginPageStatus_Archive LoginPageStatus = "archive"
	LoginPageStatus_Delete  LoginPageStatus = "delete"
	LoginPageStatus_Draft   LoginPageStatus = "draft"
	LoginPageStatus_Ready   LoginPageStatus = "ready"
	LoginPageStatus_Unknown LoginPageStatus = "unknown"
)

func PossibleValuesForLoginPageStatus() []string {
	return []string{
		string(LoginPageStatus_Archive),
		string(LoginPageStatus_Delete),
		string(LoginPageStatus_Draft),
		string(LoginPageStatus_Ready),
		string(LoginPageStatus_Unknown),
	}
}

func (s *LoginPageStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLoginPageStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLoginPageStatus(input string) (*LoginPageStatus, error) {
	vals := map[string]LoginPageStatus{
		"archive": LoginPageStatus_Archive,
		"delete":  LoginPageStatus_Delete,
		"draft":   LoginPageStatus_Draft,
		"ready":   LoginPageStatus_Ready,
		"unknown": LoginPageStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoginPageStatus(input)
	return &out, nil
}
