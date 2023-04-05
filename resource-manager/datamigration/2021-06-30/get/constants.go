package get

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommandState string

const (
	CommandStateAccepted  CommandState = "Accepted"
	CommandStateFailed    CommandState = "Failed"
	CommandStateRunning   CommandState = "Running"
	CommandStateSucceeded CommandState = "Succeeded"
	CommandStateUnknown   CommandState = "Unknown"
)

func PossibleValuesForCommandState() []string {
	return []string{
		string(CommandStateAccepted),
		string(CommandStateFailed),
		string(CommandStateRunning),
		string(CommandStateSucceeded),
		string(CommandStateUnknown),
	}
}

type ProjectProvisioningState string

const (
	ProjectProvisioningStateDeleting  ProjectProvisioningState = "Deleting"
	ProjectProvisioningStateSucceeded ProjectProvisioningState = "Succeeded"
)

func PossibleValuesForProjectProvisioningState() []string {
	return []string{
		string(ProjectProvisioningStateDeleting),
		string(ProjectProvisioningStateSucceeded),
	}
}

type ProjectSourcePlatform string

const (
	ProjectSourcePlatformMongoDb    ProjectSourcePlatform = "MongoDb"
	ProjectSourcePlatformMySQL      ProjectSourcePlatform = "MySQL"
	ProjectSourcePlatformPostgreSql ProjectSourcePlatform = "PostgreSql"
	ProjectSourcePlatformSQL        ProjectSourcePlatform = "SQL"
	ProjectSourcePlatformUnknown    ProjectSourcePlatform = "Unknown"
)

func PossibleValuesForProjectSourcePlatform() []string {
	return []string{
		string(ProjectSourcePlatformMongoDb),
		string(ProjectSourcePlatformMySQL),
		string(ProjectSourcePlatformPostgreSql),
		string(ProjectSourcePlatformSQL),
		string(ProjectSourcePlatformUnknown),
	}
}

type ProjectTargetPlatform string

const (
	ProjectTargetPlatformAzureDbForMySql      ProjectTargetPlatform = "AzureDbForMySql"
	ProjectTargetPlatformAzureDbForPostgreSql ProjectTargetPlatform = "AzureDbForPostgreSql"
	ProjectTargetPlatformMongoDb              ProjectTargetPlatform = "MongoDb"
	ProjectTargetPlatformSQLDB                ProjectTargetPlatform = "SQLDB"
	ProjectTargetPlatformSQLMI                ProjectTargetPlatform = "SQLMI"
	ProjectTargetPlatformUnknown              ProjectTargetPlatform = "Unknown"
)

func PossibleValuesForProjectTargetPlatform() []string {
	return []string{
		string(ProjectTargetPlatformAzureDbForMySql),
		string(ProjectTargetPlatformAzureDbForPostgreSql),
		string(ProjectTargetPlatformMongoDb),
		string(ProjectTargetPlatformSQLDB),
		string(ProjectTargetPlatformSQLMI),
		string(ProjectTargetPlatformUnknown),
	}
}

type ResourceSkuCapacityScaleType string

const (
	ResourceSkuCapacityScaleTypeAutomatic ResourceSkuCapacityScaleType = "Automatic"
	ResourceSkuCapacityScaleTypeManual    ResourceSkuCapacityScaleType = "Manual"
	ResourceSkuCapacityScaleTypeNone      ResourceSkuCapacityScaleType = "None"
)

func PossibleValuesForResourceSkuCapacityScaleType() []string {
	return []string{
		string(ResourceSkuCapacityScaleTypeAutomatic),
		string(ResourceSkuCapacityScaleTypeManual),
		string(ResourceSkuCapacityScaleTypeNone),
	}
}

type ResourceSkuRestrictionsReasonCode string

const (
	ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription ResourceSkuRestrictionsReasonCode = "NotAvailableForSubscription"
	ResourceSkuRestrictionsReasonCodeQuotaId                     ResourceSkuRestrictionsReasonCode = "QuotaId"
)

func PossibleValuesForResourceSkuRestrictionsReasonCode() []string {
	return []string{
		string(ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription),
		string(ResourceSkuRestrictionsReasonCodeQuotaId),
	}
}

type ResourceSkuRestrictionsType string

const (
	ResourceSkuRestrictionsTypeLocation ResourceSkuRestrictionsType = "location"
)

func PossibleValuesForResourceSkuRestrictionsType() []string {
	return []string{
		string(ResourceSkuRestrictionsTypeLocation),
	}
}

type ServiceProvisioningState string

const (
	ServiceProvisioningStateAccepted      ServiceProvisioningState = "Accepted"
	ServiceProvisioningStateDeleting      ServiceProvisioningState = "Deleting"
	ServiceProvisioningStateDeploying     ServiceProvisioningState = "Deploying"
	ServiceProvisioningStateFailed        ServiceProvisioningState = "Failed"
	ServiceProvisioningStateFailedToStart ServiceProvisioningState = "FailedToStart"
	ServiceProvisioningStateFailedToStop  ServiceProvisioningState = "FailedToStop"
	ServiceProvisioningStateStarting      ServiceProvisioningState = "Starting"
	ServiceProvisioningStateStopped       ServiceProvisioningState = "Stopped"
	ServiceProvisioningStateStopping      ServiceProvisioningState = "Stopping"
	ServiceProvisioningStateSucceeded     ServiceProvisioningState = "Succeeded"
)

func PossibleValuesForServiceProvisioningState() []string {
	return []string{
		string(ServiceProvisioningStateAccepted),
		string(ServiceProvisioningStateDeleting),
		string(ServiceProvisioningStateDeploying),
		string(ServiceProvisioningStateFailed),
		string(ServiceProvisioningStateFailedToStart),
		string(ServiceProvisioningStateFailedToStop),
		string(ServiceProvisioningStateStarting),
		string(ServiceProvisioningStateStopped),
		string(ServiceProvisioningStateStopping),
		string(ServiceProvisioningStateSucceeded),
	}
}

type ServiceScalability string

const (
	ServiceScalabilityAutomatic ServiceScalability = "automatic"
	ServiceScalabilityManual    ServiceScalability = "manual"
	ServiceScalabilityNone      ServiceScalability = "none"
)

func PossibleValuesForServiceScalability() []string {
	return []string{
		string(ServiceScalabilityAutomatic),
		string(ServiceScalabilityManual),
		string(ServiceScalabilityNone),
	}
}

type TaskState string

const (
	TaskStateCanceled              TaskState = "Canceled"
	TaskStateFailed                TaskState = "Failed"
	TaskStateFailedInputValidation TaskState = "FailedInputValidation"
	TaskStateFaulted               TaskState = "Faulted"
	TaskStateQueued                TaskState = "Queued"
	TaskStateRunning               TaskState = "Running"
	TaskStateSucceeded             TaskState = "Succeeded"
	TaskStateUnknown               TaskState = "Unknown"
)

func PossibleValuesForTaskState() []string {
	return []string{
		string(TaskStateCanceled),
		string(TaskStateFailed),
		string(TaskStateFailedInputValidation),
		string(TaskStateFaulted),
		string(TaskStateQueued),
		string(TaskStateRunning),
		string(TaskStateSucceeded),
		string(TaskStateUnknown),
	}
}
