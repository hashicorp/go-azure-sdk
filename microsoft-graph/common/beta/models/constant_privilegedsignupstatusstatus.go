package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedSignupStatusStatus string

const (
	PrivilegedSignupStatusStatusdisabled                      PrivilegedSignupStatusStatus = "Disabled"
	PrivilegedSignupStatusStatusnotRegisteredYet              PrivilegedSignupStatusStatus = "NotRegisteredYet"
	PrivilegedSignupStatusStatusregisteredSetupInProgress     PrivilegedSignupStatusStatus = "RegisteredSetupInProgress"
	PrivilegedSignupStatusStatusregisteredSetupNotStarted     PrivilegedSignupStatusStatus = "RegisteredSetupNotStarted"
	PrivilegedSignupStatusStatusregistrationAndSetupCompleted PrivilegedSignupStatusStatus = "RegistrationAndSetupCompleted"
	PrivilegedSignupStatusStatusregistrationFailed            PrivilegedSignupStatusStatus = "RegistrationFailed"
	PrivilegedSignupStatusStatusregistrationTimedOut          PrivilegedSignupStatusStatus = "RegistrationTimedOut"
	PrivilegedSignupStatusStatusunknown                       PrivilegedSignupStatusStatus = "Unknown"
)

func PossibleValuesForPrivilegedSignupStatusStatus() []string {
	return []string{
		string(PrivilegedSignupStatusStatusdisabled),
		string(PrivilegedSignupStatusStatusnotRegisteredYet),
		string(PrivilegedSignupStatusStatusregisteredSetupInProgress),
		string(PrivilegedSignupStatusStatusregisteredSetupNotStarted),
		string(PrivilegedSignupStatusStatusregistrationAndSetupCompleted),
		string(PrivilegedSignupStatusStatusregistrationFailed),
		string(PrivilegedSignupStatusStatusregistrationTimedOut),
		string(PrivilegedSignupStatusStatusunknown),
	}
}

func parsePrivilegedSignupStatusStatus(input string) (*PrivilegedSignupStatusStatus, error) {
	vals := map[string]PrivilegedSignupStatusStatus{
		"disabled":                      PrivilegedSignupStatusStatusdisabled,
		"notregisteredyet":              PrivilegedSignupStatusStatusnotRegisteredYet,
		"registeredsetupinprogress":     PrivilegedSignupStatusStatusregisteredSetupInProgress,
		"registeredsetupnotstarted":     PrivilegedSignupStatusStatusregisteredSetupNotStarted,
		"registrationandsetupcompleted": PrivilegedSignupStatusStatusregistrationAndSetupCompleted,
		"registrationfailed":            PrivilegedSignupStatusStatusregistrationFailed,
		"registrationtimedout":          PrivilegedSignupStatusStatusregistrationTimedOut,
		"unknown":                       PrivilegedSignupStatusStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedSignupStatusStatus(input)
	return &out, nil
}
