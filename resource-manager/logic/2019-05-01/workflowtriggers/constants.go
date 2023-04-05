package workflowtriggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DayOfWeek string

const (
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

func PossibleValuesForDayOfWeek() []string {
	return []string{
		string(DayOfWeekFriday),
		string(DayOfWeekMonday),
		string(DayOfWeekSaturday),
		string(DayOfWeekSunday),
		string(DayOfWeekThursday),
		string(DayOfWeekTuesday),
		string(DayOfWeekWednesday),
	}
}

type DaysOfWeek string

const (
	DaysOfWeekFriday    DaysOfWeek = "Friday"
	DaysOfWeekMonday    DaysOfWeek = "Monday"
	DaysOfWeekSaturday  DaysOfWeek = "Saturday"
	DaysOfWeekSunday    DaysOfWeek = "Sunday"
	DaysOfWeekThursday  DaysOfWeek = "Thursday"
	DaysOfWeekTuesday   DaysOfWeek = "Tuesday"
	DaysOfWeekWednesday DaysOfWeek = "Wednesday"
)

func PossibleValuesForDaysOfWeek() []string {
	return []string{
		string(DaysOfWeekFriday),
		string(DaysOfWeekMonday),
		string(DaysOfWeekSaturday),
		string(DaysOfWeekSunday),
		string(DaysOfWeekThursday),
		string(DaysOfWeekTuesday),
		string(DaysOfWeekWednesday),
	}
}

type KeyType string

const (
	KeyTypeNotSpecified KeyType = "NotSpecified"
	KeyTypePrimary      KeyType = "Primary"
	KeyTypeSecondary    KeyType = "Secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypeNotSpecified),
		string(KeyTypePrimary),
		string(KeyTypeSecondary),
	}
}

type RecurrenceFrequency string

const (
	RecurrenceFrequencyDay          RecurrenceFrequency = "Day"
	RecurrenceFrequencyHour         RecurrenceFrequency = "Hour"
	RecurrenceFrequencyMinute       RecurrenceFrequency = "Minute"
	RecurrenceFrequencyMonth        RecurrenceFrequency = "Month"
	RecurrenceFrequencyNotSpecified RecurrenceFrequency = "NotSpecified"
	RecurrenceFrequencySecond       RecurrenceFrequency = "Second"
	RecurrenceFrequencyWeek         RecurrenceFrequency = "Week"
	RecurrenceFrequencyYear         RecurrenceFrequency = "Year"
)

func PossibleValuesForRecurrenceFrequency() []string {
	return []string{
		string(RecurrenceFrequencyDay),
		string(RecurrenceFrequencyHour),
		string(RecurrenceFrequencyMinute),
		string(RecurrenceFrequencyMonth),
		string(RecurrenceFrequencyNotSpecified),
		string(RecurrenceFrequencySecond),
		string(RecurrenceFrequencyWeek),
		string(RecurrenceFrequencyYear),
	}
}

type WorkflowState string

const (
	WorkflowStateCompleted    WorkflowState = "Completed"
	WorkflowStateDeleted      WorkflowState = "Deleted"
	WorkflowStateDisabled     WorkflowState = "Disabled"
	WorkflowStateEnabled      WorkflowState = "Enabled"
	WorkflowStateNotSpecified WorkflowState = "NotSpecified"
	WorkflowStateSuspended    WorkflowState = "Suspended"
)

func PossibleValuesForWorkflowState() []string {
	return []string{
		string(WorkflowStateCompleted),
		string(WorkflowStateDeleted),
		string(WorkflowStateDisabled),
		string(WorkflowStateEnabled),
		string(WorkflowStateNotSpecified),
		string(WorkflowStateSuspended),
	}
}

type WorkflowStatus string

const (
	WorkflowStatusAborted      WorkflowStatus = "Aborted"
	WorkflowStatusCancelled    WorkflowStatus = "Cancelled"
	WorkflowStatusFailed       WorkflowStatus = "Failed"
	WorkflowStatusFaulted      WorkflowStatus = "Faulted"
	WorkflowStatusIgnored      WorkflowStatus = "Ignored"
	WorkflowStatusNotSpecified WorkflowStatus = "NotSpecified"
	WorkflowStatusPaused       WorkflowStatus = "Paused"
	WorkflowStatusRunning      WorkflowStatus = "Running"
	WorkflowStatusSkipped      WorkflowStatus = "Skipped"
	WorkflowStatusSucceeded    WorkflowStatus = "Succeeded"
	WorkflowStatusSuspended    WorkflowStatus = "Suspended"
	WorkflowStatusTimedOut     WorkflowStatus = "TimedOut"
	WorkflowStatusWaiting      WorkflowStatus = "Waiting"
)

