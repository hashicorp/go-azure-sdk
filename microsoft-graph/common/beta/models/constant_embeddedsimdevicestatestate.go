package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMDeviceStateState string

const (
	EmbeddedSIMDeviceStateStatedeleted       EmbeddedSIMDeviceStateState = "Deleted"
	EmbeddedSIMDeviceStateStatedeleting      EmbeddedSIMDeviceStateState = "Deleting"
	EmbeddedSIMDeviceStateStateerror         EmbeddedSIMDeviceStateState = "Error"
	EmbeddedSIMDeviceStateStatefailed        EmbeddedSIMDeviceStateState = "Failed"
	EmbeddedSIMDeviceStateStateinstalled     EmbeddedSIMDeviceStateState = "Installed"
	EmbeddedSIMDeviceStateStateinstalling    EmbeddedSIMDeviceStateState = "Installing"
	EmbeddedSIMDeviceStateStatenotEvaluated  EmbeddedSIMDeviceStateState = "NotEvaluated"
	EmbeddedSIMDeviceStateStateremovedByUser EmbeddedSIMDeviceStateState = "RemovedByUser"
)

func PossibleValuesForEmbeddedSIMDeviceStateState() []string {
	return []string{
		string(EmbeddedSIMDeviceStateStatedeleted),
		string(EmbeddedSIMDeviceStateStatedeleting),
		string(EmbeddedSIMDeviceStateStateerror),
		string(EmbeddedSIMDeviceStateStatefailed),
		string(EmbeddedSIMDeviceStateStateinstalled),
		string(EmbeddedSIMDeviceStateStateinstalling),
		string(EmbeddedSIMDeviceStateStatenotEvaluated),
		string(EmbeddedSIMDeviceStateStateremovedByUser),
	}
}

func parseEmbeddedSIMDeviceStateState(input string) (*EmbeddedSIMDeviceStateState, error) {
	vals := map[string]EmbeddedSIMDeviceStateState{
		"deleted":       EmbeddedSIMDeviceStateStatedeleted,
		"deleting":      EmbeddedSIMDeviceStateStatedeleting,
		"error":         EmbeddedSIMDeviceStateStateerror,
		"failed":        EmbeddedSIMDeviceStateStatefailed,
		"installed":     EmbeddedSIMDeviceStateStateinstalled,
		"installing":    EmbeddedSIMDeviceStateStateinstalling,
		"notevaluated":  EmbeddedSIMDeviceStateStatenotEvaluated,
		"removedbyuser": EmbeddedSIMDeviceStateStateremovedByUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmbeddedSIMDeviceStateState(input)
	return &out, nil
}
