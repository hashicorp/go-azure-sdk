package globalreachconnections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalReachConnectionProvisioningState string

const (
	GlobalReachConnectionProvisioningStateCanceled  GlobalReachConnectionProvisioningState = "Canceled"
	GlobalReachConnectionProvisioningStateFailed    GlobalReachConnectionProvisioningState = "Failed"
	GlobalReachConnectionProvisioningStateSucceeded GlobalReachConnectionProvisioningState = "Succeeded"
	GlobalReachConnectionProvisioningStateUpdating  GlobalReachConnectionProvisioningState = "Updating"
)

func PossibleValuesForGlobalReachConnectionProvisioningState() []string {
	return []string{
		string(GlobalReachConnectionProvisioningStateCanceled),
		string(GlobalReachConnectionProvisioningStateFailed),
		string(GlobalReachConnectionProvisioningStateSucceeded),
		string(GlobalReachConnectionProvisioningStateUpdating),
	}
}

func (s *GlobalReachConnectionProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGlobalReachConnectionProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGlobalReachConnectionProvisioningState(input string) (*GlobalReachConnectionProvisioningState, error) {
	vals := map[string]GlobalReachConnectionProvisioningState{
		"canceled":  GlobalReachConnectionProvisioningStateCanceled,
		"failed":    GlobalReachConnectionProvisioningStateFailed,
		"succeeded": GlobalReachConnectionProvisioningStateSucceeded,
		"updating":  GlobalReachConnectionProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GlobalReachConnectionProvisioningState(input)
	return &out, nil
}

type GlobalReachConnectionStatus string

const (
	GlobalReachConnectionStatusConnected    GlobalReachConnectionStatus = "Connected"
	GlobalReachConnectionStatusConnecting   GlobalReachConnectionStatus = "Connecting"
	GlobalReachConnectionStatusDisconnected GlobalReachConnectionStatus = "Disconnected"
)

func PossibleValuesForGlobalReachConnectionStatus() []string {
	return []string{
		string(GlobalReachConnectionStatusConnected),
		string(GlobalReachConnectionStatusConnecting),
		string(GlobalReachConnectionStatusDisconnected),
	}
}

func (s *GlobalReachConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGlobalReachConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGlobalReachConnectionStatus(input string) (*GlobalReachConnectionStatus, error) {
	vals := map[string]GlobalReachConnectionStatus{
		"connected":    GlobalReachConnectionStatusConnected,
		"connecting":   GlobalReachConnectionStatusConnecting,
		"disconnected": GlobalReachConnectionStatusDisconnected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GlobalReachConnectionStatus(input)
	return &out, nil
}
