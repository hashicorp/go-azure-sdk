package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivateDeviceEsimActionResultActionState string

const (
	ActivateDeviceEsimActionResultActionStateactive       ActivateDeviceEsimActionResultActionState = "Active"
	ActivateDeviceEsimActionResultActionStatecanceled     ActivateDeviceEsimActionResultActionState = "Canceled"
	ActivateDeviceEsimActionResultActionStatedone         ActivateDeviceEsimActionResultActionState = "Done"
	ActivateDeviceEsimActionResultActionStatefailed       ActivateDeviceEsimActionResultActionState = "Failed"
	ActivateDeviceEsimActionResultActionStatenone         ActivateDeviceEsimActionResultActionState = "None"
	ActivateDeviceEsimActionResultActionStatenotSupported ActivateDeviceEsimActionResultActionState = "NotSupported"
	ActivateDeviceEsimActionResultActionStatepending      ActivateDeviceEsimActionResultActionState = "Pending"
)

func PossibleValuesForActivateDeviceEsimActionResultActionState() []string {
	return []string{
		string(ActivateDeviceEsimActionResultActionStateactive),
		string(ActivateDeviceEsimActionResultActionStatecanceled),
		string(ActivateDeviceEsimActionResultActionStatedone),
		string(ActivateDeviceEsimActionResultActionStatefailed),
		string(ActivateDeviceEsimActionResultActionStatenone),
		string(ActivateDeviceEsimActionResultActionStatenotSupported),
		string(ActivateDeviceEsimActionResultActionStatepending),
	}
}

func parseActivateDeviceEsimActionResultActionState(input string) (*ActivateDeviceEsimActionResultActionState, error) {
	vals := map[string]ActivateDeviceEsimActionResultActionState{
		"active":       ActivateDeviceEsimActionResultActionStateactive,
		"canceled":     ActivateDeviceEsimActionResultActionStatecanceled,
		"done":         ActivateDeviceEsimActionResultActionStatedone,
		"failed":       ActivateDeviceEsimActionResultActionStatefailed,
		"none":         ActivateDeviceEsimActionResultActionStatenone,
		"notsupported": ActivateDeviceEsimActionResultActionStatenotSupported,
		"pending":      ActivateDeviceEsimActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActivateDeviceEsimActionResultActionState(input)
	return &out, nil
}
