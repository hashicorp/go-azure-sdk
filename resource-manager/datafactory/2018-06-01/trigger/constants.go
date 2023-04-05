package trigger

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
