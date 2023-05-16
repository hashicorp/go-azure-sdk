package triggers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventSubscriptionStatus string

const (
	EventSubscriptionStatusDeprovisioning EventSubscriptionStatus = "Deprovisioning"
	EventSubscriptionStatusDisabled       EventSubscriptionStatus = "Disabled"
	EventSubscriptionStatusEnabled        EventSubscriptionStatus = "Enabled"
	EventSubscriptionStatusProvisioning   EventSubscriptionStatus = "Provisioning"
	EventSubscriptionStatusUnknown        EventSubscriptionStatus = "Unknown"
)

func PossibleValuesForEventSubscriptionStatus() []string {
	return []string{
		string(EventSubscriptionStatusDeprovisioning),
		string(EventSubscriptionStatusDisabled),
		string(EventSubscriptionStatusEnabled),
		string(EventSubscriptionStatusProvisioning),
		string(EventSubscriptionStatusUnknown),
	}
}

func (s *EventSubscriptionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventSubscriptionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventSubscriptionStatus(input string) (*EventSubscriptionStatus, error) {
	vals := map[string]EventSubscriptionStatus{
		"deprovisioning": EventSubscriptionStatusDeprovisioning,
		"disabled":       EventSubscriptionStatusDisabled,
		"enabled":        EventSubscriptionStatusEnabled,
		"provisioning":   EventSubscriptionStatusProvisioning,
		"unknown":        EventSubscriptionStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventSubscriptionStatus(input)
	return &out, nil
}

type TriggerRuntimeState string

const (
	TriggerRuntimeStateDisabled TriggerRuntimeState = "Disabled"
	TriggerRuntimeStateStarted  TriggerRuntimeState = "Started"
	TriggerRuntimeStateStopped  TriggerRuntimeState = "Stopped"
)

func PossibleValuesForTriggerRuntimeState() []string {
	return []string{
		string(TriggerRuntimeStateDisabled),
		string(TriggerRuntimeStateStarted),
		string(TriggerRuntimeStateStopped),
	}
}

func (s *TriggerRuntimeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerRuntimeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerRuntimeState(input string) (*TriggerRuntimeState, error) {
	vals := map[string]TriggerRuntimeState{
		"disabled": TriggerRuntimeStateDisabled,
		"started":  TriggerRuntimeStateStarted,
		"stopped":  TriggerRuntimeStateStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerRuntimeState(input)
	return &out, nil
}
