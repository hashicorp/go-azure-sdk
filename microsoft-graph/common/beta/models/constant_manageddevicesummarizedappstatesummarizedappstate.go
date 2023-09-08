package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceSummarizedAppStateSummarizedAppState string

const (
	ManagedDeviceSummarizedAppStateSummarizedAppStatefail          ManagedDeviceSummarizedAppStateSummarizedAppState = "Fail"
	ManagedDeviceSummarizedAppStateSummarizedAppStatenotApplicable ManagedDeviceSummarizedAppStateSummarizedAppState = "NotApplicable"
	ManagedDeviceSummarizedAppStateSummarizedAppStatepending       ManagedDeviceSummarizedAppStateSummarizedAppState = "Pending"
	ManagedDeviceSummarizedAppStateSummarizedAppStatescriptError   ManagedDeviceSummarizedAppStateSummarizedAppState = "ScriptError"
	ManagedDeviceSummarizedAppStateSummarizedAppStatesuccess       ManagedDeviceSummarizedAppStateSummarizedAppState = "Success"
	ManagedDeviceSummarizedAppStateSummarizedAppStateunknown       ManagedDeviceSummarizedAppStateSummarizedAppState = "Unknown"
)

func PossibleValuesForManagedDeviceSummarizedAppStateSummarizedAppState() []string {
	return []string{
		string(ManagedDeviceSummarizedAppStateSummarizedAppStatefail),
		string(ManagedDeviceSummarizedAppStateSummarizedAppStatenotApplicable),
		string(ManagedDeviceSummarizedAppStateSummarizedAppStatepending),
		string(ManagedDeviceSummarizedAppStateSummarizedAppStatescriptError),
		string(ManagedDeviceSummarizedAppStateSummarizedAppStatesuccess),
		string(ManagedDeviceSummarizedAppStateSummarizedAppStateunknown),
	}
}

func parseManagedDeviceSummarizedAppStateSummarizedAppState(input string) (*ManagedDeviceSummarizedAppStateSummarizedAppState, error) {
	vals := map[string]ManagedDeviceSummarizedAppStateSummarizedAppState{
		"fail":          ManagedDeviceSummarizedAppStateSummarizedAppStatefail,
		"notapplicable": ManagedDeviceSummarizedAppStateSummarizedAppStatenotApplicable,
		"pending":       ManagedDeviceSummarizedAppStateSummarizedAppStatepending,
		"scripterror":   ManagedDeviceSummarizedAppStateSummarizedAppStatescriptError,
		"success":       ManagedDeviceSummarizedAppStateSummarizedAppStatesuccess,
		"unknown":       ManagedDeviceSummarizedAppStateSummarizedAppStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceSummarizedAppStateSummarizedAppState(input)
	return &out, nil
}
