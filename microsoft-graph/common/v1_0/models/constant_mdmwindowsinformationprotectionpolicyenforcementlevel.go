package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MdmWindowsInformationProtectionPolicyEnforcementLevel string

const (
	MdmWindowsInformationProtectionPolicyEnforcementLevelencryptAndAuditOnly   MdmWindowsInformationProtectionPolicyEnforcementLevel = "EncryptAndAuditOnly"
	MdmWindowsInformationProtectionPolicyEnforcementLevelencryptAuditAndBlock  MdmWindowsInformationProtectionPolicyEnforcementLevel = "EncryptAuditAndBlock"
	MdmWindowsInformationProtectionPolicyEnforcementLevelencryptAuditAndPrompt MdmWindowsInformationProtectionPolicyEnforcementLevel = "EncryptAuditAndPrompt"
	MdmWindowsInformationProtectionPolicyEnforcementLevelnoProtection          MdmWindowsInformationProtectionPolicyEnforcementLevel = "NoProtection"
)

func PossibleValuesForMdmWindowsInformationProtectionPolicyEnforcementLevel() []string {
	return []string{
		string(MdmWindowsInformationProtectionPolicyEnforcementLevelencryptAndAuditOnly),
		string(MdmWindowsInformationProtectionPolicyEnforcementLevelencryptAuditAndBlock),
		string(MdmWindowsInformationProtectionPolicyEnforcementLevelencryptAuditAndPrompt),
		string(MdmWindowsInformationProtectionPolicyEnforcementLevelnoProtection),
	}
}

func parseMdmWindowsInformationProtectionPolicyEnforcementLevel(input string) (*MdmWindowsInformationProtectionPolicyEnforcementLevel, error) {
	vals := map[string]MdmWindowsInformationProtectionPolicyEnforcementLevel{
		"encryptandauditonly":   MdmWindowsInformationProtectionPolicyEnforcementLevelencryptAndAuditOnly,
		"encryptauditandblock":  MdmWindowsInformationProtectionPolicyEnforcementLevelencryptAuditAndBlock,
		"encryptauditandprompt": MdmWindowsInformationProtectionPolicyEnforcementLevelencryptAuditAndPrompt,
		"noprotection":          MdmWindowsInformationProtectionPolicyEnforcementLevelnoProtection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MdmWindowsInformationProtectionPolicyEnforcementLevel(input)
	return &out, nil
}
