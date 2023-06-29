package kafkaconfiguration

import "strings"

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
