package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionPolicyPinLowercaseLetters string

const (
	WindowsInformationProtectionPolicyPinLowercaseLettersallow             WindowsInformationProtectionPolicyPinLowercaseLetters = "Allow"
	WindowsInformationProtectionPolicyPinLowercaseLettersnotAllow          WindowsInformationProtectionPolicyPinLowercaseLetters = "NotAllow"
	WindowsInformationProtectionPolicyPinLowercaseLettersrequireAtLeastOne WindowsInformationProtectionPolicyPinLowercaseLetters = "RequireAtLeastOne"
)

func PossibleValuesForWindowsInformationProtectionPolicyPinLowercaseLetters() []string {
	return []string{
		string(WindowsInformationProtectionPolicyPinLowercaseLettersallow),
		string(WindowsInformationProtectionPolicyPinLowercaseLettersnotAllow),
		string(WindowsInformationProtectionPolicyPinLowercaseLettersrequireAtLeastOne),
	}
}

func parseWindowsInformationProtectionPolicyPinLowercaseLetters(input string) (*WindowsInformationProtectionPolicyPinLowercaseLetters, error) {
	vals := map[string]WindowsInformationProtectionPolicyPinLowercaseLetters{
		"allow":             WindowsInformationProtectionPolicyPinLowercaseLettersallow,
		"notallow":          WindowsInformationProtectionPolicyPinLowercaseLettersnotAllow,
		"requireatleastone": WindowsInformationProtectionPolicyPinLowercaseLettersrequireAtLeastOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionPolicyPinLowercaseLetters(input)
	return &out, nil
}