func PossibleValuesForWorkflowStatus() []string {
	return []string{
		string(WorkflowStatusAborted),
		string(WorkflowStatusCancelled),
		string(WorkflowStatusFailed),
		string(WorkflowStatusFaulted),
		string(WorkflowStatusIgnored),
		string(WorkflowStatusNotSpecified),
		string(WorkflowStatusPaused),
		string(WorkflowStatusRunning),
		string(WorkflowStatusSkipped),
		string(WorkflowStatusSucceeded),
		string(WorkflowStatusSuspended),
		string(WorkflowStatusTimedOut),
		string(WorkflowStatusWaiting),
	}
}

type WorkflowTriggerProvisioningState string

const (
	WorkflowTriggerProvisioningStateAccepted      WorkflowTriggerProvisioningState = "Accepted"
	WorkflowTriggerProvisioningStateCanceled      WorkflowTriggerProvisioningState = "Canceled"
	WorkflowTriggerProvisioningStateCompleted     WorkflowTriggerProvisioningState = "Completed"
	WorkflowTriggerProvisioningStateCreated       WorkflowTriggerProvisioningState = "Created"
	WorkflowTriggerProvisioningStateCreating      WorkflowTriggerProvisioningState = "Creating"
	WorkflowTriggerProvisioningStateDeleted       WorkflowTriggerProvisioningState = "Deleted"
	WorkflowTriggerProvisioningStateDeleting      WorkflowTriggerProvisioningState = "Deleting"
	WorkflowTriggerProvisioningStateFailed        WorkflowTriggerProvisioningState = "Failed"
	WorkflowTriggerProvisioningStateMoving        WorkflowTriggerProvisioningState = "Moving"
	WorkflowTriggerProvisioningStateNotSpecified  WorkflowTriggerProvisioningState = "NotSpecified"
	WorkflowTriggerProvisioningStateReady         WorkflowTriggerProvisioningState = "Ready"
	WorkflowTriggerProvisioningStateRegistered    WorkflowTriggerProvisioningState = "Registered"
	WorkflowTriggerProvisioningStateRegistering   WorkflowTriggerProvisioningState = "Registering"
	WorkflowTriggerProvisioningStateRunning       WorkflowTriggerProvisioningState = "Running"
	WorkflowTriggerProvisioningStateSucceeded     WorkflowTriggerProvisioningState = "Succeeded"
	WorkflowTriggerProvisioningStateUnregistered  WorkflowTriggerProvisioningState = "Unregistered"
	WorkflowTriggerProvisioningStateUnregistering WorkflowTriggerProvisioningState = "Unregistering"
	WorkflowTriggerProvisioningStateUpdating      WorkflowTriggerProvisioningState = "Updating"
)

func PossibleValuesForWorkflowTriggerProvisioningState() []string {
	return []string{
		string(WorkflowTriggerProvisioningStateAccepted),
		string(WorkflowTriggerProvisioningStateCanceled),
		string(WorkflowTriggerProvisioningStateCompleted),
		string(WorkflowTriggerProvisioningStateCreated),
		string(WorkflowTriggerProvisioningStateCreating),
		string(WorkflowTriggerProvisioningStateDeleted),
		string(WorkflowTriggerProvisioningStateDeleting),
		string(WorkflowTriggerProvisioningStateFailed),
		string(WorkflowTriggerProvisioningStateMoving),
		string(WorkflowTriggerProvisioningStateNotSpecified),
		string(WorkflowTriggerProvisioningStateReady),
		string(WorkflowTriggerProvisioningStateRegistered),
		string(WorkflowTriggerProvisioningStateRegistering),
		string(WorkflowTriggerProvisioningStateRunning),
		string(WorkflowTriggerProvisioningStateSucceeded),
		string(WorkflowTriggerProvisioningStateUnregistered),
		string(WorkflowTriggerProvisioningStateUnregistering),
		string(WorkflowTriggerProvisioningStateUpdating),
	}
}
