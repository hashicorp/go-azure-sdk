package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMDeviceStateState string

const (
	EmbeddedSIMDeviceStateState_Deleted       EmbeddedSIMDeviceStateState = "deleted"
	EmbeddedSIMDeviceStateState_Deleting      EmbeddedSIMDeviceStateState = "deleting"
	EmbeddedSIMDeviceStateState_Error         EmbeddedSIMDeviceStateState = "error"
	EmbeddedSIMDeviceStateState_Failed        EmbeddedSIMDeviceStateState = "failed"
	EmbeddedSIMDeviceStateState_Installed     EmbeddedSIMDeviceStateState = "installed"
	EmbeddedSIMDeviceStateState_Installing    EmbeddedSIMDeviceStateState = "installing"
	EmbeddedSIMDeviceStateState_NotEvaluated  EmbeddedSIMDeviceStateState = "notEvaluated"
	EmbeddedSIMDeviceStateState_RemovedByUser EmbeddedSIMDeviceStateState = "removedByUser"
)

func PossibleValuesForEmbeddedSIMDeviceStateState() []string {
	return []string{
		string(EmbeddedSIMDeviceStateState_Deleted),
		string(EmbeddedSIMDeviceStateState_Deleting),
		string(EmbeddedSIMDeviceStateState_Error),
		string(EmbeddedSIMDeviceStateState_Failed),
		string(EmbeddedSIMDeviceStateState_Installed),
		string(EmbeddedSIMDeviceStateState_Installing),
		string(EmbeddedSIMDeviceStateState_NotEvaluated),
		string(EmbeddedSIMDeviceStateState_RemovedByUser),
	}
}

func (s *EmbeddedSIMDeviceStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmbeddedSIMDeviceStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmbeddedSIMDeviceStateState(input string) (*EmbeddedSIMDeviceStateState, error) {
	vals := map[string]EmbeddedSIMDeviceStateState{
		"deleted":       EmbeddedSIMDeviceStateState_Deleted,
		"deleting":      EmbeddedSIMDeviceStateState_Deleting,
		"error":         EmbeddedSIMDeviceStateState_Error,
		"failed":        EmbeddedSIMDeviceStateState_Failed,
		"installed":     EmbeddedSIMDeviceStateState_Installed,
		"installing":    EmbeddedSIMDeviceStateState_Installing,
		"notevaluated":  EmbeddedSIMDeviceStateState_NotEvaluated,
		"removedbyuser": EmbeddedSIMDeviceStateState_RemovedByUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmbeddedSIMDeviceStateState(input)
	return &out, nil
}
