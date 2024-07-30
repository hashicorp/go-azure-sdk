package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionPolicyPinLowercaseLetters string

const (
	WindowsInformationProtectionPolicyPinLowercaseLetters_Allow             WindowsInformationProtectionPolicyPinLowercaseLetters = "allow"
	WindowsInformationProtectionPolicyPinLowercaseLetters_NotAllow          WindowsInformationProtectionPolicyPinLowercaseLetters = "notAllow"
	WindowsInformationProtectionPolicyPinLowercaseLetters_RequireAtLeastOne WindowsInformationProtectionPolicyPinLowercaseLetters = "requireAtLeastOne"
)

func PossibleValuesForWindowsInformationProtectionPolicyPinLowercaseLetters() []string {
	return []string{
		string(WindowsInformationProtectionPolicyPinLowercaseLetters_Allow),
		string(WindowsInformationProtectionPolicyPinLowercaseLetters_NotAllow),
		string(WindowsInformationProtectionPolicyPinLowercaseLetters_RequireAtLeastOne),
	}
}

func (s *WindowsInformationProtectionPolicyPinLowercaseLetters) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsInformationProtectionPolicyPinLowercaseLetters(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsInformationProtectionPolicyPinLowercaseLetters(input string) (*WindowsInformationProtectionPolicyPinLowercaseLetters, error) {
	vals := map[string]WindowsInformationProtectionPolicyPinLowercaseLetters{
		"allow":             WindowsInformationProtectionPolicyPinLowercaseLetters_Allow,
		"notallow":          WindowsInformationProtectionPolicyPinLowercaseLetters_NotAllow,
		"requireatleastone": WindowsInformationProtectionPolicyPinLowercaseLetters_RequireAtLeastOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionPolicyPinLowercaseLetters(input)
	return &out, nil
}
