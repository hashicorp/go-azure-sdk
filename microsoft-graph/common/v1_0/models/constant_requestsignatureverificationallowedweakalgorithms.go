package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestSignatureVerificationAllowedWeakAlgorithms string

const (
	RequestSignatureVerificationAllowedWeakAlgorithmsrsaSha1 RequestSignatureVerificationAllowedWeakAlgorithms = "RsaSha1"
)

func PossibleValuesForRequestSignatureVerificationAllowedWeakAlgorithms() []string {
	return []string{
		string(RequestSignatureVerificationAllowedWeakAlgorithmsrsaSha1),
	}
}

func parseRequestSignatureVerificationAllowedWeakAlgorithms(input string) (*RequestSignatureVerificationAllowedWeakAlgorithms, error) {
	vals := map[string]RequestSignatureVerificationAllowedWeakAlgorithms{
		"rsasha1": RequestSignatureVerificationAllowedWeakAlgorithmsrsaSha1,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RequestSignatureVerificationAllowedWeakAlgorithms(input)
	return &out, nil
}
