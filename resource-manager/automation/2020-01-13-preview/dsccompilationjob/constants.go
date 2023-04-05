package dsccompilationjob

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobProvisioningState string

const (
	JobProvisioningStateFailed     JobProvisioningState = "Failed"
	JobProvisioningStateProcessing JobProvisioningState = "Processing"
	JobProvisioningStateSucceeded  JobProvisioningState = "Succeeded"
	JobProvisioningStateSuspended  JobProvisioningState = "Suspended"
)

func PossibleValuesForJobProvisioningState() []string {
	return []string{
		string(JobProvisioningStateFailed),
		string(JobProvisioningStateProcessing),
		string(JobProvisioningStateSucceeded),
		string(JobProvisioningStateSuspended),
	}
}

type JobStatus string

const (
	JobStatusActivating   JobStatus = "Activating"
	JobStatusBlocked      JobStatus = "Blocked"
	JobStatusCompleted    JobStatus = "Completed"
	JobStatusDisconnected JobStatus = "Disconnected"
	JobStatusFailed       JobStatus = "Failed"
	JobStatusNew          JobStatus = "New"
	JobStatusRemoving     JobStatus = "Removing"
	JobStatusResuming     JobStatus = "Resuming"
	JobStatusRunning      JobStatus = "Running"
	JobStatusStopped      JobStatus = "Stopped"
	JobStatusStopping     JobStatus = "Stopping"
	JobStatusSuspended    JobStatus = "Suspended"
	JobStatusSuspending   JobStatus = "Suspending"
)

func PossibleValuesForJobStatus() []string {
	return []string{
		string(JobStatusActivating),
		string(JobStatusBlocked),
		string(JobStatusCompleted),
		string(JobStatusDisconnected),
		string(JobStatusFailed),
		string(JobStatusNew),
		string(JobStatusRemoving),
		string(JobStatusResuming),
		string(JobStatusRunning),
		string(JobStatusStopped),
		string(JobStatusStopping),
		string(JobStatusSuspended),
		string(JobStatusSuspending),
	}
}

type JobStreamType string

const (
	JobStreamTypeAny      JobStreamType = "Any"
	JobStreamTypeDebug    JobStreamType = "Debug"
	JobStreamTypeError    JobStreamType = "Error"
	JobStreamTypeOutput   JobStreamType = "Output"
	JobStreamTypeProgress JobStreamType = "Progress"
	JobStreamTypeVerbose  JobStreamType = "Verbose"
	JobStreamTypeWarning  JobStreamType = "Warning"
)

func PossibleValuesForJobStreamType() []string {
	return []string{
		string(JobStreamTypeAny),
		string(JobStreamTypeDebug),
		string(JobStreamTypeError),
		string(JobStreamTypeOutput),
		string(JobStreamTypeProgress),
		string(JobStreamTypeVerbose),
		string(JobStreamTypeWarning),
	}
}
