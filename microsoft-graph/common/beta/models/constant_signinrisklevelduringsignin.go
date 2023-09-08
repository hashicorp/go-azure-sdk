package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInRiskLevelDuringSignIn string

const (
	SignInRiskLevelDuringSignInhidden SignInRiskLevelDuringSignIn = "Hidden"
	SignInRiskLevelDuringSignInhigh   SignInRiskLevelDuringSignIn = "High"
	SignInRiskLevelDuringSignInlow    SignInRiskLevelDuringSignIn = "Low"
	SignInRiskLevelDuringSignInmedium SignInRiskLevelDuringSignIn = "Medium"
	SignInRiskLevelDuringSignInnone   SignInRiskLevelDuringSignIn = "None"
)

func PossibleValuesForSignInRiskLevelDuringSignIn() []string {
	return []string{
		string(SignInRiskLevelDuringSignInhidden),
		string(SignInRiskLevelDuringSignInhigh),
		string(SignInRiskLevelDuringSignInlow),
		string(SignInRiskLevelDuringSignInmedium),
		string(SignInRiskLevelDuringSignInnone),
	}
}

func parseSignInRiskLevelDuringSignIn(input string) (*SignInRiskLevelDuringSignIn, error) {
	vals := map[string]SignInRiskLevelDuringSignIn{
		"hidden": SignInRiskLevelDuringSignInhidden,
		"high":   SignInRiskLevelDuringSignInhigh,
		"low":    SignInRiskLevelDuringSignInlow,
		"medium": SignInRiskLevelDuringSignInmedium,
		"none":   SignInRiskLevelDuringSignInnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInRiskLevelDuringSignIn(input)
	return &out, nil
}
