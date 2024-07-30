package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoginPageSource string

const (
	LoginPageSource_Global  LoginPageSource = "global"
	LoginPageSource_Tenant  LoginPageSource = "tenant"
	LoginPageSource_Unknown LoginPageSource = "unknown"
)

func PossibleValuesForLoginPageSource() []string {
	return []string{
		string(LoginPageSource_Global),
		string(LoginPageSource_Tenant),
		string(LoginPageSource_Unknown),
	}
}

func (s *LoginPageSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLoginPageSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLoginPageSource(input string) (*LoginPageSource, error) {
	vals := map[string]LoginPageSource{
		"global":  LoginPageSource_Global,
		"tenant":  LoginPageSource_Tenant,
		"unknown": LoginPageSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoginPageSource(input)
	return &out, nil
}
