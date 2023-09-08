package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionWipeActionStatus string

const (
	WindowsInformationProtectionWipeActionStatusactive       WindowsInformationProtectionWipeActionStatus = "Active"
	WindowsInformationProtectionWipeActionStatuscanceled     WindowsInformationProtectionWipeActionStatus = "Canceled"
	WindowsInformationProtectionWipeActionStatusdone         WindowsInformationProtectionWipeActionStatus = "Done"
	WindowsInformationProtectionWipeActionStatusfailed       WindowsInformationProtectionWipeActionStatus = "Failed"
	WindowsInformationProtectionWipeActionStatusnone         WindowsInformationProtectionWipeActionStatus = "None"
	WindowsInformationProtectionWipeActionStatusnotSupported WindowsInformationProtectionWipeActionStatus = "NotSupported"
	WindowsInformationProtectionWipeActionStatuspending      WindowsInformationProtectionWipeActionStatus = "Pending"
)

func PossibleValuesForWindowsInformationProtectionWipeActionStatus() []string {
	return []string{
		string(WindowsInformationProtectionWipeActionStatusactive),
		string(WindowsInformationProtectionWipeActionStatuscanceled),
		string(WindowsInformationProtectionWipeActionStatusdone),
		string(WindowsInformationProtectionWipeActionStatusfailed),
		string(WindowsInformationProtectionWipeActionStatusnone),
		string(WindowsInformationProtectionWipeActionStatusnotSupported),
		string(WindowsInformationProtectionWipeActionStatuspending),
	}
}

func parseWindowsInformationProtectionWipeActionStatus(input string) (*WindowsInformationProtectionWipeActionStatus, error) {
	vals := map[string]WindowsInformationProtectionWipeActionStatus{
		"active":       WindowsInformationProtectionWipeActionStatusactive,
		"canceled":     WindowsInformationProtectionWipeActionStatuscanceled,
		"done":         WindowsInformationProtectionWipeActionStatusdone,
		"failed":       WindowsInformationProtectionWipeActionStatusfailed,
		"none":         WindowsInformationProtectionWipeActionStatusnone,
		"notsupported": WindowsInformationProtectionWipeActionStatusnotSupported,
		"pending":      WindowsInformationProtectionWipeActionStatuspending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionWipeActionStatus(input)
	return &out, nil
}
