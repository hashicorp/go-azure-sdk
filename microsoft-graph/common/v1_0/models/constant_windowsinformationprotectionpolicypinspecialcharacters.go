package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionPolicyPinSpecialCharacters string

const (
	WindowsInformationProtectionPolicyPinSpecialCharactersallow             WindowsInformationProtectionPolicyPinSpecialCharacters = "Allow"
	WindowsInformationProtectionPolicyPinSpecialCharactersnotAllow          WindowsInformationProtectionPolicyPinSpecialCharacters = "NotAllow"
	WindowsInformationProtectionPolicyPinSpecialCharactersrequireAtLeastOne WindowsInformationProtectionPolicyPinSpecialCharacters = "RequireAtLeastOne"
)

func PossibleValuesForWindowsInformationProtectionPolicyPinSpecialCharacters() []string {
	return []string{
		string(WindowsInformationProtectionPolicyPinSpecialCharactersallow),
		string(WindowsInformationProtectionPolicyPinSpecialCharactersnotAllow),
		string(WindowsInformationProtectionPolicyPinSpecialCharactersrequireAtLeastOne),
	}
}

func parseWindowsInformationProtectionPolicyPinSpecialCharacters(input string) (*WindowsInformationProtectionPolicyPinSpecialCharacters, error) {
	vals := map[string]WindowsInformationProtectionPolicyPinSpecialCharacters{
		"allow":             WindowsInformationProtectionPolicyPinSpecialCharactersallow,
		"notallow":          WindowsInformationProtectionPolicyPinSpecialCharactersnotAllow,
		"requireatleastone": WindowsInformationProtectionPolicyPinSpecialCharactersrequireAtLeastOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionPolicyPinSpecialCharacters(input)
	return &out, nil
}
