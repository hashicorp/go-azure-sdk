package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogonUserLogonTypes string

const (
	LogonUserLogonTypesbatch             LogonUserLogonTypes = "Batch"
	LogonUserLogonTypesinteractive       LogonUserLogonTypes = "Interactive"
	LogonUserLogonTypesnetwork           LogonUserLogonTypes = "Network"
	LogonUserLogonTypesremoteInteractive LogonUserLogonTypes = "RemoteInteractive"
	LogonUserLogonTypesservice           LogonUserLogonTypes = "Service"
	LogonUserLogonTypesunknown           LogonUserLogonTypes = "Unknown"
)

func PossibleValuesForLogonUserLogonTypes() []string {
	return []string{
		string(LogonUserLogonTypesbatch),
		string(LogonUserLogonTypesinteractive),
		string(LogonUserLogonTypesnetwork),
		string(LogonUserLogonTypesremoteInteractive),
		string(LogonUserLogonTypesservice),
		string(LogonUserLogonTypesunknown),
	}
}

func parseLogonUserLogonTypes(input string) (*LogonUserLogonTypes, error) {
	vals := map[string]LogonUserLogonTypes{
		"batch":             LogonUserLogonTypesbatch,
		"interactive":       LogonUserLogonTypesinteractive,
		"network":           LogonUserLogonTypesnetwork,
		"remoteinteractive": LogonUserLogonTypesremoteInteractive,
		"service":           LogonUserLogonTypesservice,
		"unknown":           LogonUserLogonTypesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogonUserLogonTypes(input)
	return &out, nil
}
