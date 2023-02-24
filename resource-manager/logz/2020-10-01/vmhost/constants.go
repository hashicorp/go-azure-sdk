package vmhost

import "strings"

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

func parseVMHostUpdateStates(input string) (*VMHostUpdateStates, error) {
	vals := map[string]VMHostUpdateStates{
		"delete":  VMHostUpdateStatesDelete,
		"install": VMHostUpdateStatesInstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VMHostUpdateStates(input)
	return &out, nil
}
