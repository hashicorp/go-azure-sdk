package applyupdate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateStatus string

const (
	UpdateStatusCompleted  UpdateStatus = "Completed"
	UpdateStatusInProgress UpdateStatus = "InProgress"
	UpdateStatusPending    UpdateStatus = "Pending"
	UpdateStatusRetryLater UpdateStatus = "RetryLater"
	UpdateStatusRetryNow   UpdateStatus = "RetryNow"
)

func PossibleValuesForUpdateStatus() []string {
	return []string{
		string(UpdateStatusCompleted),
		string(UpdateStatusInProgress),
		string(UpdateStatusPending),
		string(UpdateStatusRetryLater),
		string(UpdateStatusRetryNow),
	}
}
