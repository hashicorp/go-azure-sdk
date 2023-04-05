package vmhost

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMHostUpdateStates string

const (
	VMHostUpdateStatesDelete  VMHostUpdateStates = "Delete"
	VMHostUpdateStatesInstall VMHostUpdateStates = "Install"
)

func PossibleValuesForVMHostUpdateStates() []string {
	return []string{
		string(VMHostUpdateStatesDelete),
		string(VMHostUpdateStatesInstall),
	}
}
