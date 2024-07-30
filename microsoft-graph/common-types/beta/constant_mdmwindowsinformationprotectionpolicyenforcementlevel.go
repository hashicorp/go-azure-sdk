package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MdmWindowsInformationProtectionPolicyEnforcementLevel string

const (
	MdmWindowsInformationProtectionPolicyEnforcementLevel_EncryptAndAuditOnly   MdmWindowsInformationProtectionPolicyEnforcementLevel = "encryptAndAuditOnly"
	MdmWindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndBlock  MdmWindowsInformationProtectionPolicyEnforcementLevel = "encryptAuditAndBlock"
	MdmWindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndPrompt MdmWindowsInformationProtectionPolicyEnforcementLevel = "encryptAuditAndPrompt"
	MdmWindowsInformationProtectionPolicyEnforcementLevel_NoProtection          MdmWindowsInformationProtectionPolicyEnforcementLevel = "noProtection"
)

func PossibleValuesForMdmWindowsInformationProtectionPolicyEnforcementLevel() []string {
	return []string{
		string(MdmWindowsInformationProtectionPolicyEnforcementLevel_EncryptAndAuditOnly),
		string(MdmWindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndBlock),
		string(MdmWindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndPrompt),
		string(MdmWindowsInformationProtectionPolicyEnforcementLevel_NoProtection),
	}
}

func (s *MdmWindowsInformationProtectionPolicyEnforcementLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMdmWindowsInformationProtectionPolicyEnforcementLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMdmWindowsInformationProtectionPolicyEnforcementLevel(input string) (*MdmWindowsInformationProtectionPolicyEnforcementLevel, error) {
	vals := map[string]MdmWindowsInformationProtectionPolicyEnforcementLevel{
		"encryptandauditonly":   MdmWindowsInformationProtectionPolicyEnforcementLevel_EncryptAndAuditOnly,
		"encryptauditandblock":  MdmWindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndBlock,
		"encryptauditandprompt": MdmWindowsInformationProtectionPolicyEnforcementLevel_EncryptAuditAndPrompt,
		"noprotection":          MdmWindowsInformationProtectionPolicyEnforcementLevel_NoProtection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MdmWindowsInformationProtectionPolicyEnforcementLevel(input)
	return &out, nil
}
