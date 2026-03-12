package msdeploystatuses

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MSDeployLogEntryType string

const (
	MSDeployLogEntryTypeError   MSDeployLogEntryType = "Error"
	MSDeployLogEntryTypeMessage MSDeployLogEntryType = "Message"
	MSDeployLogEntryTypeWarning MSDeployLogEntryType = "Warning"
)

func PossibleValuesForMSDeployLogEntryType() []string {
	return []string{
		string(MSDeployLogEntryTypeError),
		string(MSDeployLogEntryTypeMessage),
		string(MSDeployLogEntryTypeWarning),
	}
}

func (s *MSDeployLogEntryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMSDeployLogEntryType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMSDeployLogEntryType(input string) (*MSDeployLogEntryType, error) {
	vals := map[string]MSDeployLogEntryType{
		"error":   MSDeployLogEntryTypeError,
		"message": MSDeployLogEntryTypeMessage,
		"warning": MSDeployLogEntryTypeWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MSDeployLogEntryType(input)
	return &out, nil
}

type MSDeployProvisioningState string

const (
	MSDeployProvisioningStateAccepted  MSDeployProvisioningState = "accepted"
	MSDeployProvisioningStateCanceled  MSDeployProvisioningState = "canceled"
	MSDeployProvisioningStateFailed    MSDeployProvisioningState = "failed"
	MSDeployProvisioningStateRunning   MSDeployProvisioningState = "running"
	MSDeployProvisioningStateSucceeded MSDeployProvisioningState = "succeeded"
)

func PossibleValuesForMSDeployProvisioningState() []string {
	return []string{
		string(MSDeployProvisioningStateAccepted),
		string(MSDeployProvisioningStateCanceled),
		string(MSDeployProvisioningStateFailed),
		string(MSDeployProvisioningStateRunning),
		string(MSDeployProvisioningStateSucceeded),
	}
}

func (s *MSDeployProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMSDeployProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMSDeployProvisioningState(input string) (*MSDeployProvisioningState, error) {
	vals := map[string]MSDeployProvisioningState{
		"accepted":  MSDeployProvisioningStateAccepted,
		"canceled":  MSDeployProvisioningStateCanceled,
		"failed":    MSDeployProvisioningStateFailed,
		"running":   MSDeployProvisioningStateRunning,
		"succeeded": MSDeployProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MSDeployProvisioningState(input)
	return &out, nil
}
