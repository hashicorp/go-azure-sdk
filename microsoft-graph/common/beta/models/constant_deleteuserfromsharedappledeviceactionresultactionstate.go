package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteUserFromSharedAppleDeviceActionResultActionState string

const (
	DeleteUserFromSharedAppleDeviceActionResultActionStateactive       DeleteUserFromSharedAppleDeviceActionResultActionState = "Active"
	DeleteUserFromSharedAppleDeviceActionResultActionStatecanceled     DeleteUserFromSharedAppleDeviceActionResultActionState = "Canceled"
	DeleteUserFromSharedAppleDeviceActionResultActionStatedone         DeleteUserFromSharedAppleDeviceActionResultActionState = "Done"
	DeleteUserFromSharedAppleDeviceActionResultActionStatefailed       DeleteUserFromSharedAppleDeviceActionResultActionState = "Failed"
	DeleteUserFromSharedAppleDeviceActionResultActionStatenone         DeleteUserFromSharedAppleDeviceActionResultActionState = "None"
	DeleteUserFromSharedAppleDeviceActionResultActionStatenotSupported DeleteUserFromSharedAppleDeviceActionResultActionState = "NotSupported"
	DeleteUserFromSharedAppleDeviceActionResultActionStatepending      DeleteUserFromSharedAppleDeviceActionResultActionState = "Pending"
)

func PossibleValuesForDeleteUserFromSharedAppleDeviceActionResultActionState() []string {
	return []string{
		string(DeleteUserFromSharedAppleDeviceActionResultActionStateactive),
		string(DeleteUserFromSharedAppleDeviceActionResultActionStatecanceled),
		string(DeleteUserFromSharedAppleDeviceActionResultActionStatedone),
		string(DeleteUserFromSharedAppleDeviceActionResultActionStatefailed),
		string(DeleteUserFromSharedAppleDeviceActionResultActionStatenone),
		string(DeleteUserFromSharedAppleDeviceActionResultActionStatenotSupported),
		string(DeleteUserFromSharedAppleDeviceActionResultActionStatepending),
	}
}

func parseDeleteUserFromSharedAppleDeviceActionResultActionState(input string) (*DeleteUserFromSharedAppleDeviceActionResultActionState, error) {
	vals := map[string]DeleteUserFromSharedAppleDeviceActionResultActionState{
		"active":       DeleteUserFromSharedAppleDeviceActionResultActionStateactive,
		"canceled":     DeleteUserFromSharedAppleDeviceActionResultActionStatecanceled,
		"done":         DeleteUserFromSharedAppleDeviceActionResultActionStatedone,
		"failed":       DeleteUserFromSharedAppleDeviceActionResultActionStatefailed,
		"none":         DeleteUserFromSharedAppleDeviceActionResultActionStatenone,
		"notsupported": DeleteUserFromSharedAppleDeviceActionResultActionStatenotSupported,
		"pending":      DeleteUserFromSharedAppleDeviceActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeleteUserFromSharedAppleDeviceActionResultActionState(input)
	return &out, nil
}
