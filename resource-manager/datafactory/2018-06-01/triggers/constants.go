package triggers

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
