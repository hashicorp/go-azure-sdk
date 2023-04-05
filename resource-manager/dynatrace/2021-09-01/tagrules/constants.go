package tagrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateNotSpecified),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type SendAadLogsStatus string

const (
	SendAadLogsStatusDisabled SendAadLogsStatus = "Disabled"
	SendAadLogsStatusEnabled  SendAadLogsStatus = "Enabled"
)

func PossibleValuesForSendAadLogsStatus() []string {
	return []string{
		string(SendAadLogsStatusDisabled),
		string(SendAadLogsStatusEnabled),
	}
}

type SendActivityLogsStatus string

const (
	SendActivityLogsStatusDisabled SendActivityLogsStatus = "Disabled"
	SendActivityLogsStatusEnabled  SendActivityLogsStatus = "Enabled"
)

func PossibleValuesForSendActivityLogsStatus() []string {
	return []string{
		string(SendActivityLogsStatusDisabled),
		string(SendActivityLogsStatusEnabled),
	}
}

type SendSubscriptionLogsStatus string

const (
	SendSubscriptionLogsStatusDisabled SendSubscriptionLogsStatus = "Disabled"
	SendSubscriptionLogsStatusEnabled  SendSubscriptionLogsStatus = "Enabled"
)

func PossibleValuesForSendSubscriptionLogsStatus() []string {
	return []string{
		string(SendSubscriptionLogsStatusDisabled),
		string(SendSubscriptionLogsStatusEnabled),
	}
}

type TagAction string

const (
	TagActionExclude TagAction = "Exclude"
	TagActionInclude TagAction = "Include"
)

func PossibleValuesForTagAction() []string {
	return []string{
		string(TagActionExclude),
		string(TagActionInclude),
	}
}
