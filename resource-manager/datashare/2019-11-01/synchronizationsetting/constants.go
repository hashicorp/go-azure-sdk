package synchronizationsetting

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateSucceeded),
	}
}

type RecurrenceInterval string

const (
	RecurrenceIntervalDay  RecurrenceInterval = "Day"
	RecurrenceIntervalHour RecurrenceInterval = "Hour"
)

func PossibleValuesForRecurrenceInterval() []string {
	return []string{
		string(RecurrenceIntervalDay),
		string(RecurrenceIntervalHour),
	}
}

type Status string

const (
	StatusAccepted         Status = "Accepted"
	StatusCanceled         Status = "Canceled"
	StatusFailed           Status = "Failed"
	StatusInProgress       Status = "InProgress"
	StatusSucceeded        Status = "Succeeded"
	StatusTransientFailure Status = "TransientFailure"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusAccepted),
		string(StatusCanceled),
		string(StatusFailed),
		string(StatusInProgress),
		string(StatusSucceeded),
		string(StatusTransientFailure),
	}
}

type SynchronizationSettingKind string

const (
	SynchronizationSettingKindScheduleBased SynchronizationSettingKind = "ScheduleBased"
)

func PossibleValuesForSynchronizationSettingKind() []string {
	return []string{
		string(SynchronizationSettingKindScheduleBased),
	}
}
