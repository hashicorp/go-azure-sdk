package emailregistration

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
