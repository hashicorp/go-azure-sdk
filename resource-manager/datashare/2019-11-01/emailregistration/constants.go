package emailregistration

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistrationStatus string

const (
	RegistrationStatusActivated                   RegistrationStatus = "Activated"
	RegistrationStatusActivationAttemptsExhausted RegistrationStatus = "ActivationAttemptsExhausted"
	RegistrationStatusActivationPending           RegistrationStatus = "ActivationPending"
)

func PossibleValuesForRegistrationStatus() []string {
	return []string{
		string(RegistrationStatusActivated),
		string(RegistrationStatusActivationAttemptsExhausted),
		string(RegistrationStatusActivationPending),
	}
}

func parseRegistrationStatus(input string) (*RegistrationStatus, error) {
	vals := map[string]RegistrationStatus{
		"activated":                   RegistrationStatusActivated,
		"activationattemptsexhausted": RegistrationStatusActivationAttemptsExhausted,
		"activationpending":           RegistrationStatusActivationPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistrationStatus(input)
	return &out, nil
}
