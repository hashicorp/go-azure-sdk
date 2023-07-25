package kafkaconfiguration

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialsType string

const (
	CredentialsTypeNone           CredentialsType = "None"
	CredentialsTypeSystemAssigned CredentialsType = "SystemAssigned"
	CredentialsTypeUserAssigned   CredentialsType = "UserAssigned"
)

func PossibleValuesForCredentialsType() []string {
	return []string{
		string(CredentialsTypeNone),
		string(CredentialsTypeSystemAssigned),
		string(CredentialsTypeUserAssigned),
	}
}

func (s *CredentialsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCredentialsType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCredentialsType(input string) (*CredentialsType, error) {
	vals := map[string]CredentialsType{
		"none":           CredentialsTypeNone,
		"systemassigned": CredentialsTypeSystemAssigned,
		"userassigned":   CredentialsTypeUserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CredentialsType(input)
	return &out, nil
}

type EventHubType string

const (
	EventHubTypeHook         EventHubType = "Hook"
	EventHubTypeNotification EventHubType = "Notification"
)

func PossibleValuesForEventHubType() []string {
	return []string{
		string(EventHubTypeHook),
		string(EventHubTypeNotification),
	}
}

func (s *EventHubType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventHubType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventHubType(input string) (*EventHubType, error) {
	vals := map[string]EventHubType{
		"hook":         EventHubTypeHook,
		"notification": EventHubTypeNotification,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventHubType(input)
	return &out, nil
}

type EventStreamingState string

const (
	EventStreamingStateDisabled EventStreamingState = "Disabled"
	EventStreamingStateEnabled  EventStreamingState = "Enabled"
)

func PossibleValuesForEventStreamingState() []string {
	return []string{
		string(EventStreamingStateDisabled),
		string(EventStreamingStateEnabled),
	}
}

func (s *EventStreamingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventStreamingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventStreamingState(input string) (*EventStreamingState, error) {
	vals := map[string]EventStreamingState{
		"disabled": EventStreamingStateDisabled,
		"enabled":  EventStreamingStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventStreamingState(input)
	return &out, nil
}

type EventStreamingType string

const (
	EventStreamingTypeAzure   EventStreamingType = "Azure"
	EventStreamingTypeManaged EventStreamingType = "Managed"
	EventStreamingTypeNone    EventStreamingType = "None"
)

func PossibleValuesForEventStreamingType() []string {
	return []string{
		string(EventStreamingTypeAzure),
		string(EventStreamingTypeManaged),
		string(EventStreamingTypeNone),
	}
}

func (s *EventStreamingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventStreamingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventStreamingType(input string) (*EventStreamingType, error) {
	vals := map[string]EventStreamingType{
		"azure":   EventStreamingTypeAzure,
		"managed": EventStreamingTypeManaged,
		"none":    EventStreamingTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventStreamingType(input)
	return &out, nil
}
