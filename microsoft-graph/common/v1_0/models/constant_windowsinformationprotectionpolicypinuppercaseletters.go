package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionPolicyPinUppercaseLetters string

const (
	WindowsInformationProtectionPolicyPinUppercaseLettersallow             WindowsInformationProtectionPolicyPinUppercaseLetters = "Allow"
	WindowsInformationProtectionPolicyPinUppercaseLettersnotAllow          WindowsInformationProtectionPolicyPinUppercaseLetters = "NotAllow"
	WindowsInformationProtectionPolicyPinUppercaseLettersrequireAtLeastOne WindowsInformationProtectionPolicyPinUppercaseLetters = "RequireAtLeastOne"
)

func PossibleValuesForWindowsInformationProtectionPolicyPinUppercaseLetters() []string {
	return []string{
		string(WindowsInformationProtectionPolicyPinUppercaseLettersallow),
		string(WindowsInformationProtectionPolicyPinUppercaseLettersnotAllow),
		string(WindowsInformationProtectionPolicyPinUppercaseLettersrequireAtLeastOne),
	}
}

func parseWindowsInformationProtectionPolicyPinUppercaseLetters(input string) (*WindowsInformationProtectionPolicyPinUppercaseLetters, error) {
	vals := map[string]WindowsInformationProtectionPolicyPinUppercaseLetters{
		"allow":             WindowsInformationProtectionPolicyPinUppercaseLettersallow,
		"notallow":          WindowsInformationProtectionPolicyPinUppercaseLettersnotAllow,
		"requireatleastone": WindowsInformationProtectionPolicyPinUppercaseLettersrequireAtLeastOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionPolicyPinUppercaseLetters(input)
	return &out, nil
}
