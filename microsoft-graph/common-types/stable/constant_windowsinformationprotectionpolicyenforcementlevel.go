package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionPolicyEnforcementLevel string

const (
	WindowsInformationProtectionPolicyEnforcementLevel_EncryptAndAuditOnly   WindowsInformationProtectionPolicyEnforcementLevel = "encryptAndAuditOnly"
	WindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndBlock  WindowsInformationProtectionPolicyEnforcementLevel = "encryptAuditAndBlock"
	WindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndPrompt WindowsInformationProtectionPolicyEnforcementLevel = "encryptAuditAndPrompt"
	WindowsInformationProtectionPolicyEnforcementLevel_NoProtection          WindowsInformationProtectionPolicyEnforcementLevel = "noProtection"
)

func PossibleValuesForWindowsInformationProtectionPolicyEnforcementLevel() []string {
	return []string{
		string(WindowsInformationProtectionPolicyEnforcementLevel_EncryptAndAuditOnly),
		string(WindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndBlock),
		string(WindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndPrompt),
		string(WindowsInformationProtectionPolicyEnforcementLevel_NoProtection),
	}
}

func (s *WindowsInformationProtectionPolicyEnforcementLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsInformationProtectionPolicyEnforcementLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsInformationProtectionPolicyEnforcementLevel(input string) (*WindowsInformationProtectionPolicyEnforcementLevel, error) {
	vals := map[string]WindowsInformationProtectionPolicyEnforcementLevel{
		"encryptandauditonly":   WindowsInformationProtectionPolicyEnforcementLevel_EncryptAndAuditOnly,
		"encryptauditandblock":  WindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndBlock,
		"encryptauditandprompt": WindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndPrompt,
		"noprotection":          WindowsInformationProtectionPolicyEnforcementLevel_NoProtection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionPolicyEnforcementLevel(input)
	return &out, nil
}
