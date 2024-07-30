package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessGrantControlsBuiltInControls string

const (
	ConditionalAccessGrantControlsBuiltInControls_ApprovedApplication  ConditionalAccessGrantControlsBuiltInControls = "approvedApplication"
	ConditionalAccessGrantControlsBuiltInControls_Block                ConditionalAccessGrantControlsBuiltInControls = "block"
	ConditionalAccessGrantControlsBuiltInControls_CompliantApplication ConditionalAccessGrantControlsBuiltInControls = "compliantApplication"
	ConditionalAccessGrantControlsBuiltInControls_CompliantDevice      ConditionalAccessGrantControlsBuiltInControls = "compliantDevice"
	ConditionalAccessGrantControlsBuiltInControls_DomainJoinedDevice   ConditionalAccessGrantControlsBuiltInControls = "domainJoinedDevice"
	ConditionalAccessGrantControlsBuiltInControls_Mfa                  ConditionalAccessGrantControlsBuiltInControls = "mfa"
	ConditionalAccessGrantControlsBuiltInControls_PasswordChange       ConditionalAccessGrantControlsBuiltInControls = "passwordChange"
)

func PossibleValuesForConditionalAccessGrantControlsBuiltInControls() []string {
	return []string{
		string(ConditionalAccessGrantControlsBuiltInControls_ApprovedApplication),
		string(ConditionalAccessGrantControlsBuiltInControls_Block),
		string(ConditionalAccessGrantControlsBuiltInControls_CompliantApplication),
		string(ConditionalAccessGrantControlsBuiltInControls_CompliantDevice),
		string(ConditionalAccessGrantControlsBuiltInControls_DomainJoinedDevice),
		string(ConditionalAccessGrantControlsBuiltInControls_Mfa),
		string(ConditionalAccessGrantControlsBuiltInControls_PasswordChange),
	}
}

func (s *ConditionalAccessGrantControlsBuiltInControls) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessGrantControlsBuiltInControls(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessGrantControlsBuiltInControls(input string) (*ConditionalAccessGrantControlsBuiltInControls, error) {
	vals := map[string]ConditionalAccessGrantControlsBuiltInControls{
		"approvedapplication":  ConditionalAccessGrantControlsBuiltInControls_ApprovedApplication,
		"block":                ConditionalAccessGrantControlsBuiltInControls_Block,
		"compliantapplication": ConditionalAccessGrantControlsBuiltInControls_CompliantApplication,
		"compliantdevice":      ConditionalAccessGrantControlsBuiltInControls_CompliantDevice,
		"domainjoineddevice":   ConditionalAccessGrantControlsBuiltInControls_DomainJoinedDevice,
		"mfa":                  ConditionalAccessGrantControlsBuiltInControls_Mfa,
		"passwordchange":       ConditionalAccessGrantControlsBuiltInControls_PasswordChange,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessGrantControlsBuiltInControls(input)
	return &out, nil
}
