package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionPolicyPinUppercaseLetters string

const (
	WindowsInformationProtectionPolicyPinUppercaseLetters_Allow             WindowsInformationProtectionPolicyPinUppercaseLetters = "allow"
	WindowsInformationProtectionPolicyPinUppercaseLetters_NotAllow          WindowsInformationProtectionPolicyPinUppercaseLetters = "notAllow"
	WindowsInformationProtectionPolicyPinUppercaseLetters_RequireAtLeastOne WindowsInformationProtectionPolicyPinUppercaseLetters = "requireAtLeastOne"
)

func PossibleValuesForWindowsInformationProtectionPolicyPinUppercaseLetters() []string {
	return []string{
		string(WindowsInformationProtectionPolicyPinUppercaseLetters_Allow),
		string(WindowsInformationProtectionPolicyPinUppercaseLetters_NotAllow),
		string(WindowsInformationProtectionPolicyPinUppercaseLetters_RequireAtLeastOne),
	}
}

func (s *WindowsInformationProtectionPolicyPinUppercaseLetters) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsInformationProtectionPolicyPinUppercaseLetters(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsInformationProtectionPolicyPinUppercaseLetters(input string) (*WindowsInformationProtectionPolicyPinUppercaseLetters, error) {
	vals := map[string]WindowsInformationProtectionPolicyPinUppercaseLetters{
		"allow":             WindowsInformationProtectionPolicyPinUppercaseLetters_Allow,
		"notallow":          WindowsInformationProtectionPolicyPinUppercaseLetters_NotAllow,
		"requireatleastone": WindowsInformationProtectionPolicyPinUppercaseLetters_RequireAtLeastOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionPolicyPinUppercaseLetters(input)
	return &out, nil
}
