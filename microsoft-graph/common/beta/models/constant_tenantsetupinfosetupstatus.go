package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantSetupInfoSetupStatus string

const (
	TenantSetupInfoSetupStatusdisabled                      TenantSetupInfoSetupStatus = "Disabled"
	TenantSetupInfoSetupStatusnotRegisteredYet              TenantSetupInfoSetupStatus = "NotRegisteredYet"
	TenantSetupInfoSetupStatusregisteredSetupInProgress     TenantSetupInfoSetupStatus = "RegisteredSetupInProgress"
	TenantSetupInfoSetupStatusregisteredSetupNotStarted     TenantSetupInfoSetupStatus = "RegisteredSetupNotStarted"
	TenantSetupInfoSetupStatusregistrationAndSetupCompleted TenantSetupInfoSetupStatus = "RegistrationAndSetupCompleted"
	TenantSetupInfoSetupStatusregistrationFailed            TenantSetupInfoSetupStatus = "RegistrationFailed"
	TenantSetupInfoSetupStatusregistrationTimedOut          TenantSetupInfoSetupStatus = "RegistrationTimedOut"
	TenantSetupInfoSetupStatusunknown                       TenantSetupInfoSetupStatus = "Unknown"
)

func PossibleValuesForTenantSetupInfoSetupStatus() []string {
	return []string{
		string(TenantSetupInfoSetupStatusdisabled),
		string(TenantSetupInfoSetupStatusnotRegisteredYet),
		string(TenantSetupInfoSetupStatusregisteredSetupInProgress),
		string(TenantSetupInfoSetupStatusregisteredSetupNotStarted),
		string(TenantSetupInfoSetupStatusregistrationAndSetupCompleted),
		string(TenantSetupInfoSetupStatusregistrationFailed),
		string(TenantSetupInfoSetupStatusregistrationTimedOut),
		string(TenantSetupInfoSetupStatusunknown),
	}
}

func parseTenantSetupInfoSetupStatus(input string) (*TenantSetupInfoSetupStatus, error) {
	vals := map[string]TenantSetupInfoSetupStatus{
		"disabled":                      TenantSetupInfoSetupStatusdisabled,
		"notregisteredyet":              TenantSetupInfoSetupStatusnotRegisteredYet,
		"registeredsetupinprogress":     TenantSetupInfoSetupStatusregisteredSetupInProgress,
		"registeredsetupnotstarted":     TenantSetupInfoSetupStatusregisteredSetupNotStarted,
		"registrationandsetupcompleted": TenantSetupInfoSetupStatusregistrationAndSetupCompleted,
		"registrationfailed":            TenantSetupInfoSetupStatusregistrationFailed,
		"registrationtimedout":          TenantSetupInfoSetupStatusregistrationTimedOut,
		"unknown":                       TenantSetupInfoSetupStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TenantSetupInfoSetupStatus(input)
	return &out, nil
}
