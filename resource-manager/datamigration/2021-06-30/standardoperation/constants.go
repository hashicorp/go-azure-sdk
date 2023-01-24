package standardoperation

import "strings"

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

func parseCommandState(input string) (*CommandState, error) {
	vals := map[string]CommandState{
		"accepted":  CommandStateAccepted,
		"failed":    CommandStateFailed,
		"running":   CommandStateRunning,
		"succeeded": CommandStateSucceeded,
		"unknown":   CommandStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommandState(input)
	return &out, nil
}

type NameCheckFailureReason string

const (
	NameCheckFailureReasonAlreadyExists NameCheckFailureReason = "AlreadyExists"
	NameCheckFailureReasonInvalid       NameCheckFailureReason = "Invalid"
)

func PossibleValuesForNameCheckFailureReason() []string {
	return []string{
		string(NameCheckFailureReasonAlreadyExists),
		string(NameCheckFailureReasonInvalid),
	}
}

func parseNameCheckFailureReason(input string) (*NameCheckFailureReason, error) {
	vals := map[string]NameCheckFailureReason{
		"alreadyexists": NameCheckFailureReasonAlreadyExists,
		"invalid":       NameCheckFailureReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NameCheckFailureReason(input)
	return &out, nil
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

func parseProjectProvisioningState(input string) (*ProjectProvisioningState, error) {
	vals := map[string]ProjectProvisioningState{
		"deleting":  ProjectProvisioningStateDeleting,
		"succeeded": ProjectProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectProvisioningState(input)
	return &out, nil
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

func parseProjectSourcePlatform(input string) (*ProjectSourcePlatform, error) {
	vals := map[string]ProjectSourcePlatform{
		"mongodb":    ProjectSourcePlatformMongoDb,
		"mysql":      ProjectSourcePlatformMySQL,
		"postgresql": ProjectSourcePlatformPostgreSql,
		"sql":        ProjectSourcePlatformSQL,
		"unknown":    ProjectSourcePlatformUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectSourcePlatform(input)
	return &out, nil
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

func parseProjectTargetPlatform(input string) (*ProjectTargetPlatform, error) {
	vals := map[string]ProjectTargetPlatform{
		"azuredbformysql":      ProjectTargetPlatformAzureDbForMySql,
		"azuredbforpostgresql": ProjectTargetPlatformAzureDbForPostgreSql,
		"mongodb":              ProjectTargetPlatformMongoDb,
		"sqldb":                ProjectTargetPlatformSQLDB,
		"sqlmi":                ProjectTargetPlatformSQLMI,
		"unknown":              ProjectTargetPlatformUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectTargetPlatform(input)
	return &out, nil
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

func parseResourceSkuCapacityScaleType(input string) (*ResourceSkuCapacityScaleType, error) {
	vals := map[string]ResourceSkuCapacityScaleType{
		"automatic": ResourceSkuCapacityScaleTypeAutomatic,
		"manual":    ResourceSkuCapacityScaleTypeManual,
		"none":      ResourceSkuCapacityScaleTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceSkuCapacityScaleType(input)
	return &out, nil
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

func parseResourceSkuRestrictionsReasonCode(input string) (*ResourceSkuRestrictionsReasonCode, error) {
	vals := map[string]ResourceSkuRestrictionsReasonCode{
		"notavailableforsubscription": ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription,
		"quotaid":                     ResourceSkuRestrictionsReasonCodeQuotaId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceSkuRestrictionsReasonCode(input)
	return &out, nil
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

func parseResourceSkuRestrictionsType(input string) (*ResourceSkuRestrictionsType, error) {
	vals := map[string]ResourceSkuRestrictionsType{
		"location": ResourceSkuRestrictionsTypeLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceSkuRestrictionsType(input)
	return &out, nil
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

func parseServiceProvisioningState(input string) (*ServiceProvisioningState, error) {
	vals := map[string]ServiceProvisioningState{
		"accepted":      ServiceProvisioningStateAccepted,
		"deleting":      ServiceProvisioningStateDeleting,
		"deploying":     ServiceProvisioningStateDeploying,
		"failed":        ServiceProvisioningStateFailed,
		"failedtostart": ServiceProvisioningStateFailedToStart,
		"failedtostop":  ServiceProvisioningStateFailedToStop,
		"starting":      ServiceProvisioningStateStarting,
		"stopped":       ServiceProvisioningStateStopped,
		"stopping":      ServiceProvisioningStateStopping,
		"succeeded":     ServiceProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceProvisioningState(input)
	return &out, nil
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

func parseServiceScalability(input string) (*ServiceScalability, error) {
	vals := map[string]ServiceScalability{
		"automatic": ServiceScalabilityAutomatic,
		"manual":    ServiceScalabilityManual,
		"none":      ServiceScalabilityNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceScalability(input)
	return &out, nil
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

func parseTaskState(input string) (*TaskState, error) {
	vals := map[string]TaskState{
		"canceled":              TaskStateCanceled,
		"failed":                TaskStateFailed,
		"failedinputvalidation": TaskStateFailedInputValidation,
		"faulted":               TaskStateFaulted,
		"queued":                TaskStateQueued,
		"running":               TaskStateRunning,
		"succeeded":             TaskStateSucceeded,
		"unknown":               TaskStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaskState(input)
	return &out, nil
}
