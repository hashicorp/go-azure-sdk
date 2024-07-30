package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteUserFromSharedAppleDeviceActionResultActionState string

const (
	DeleteUserFromSharedAppleDeviceActionResultActionState_Active       DeleteUserFromSharedAppleDeviceActionResultActionState = "active"
	DeleteUserFromSharedAppleDeviceActionResultActionState_Canceled     DeleteUserFromSharedAppleDeviceActionResultActionState = "canceled"
	DeleteUserFromSharedAppleDeviceActionResultActionState_Done         DeleteUserFromSharedAppleDeviceActionResultActionState = "done"
	DeleteUserFromSharedAppleDeviceActionResultActionState_Failed       DeleteUserFromSharedAppleDeviceActionResultActionState = "failed"
	DeleteUserFromSharedAppleDeviceActionResultActionState_None         DeleteUserFromSharedAppleDeviceActionResultActionState = "none"
	DeleteUserFromSharedAppleDeviceActionResultActionState_NotSupported DeleteUserFromSharedAppleDeviceActionResultActionState = "notSupported"
	DeleteUserFromSharedAppleDeviceActionResultActionState_Pending      DeleteUserFromSharedAppleDeviceActionResultActionState = "pending"
)

func PossibleValuesForDeleteUserFromSharedAppleDeviceActionResultActionState() []string {
	return []string{
		string(DeleteUserFromSharedAppleDeviceActionResultActionState_Active),
		string(DeleteUserFromSharedAppleDeviceActionResultActionState_Canceled),
		string(DeleteUserFromSharedAppleDeviceActionResultActionState_Done),
		string(DeleteUserFromSharedAppleDeviceActionResultActionState_Failed),
		string(DeleteUserFromSharedAppleDeviceActionResultActionState_None),
		string(DeleteUserFromSharedAppleDeviceActionResultActionState_NotSupported),
		string(DeleteUserFromSharedAppleDeviceActionResultActionState_Pending),
	}
}

func (s *DeleteUserFromSharedAppleDeviceActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeleteUserFromSharedAppleDeviceActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeleteUserFromSharedAppleDeviceActionResultActionState(input string) (*DeleteUserFromSharedAppleDeviceActionResultActionState, error) {
	vals := map[string]DeleteUserFromSharedAppleDeviceActionResultActionState{
		"active":       DeleteUserFromSharedAppleDeviceActionResultActionState_Active,
		"canceled":     DeleteUserFromSharedAppleDeviceActionResultActionState_Canceled,
		"done":         DeleteUserFromSharedAppleDeviceActionResultActionState_Done,
		"failed":       DeleteUserFromSharedAppleDeviceActionResultActionState_Failed,
		"none":         DeleteUserFromSharedAppleDeviceActionResultActionState_None,
		"notsupported": DeleteUserFromSharedAppleDeviceActionResultActionState_NotSupported,
		"pending":      DeleteUserFromSharedAppleDeviceActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeleteUserFromSharedAppleDeviceActionResultActionState(input)
	return &out, nil
}
