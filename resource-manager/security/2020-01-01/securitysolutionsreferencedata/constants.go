package securitysolutionsreferencedata

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFamily string

const (
	SecurityFamilyNgfw    SecurityFamily = "Ngfw"
	SecurityFamilySaasWaf SecurityFamily = "SaasWaf"
	SecurityFamilyVa      SecurityFamily = "Va"
	SecurityFamilyWaf     SecurityFamily = "Waf"
)

func PossibleValuesForSecurityFamily() []string {
	return []string{
		string(SecurityFamilyNgfw),
		string(SecurityFamilySaasWaf),
		string(SecurityFamilyVa),
		string(SecurityFamilyWaf),
	}
}

func parseSecurityFamily(input string) (*SecurityFamily, error) {
	vals := map[string]SecurityFamily{
		"ngfw":    SecurityFamilyNgfw,
		"saaswaf": SecurityFamilySaasWaf,
		"va":      SecurityFamilyVa,
		"waf":     SecurityFamilyWaf,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFamily(input)
	return &out, nil
}
