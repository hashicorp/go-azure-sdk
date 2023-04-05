package clusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterProvisioningState string

const (
	ClusterProvisioningStateCanceled   ClusterProvisioningState = "Canceled"
	ClusterProvisioningStateFailed     ClusterProvisioningState = "Failed"
	ClusterProvisioningStateInProgress ClusterProvisioningState = "InProgress"
	ClusterProvisioningStateSucceeded  ClusterProvisioningState = "Succeeded"
)

func PossibleValuesForClusterProvisioningState() []string {
	return []string{
		string(ClusterProvisioningStateCanceled),
		string(ClusterProvisioningStateFailed),
		string(ClusterProvisioningStateInProgress),
		string(ClusterProvisioningStateSucceeded),
	}
}

type ClusterSkuName string

const (
	ClusterSkuNameDefault ClusterSkuName = "Default"
)

func PossibleValuesForClusterSkuName() []string {
	return []string{
		string(ClusterSkuNameDefault),
	}
}

type JobState string

const (
	JobStateCreated    JobState = "Created"
	JobStateDegraded   JobState = "Degraded"
	JobStateDeleting   JobState = "Deleting"
	JobStateFailed     JobState = "Failed"
	JobStateRestarting JobState = "Restarting"
	JobStateRunning    JobState = "Running"
	JobStateScaling    JobState = "Scaling"
	JobStateStarting   JobState = "Starting"
	JobStateStopped    JobState = "Stopped"
	JobStateStopping   JobState = "Stopping"
)

func PossibleValuesForJobState() []string {
	return []string{
		string(JobStateCreated),
		string(JobStateDegraded),
		string(JobStateDeleting),
		string(JobStateFailed),
		string(JobStateRestarting),
		string(JobStateRunning),
		string(JobStateScaling),
		string(JobStateStarting),
		string(JobStateStopped),
		string(JobStateStopping),
	}
}
