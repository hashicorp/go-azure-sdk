package partnerdestinations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerDestinationActivationState string

const (
	PartnerDestinationActivationStateActivated      PartnerDestinationActivationState = "Activated"
	PartnerDestinationActivationStateNeverActivated PartnerDestinationActivationState = "NeverActivated"
)

func PossibleValuesForPartnerDestinationActivationState() []string {
	return []string{
		string(PartnerDestinationActivationStateActivated),
		string(PartnerDestinationActivationStateNeverActivated),
	}
}

func (s *PartnerDestinationActivationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerDestinationActivationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerDestinationActivationState(input string) (*PartnerDestinationActivationState, error) {
	vals := map[string]PartnerDestinationActivationState{
		"activated":      PartnerDestinationActivationStateActivated,
		"neveractivated": PartnerDestinationActivationStateNeverActivated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerDestinationActivationState(input)
	return &out, nil
}

type PartnerDestinationProvisioningState string

const (
	PartnerDestinationProvisioningStateCanceled                                 PartnerDestinationProvisioningState = "Canceled"
	PartnerDestinationProvisioningStateCreating                                 PartnerDestinationProvisioningState = "Creating"
	PartnerDestinationProvisioningStateDeleting                                 PartnerDestinationProvisioningState = "Deleting"
	PartnerDestinationProvisioningStateFailed                                   PartnerDestinationProvisioningState = "Failed"
	PartnerDestinationProvisioningStateIdleDueToMirroredChannelResourceDeletion PartnerDestinationProvisioningState = "IdleDueToMirroredChannelResourceDeletion"
	PartnerDestinationProvisioningStateSucceeded                                PartnerDestinationProvisioningState = "Succeeded"
	PartnerDestinationProvisioningStateUpdating                                 PartnerDestinationProvisioningState = "Updating"
)

func PossibleValuesForPartnerDestinationProvisioningState() []string {
	return []string{
		string(PartnerDestinationProvisioningStateCanceled),
		string(PartnerDestinationProvisioningStateCreating),
		string(PartnerDestinationProvisioningStateDeleting),
		string(PartnerDestinationProvisioningStateFailed),
		string(PartnerDestinationProvisioningStateIdleDueToMirroredChannelResourceDeletion),
		string(PartnerDestinationProvisioningStateSucceeded),
		string(PartnerDestinationProvisioningStateUpdating),
	}
}

func (s *PartnerDestinationProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerDestinationProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerDestinationProvisioningState(input string) (*PartnerDestinationProvisioningState, error) {
	vals := map[string]PartnerDestinationProvisioningState{
		"canceled": PartnerDestinationProvisioningStateCanceled,
		"creating": PartnerDestinationProvisioningStateCreating,
		"deleting": PartnerDestinationProvisioningStateDeleting,
		"failed":   PartnerDestinationProvisioningStateFailed,
		"idleduetomirroredchannelresourcedeletion": PartnerDestinationProvisioningStateIdleDueToMirroredChannelResourceDeletion,
		"succeeded": PartnerDestinationProvisioningStateSucceeded,
		"updating":  PartnerDestinationProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerDestinationProvisioningState(input)
	return &out, nil
}
