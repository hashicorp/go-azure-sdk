package machines

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
