package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerEventType string

const (
	TriggerEventTypeFileEvent          TriggerEventType = "FileEvent"
	TriggerEventTypePeriodicTimerEvent TriggerEventType = "PeriodicTimerEvent"
)

func PossibleValuesForTriggerEventType() []string {
	return []string{
		string(TriggerEventTypeFileEvent),
		string(TriggerEventTypePeriodicTimerEvent),
	}
}
