package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityStateLogonType string

const (
	UserSecurityStateLogonType_Batch             UserSecurityStateLogonType = "batch"
	UserSecurityStateLogonType_Interactive       UserSecurityStateLogonType = "interactive"
	UserSecurityStateLogonType_Network           UserSecurityStateLogonType = "network"
	UserSecurityStateLogonType_RemoteInteractive UserSecurityStateLogonType = "remoteInteractive"
	UserSecurityStateLogonType_Service           UserSecurityStateLogonType = "service"
	UserSecurityStateLogonType_Unknown           UserSecurityStateLogonType = "unknown"
)

func PossibleValuesForUserSecurityStateLogonType() []string {
	return []string{
		string(UserSecurityStateLogonType_Batch),
		string(UserSecurityStateLogonType_Interactive),
		string(UserSecurityStateLogonType_Network),
		string(UserSecurityStateLogonType_RemoteInteractive),
		string(UserSecurityStateLogonType_Service),
		string(UserSecurityStateLogonType_Unknown),
	}
}

func (s *UserSecurityStateLogonType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserSecurityStateLogonType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserSecurityStateLogonType(input string) (*UserSecurityStateLogonType, error) {
	vals := map[string]UserSecurityStateLogonType{
		"batch":             UserSecurityStateLogonType_Batch,
		"interactive":       UserSecurityStateLogonType_Interactive,
		"network":           UserSecurityStateLogonType_Network,
		"remoteinteractive": UserSecurityStateLogonType_RemoteInteractive,
		"service":           UserSecurityStateLogonType_Service,
		"unknown":           UserSecurityStateLogonType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserSecurityStateLogonType(input)
	return &out, nil
}
