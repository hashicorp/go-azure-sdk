package vnetpeering

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PeeringProvisioningState string

const (
	PeeringProvisioningStateDeleting  PeeringProvisioningState = "Deleting"
	PeeringProvisioningStateFailed    PeeringProvisioningState = "Failed"
	PeeringProvisioningStateSucceeded PeeringProvisioningState = "Succeeded"
	PeeringProvisioningStateUpdating  PeeringProvisioningState = "Updating"
)

func PossibleValuesForPeeringProvisioningState() []string {
	return []string{
		string(PeeringProvisioningStateDeleting),
		string(PeeringProvisioningStateFailed),
		string(PeeringProvisioningStateSucceeded),
		string(PeeringProvisioningStateUpdating),
	}
}

func (s *PeeringProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPeeringProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PeeringProvisioningState(decoded)
	return nil
}

type PeeringState string

const (
	PeeringStateConnected    PeeringState = "Connected"
	PeeringStateDisconnected PeeringState = "Disconnected"
	PeeringStateInitiated    PeeringState = "Initiated"
)

func PossibleValuesForPeeringState() []string {
	return []string{
		string(PeeringStateConnected),
		string(PeeringStateDisconnected),
		string(PeeringStateInitiated),
	}
}

func (s *PeeringState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPeeringState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PeeringState(decoded)
	return nil
}
