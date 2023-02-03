package machines

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualDiskMode string

const (
	VirtualDiskModeAppend                   VirtualDiskMode = "append"
	VirtualDiskModeIndependentNonpersistent VirtualDiskMode = "independent_nonpersistent"
	VirtualDiskModeIndependentPersistent    VirtualDiskMode = "independent_persistent"
	VirtualDiskModeNonpersistent            VirtualDiskMode = "nonpersistent"
	VirtualDiskModePersistent               VirtualDiskMode = "persistent"
	VirtualDiskModeUndoable                 VirtualDiskMode = "undoable"
)

func PossibleValuesForVirtualDiskMode() []string {
	return []string{
		string(VirtualDiskModeAppend),
		string(VirtualDiskModeIndependentNonpersistent),
		string(VirtualDiskModeIndependentPersistent),
		string(VirtualDiskModeNonpersistent),
		string(VirtualDiskModePersistent),
		string(VirtualDiskModeUndoable),
	}
}

func parseVirtualDiskMode(input string) (*VirtualDiskMode, error) {
	vals := map[string]VirtualDiskMode{
		"append":                    VirtualDiskModeAppend,
		"independent_nonpersistent": VirtualDiskModeIndependentNonpersistent,
		"independent_persistent":    VirtualDiskModeIndependentPersistent,
		"nonpersistent":             VirtualDiskModeNonpersistent,
		"persistent":                VirtualDiskModePersistent,
		"undoable":                  VirtualDiskModeUndoable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualDiskMode(input)
	return &out, nil
}
