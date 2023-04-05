package replicationrecoveryplans

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type A2ARpRecoveryPointType string

const (
	A2ARpRecoveryPointTypeLatest                      A2ARpRecoveryPointType = "Latest"
	A2ARpRecoveryPointTypeLatestApplicationConsistent A2ARpRecoveryPointType = "LatestApplicationConsistent"
	A2ARpRecoveryPointTypeLatestCrashConsistent       A2ARpRecoveryPointType = "LatestCrashConsistent"
	A2ARpRecoveryPointTypeLatestProcessed             A2ARpRecoveryPointType = "LatestProcessed"
)

func PossibleValuesForA2ARpRecoveryPointType() []string {
	return []string{
		string(A2ARpRecoveryPointTypeLatest),
		string(A2ARpRecoveryPointTypeLatestApplicationConsistent),
		string(A2ARpRecoveryPointTypeLatestCrashConsistent),
		string(A2ARpRecoveryPointTypeLatestProcessed),
	}
}

type AlternateLocationRecoveryOption string

const (
	AlternateLocationRecoveryOptionCreateVMIfNotFound AlternateLocationRecoveryOption = "CreateVmIfNotFound"
	AlternateLocationRecoveryOptionNoAction           AlternateLocationRecoveryOption = "NoAction"
)

func PossibleValuesForAlternateLocationRecoveryOption() []string {
	return []string{
		string(AlternateLocationRecoveryOptionCreateVMIfNotFound),
		string(AlternateLocationRecoveryOptionNoAction),
	}
}

type DataSyncStatus string

const (
	DataSyncStatusForDownTime        DataSyncStatus = "ForDownTime"
	DataSyncStatusForSynchronization DataSyncStatus = "ForSynchronization"
)

func PossibleValuesForDataSyncStatus() []string {
	return []string{
		string(DataSyncStatusForDownTime),
		string(DataSyncStatusForSynchronization),
	}
}

type FailoverDeploymentModel string

const (
	FailoverDeploymentModelClassic         FailoverDeploymentModel = "Classic"
	FailoverDeploymentModelNotApplicable   FailoverDeploymentModel = "NotApplicable"
	FailoverDeploymentModelResourceManager FailoverDeploymentModel = "ResourceManager"
)

func PossibleValuesForFailoverDeploymentModel() []string {
	return []string{
		string(FailoverDeploymentModelClassic),
		string(FailoverDeploymentModelNotApplicable),
		string(FailoverDeploymentModelResourceManager),
	}
}

type HyperVReplicaAzureRpRecoveryPointType string

const (
	HyperVReplicaAzureRpRecoveryPointTypeLatest                      HyperVReplicaAzureRpRecoveryPointType = "Latest"
	HyperVReplicaAzureRpRecoveryPointTypeLatestApplicationConsistent HyperVReplicaAzureRpRecoveryPointType = "LatestApplicationConsistent"
	HyperVReplicaAzureRpRecoveryPointTypeLatestProcessed             HyperVReplicaAzureRpRecoveryPointType = "LatestProcessed"
)

func PossibleValuesForHyperVReplicaAzureRpRecoveryPointType() []string {
	return []string{
		string(HyperVReplicaAzureRpRecoveryPointTypeLatest),
		string(HyperVReplicaAzureRpRecoveryPointTypeLatestApplicationConsistent),
		string(HyperVReplicaAzureRpRecoveryPointTypeLatestProcessed),
	}
}

type InMageRcmFailbackRecoveryPointType string

const (
	InMageRcmFailbackRecoveryPointTypeApplicationConsistent InMageRcmFailbackRecoveryPointType = "ApplicationConsistent"
	InMageRcmFailbackRecoveryPointTypeCrashConsistent       InMageRcmFailbackRecoveryPointType = "CrashConsistent"
)

func PossibleValuesForInMageRcmFailbackRecoveryPointType() []string {
	return []string{
		string(InMageRcmFailbackRecoveryPointTypeApplicationConsistent),
		string(InMageRcmFailbackRecoveryPointTypeCrashConsistent),
	}
}

type InMageV2RpRecoveryPointType string

const (
	InMageV2RpRecoveryPointTypeLatest                      InMageV2RpRecoveryPointType = "Latest"
	InMageV2RpRecoveryPointTypeLatestApplicationConsistent InMageV2RpRecoveryPointType = "LatestApplicationConsistent"
	InMageV2RpRecoveryPointTypeLatestCrashConsistent       InMageV2RpRecoveryPointType = "LatestCrashConsistent"
	InMageV2RpRecoveryPointTypeLatestProcessed             InMageV2RpRecoveryPointType = "LatestProcessed"
)

func PossibleValuesForInMageV2RpRecoveryPointType() []string {
	return []string{
		string(InMageV2RpRecoveryPointTypeLatest),
		string(InMageV2RpRecoveryPointTypeLatestApplicationConsistent),
		string(InMageV2RpRecoveryPointTypeLatestCrashConsistent),
		string(InMageV2RpRecoveryPointTypeLatestProcessed),
	}
}

type MultiVMSyncPointOption string

const (
	MultiVMSyncPointOptionUseMultiVMSyncRecoveryPoint MultiVMSyncPointOption = "UseMultiVmSyncRecoveryPoint"
	MultiVMSyncPointOptionUsePerVMRecoveryPoint       MultiVMSyncPointOption = "UsePerVmRecoveryPoint"
)

func PossibleValuesForMultiVMSyncPointOption() []string {
	return []string{
		string(MultiVMSyncPointOptionUseMultiVMSyncRecoveryPoint),
		string(MultiVMSyncPointOptionUsePerVMRecoveryPoint),
	}
}

type PossibleOperationsDirections string

const (
	PossibleOperationsDirectionsPrimaryToRecovery PossibleOperationsDirections = "PrimaryToRecovery"
	PossibleOperationsDirectionsRecoveryToPrimary PossibleOperationsDirections = "RecoveryToPrimary"
)

func PossibleValuesForPossibleOperationsDirections() []string {
	return []string{
		string(PossibleOperationsDirectionsPrimaryToRecovery),
		string(PossibleOperationsDirectionsRecoveryToPrimary),
	}
}

type RecoveryPlanActionLocation string

const (
	RecoveryPlanActionLocationPrimary  RecoveryPlanActionLocation = "Primary"
	RecoveryPlanActionLocationRecovery RecoveryPlanActionLocation = "Recovery"
)

func PossibleValuesForRecoveryPlanActionLocation() []string {
	return []string{
		string(RecoveryPlanActionLocationPrimary),
		string(RecoveryPlanActionLocationRecovery),
	}
}

type RecoveryPlanGroupType string

const (
	RecoveryPlanGroupTypeBoot     RecoveryPlanGroupType = "Boot"
	RecoveryPlanGroupTypeFailover RecoveryPlanGroupType = "Failover"
	RecoveryPlanGroupTypeShutdown RecoveryPlanGroupType = "Shutdown"
)

func PossibleValuesForRecoveryPlanGroupType() []string {
	return []string{
		string(RecoveryPlanGroupTypeBoot),
		string(RecoveryPlanGroupTypeFailover),
		string(RecoveryPlanGroupTypeShutdown),
	}
}

type RecoveryPlanPointType string

const (
	RecoveryPlanPointTypeLatest                      RecoveryPlanPointType = "Latest"
	RecoveryPlanPointTypeLatestApplicationConsistent RecoveryPlanPointType = "LatestApplicationConsistent"
	RecoveryPlanPointTypeLatestCrashConsistent       RecoveryPlanPointType = "LatestCrashConsistent"
	RecoveryPlanPointTypeLatestProcessed             RecoveryPlanPointType = "LatestProcessed"
)

func PossibleValuesForRecoveryPlanPointType() []string {
	return []string{
		string(RecoveryPlanPointTypeLatest),
		string(RecoveryPlanPointTypeLatestApplicationConsistent),
		string(RecoveryPlanPointTypeLatestCrashConsistent),
		string(RecoveryPlanPointTypeLatestProcessed),
	}
}

type ReplicationProtectedItemOperation string

const (
	ReplicationProtectedItemOperationCancelFailover      ReplicationProtectedItemOperation = "CancelFailover"
	ReplicationProtectedItemOperationChangePit           ReplicationProtectedItemOperation = "ChangePit"
	ReplicationProtectedItemOperationCommit              ReplicationProtectedItemOperation = "Commit"
	ReplicationProtectedItemOperationCompleteMigration   ReplicationProtectedItemOperation = "CompleteMigration"
	ReplicationProtectedItemOperationDisableProtection   ReplicationProtectedItemOperation = "DisableProtection"
	ReplicationProtectedItemOperationFailback            ReplicationProtectedItemOperation = "Failback"
	ReplicationProtectedItemOperationFinalizeFailback    ReplicationProtectedItemOperation = "FinalizeFailback"
	ReplicationProtectedItemOperationPlannedFailover     ReplicationProtectedItemOperation = "PlannedFailover"
	ReplicationProtectedItemOperationRepairReplication   ReplicationProtectedItemOperation = "RepairReplication"
	ReplicationProtectedItemOperationReverseReplicate    ReplicationProtectedItemOperation = "ReverseReplicate"
	ReplicationProtectedItemOperationSwitchProtection    ReplicationProtectedItemOperation = "SwitchProtection"
	ReplicationProtectedItemOperationTestFailover        ReplicationProtectedItemOperation = "TestFailover"
	ReplicationProtectedItemOperationTestFailoverCleanup ReplicationProtectedItemOperation = "TestFailoverCleanup"
	ReplicationProtectedItemOperationUnplannedFailover   ReplicationProtectedItemOperation = "UnplannedFailover"
)

func PossibleValuesForReplicationProtectedItemOperation() []string {
	return []string{
		string(ReplicationProtectedItemOperationCancelFailover),
		string(ReplicationProtectedItemOperationChangePit),
		string(ReplicationProtectedItemOperationCommit),
		string(ReplicationProtectedItemOperationCompleteMigration),
		string(ReplicationProtectedItemOperationDisableProtection),
		string(ReplicationProtectedItemOperationFailback),
		string(ReplicationProtectedItemOperationFinalizeFailback),
		string(ReplicationProtectedItemOperationPlannedFailover),
		string(ReplicationProtectedItemOperationRepairReplication),
		string(ReplicationProtectedItemOperationReverseReplicate),
		string(ReplicationProtectedItemOperationSwitchProtection),
		string(ReplicationProtectedItemOperationTestFailover),
		string(ReplicationProtectedItemOperationTestFailoverCleanup),
		string(ReplicationProtectedItemOperationUnplannedFailover),
	}
}

type RpInMageRecoveryPointType string

const (
	RpInMageRecoveryPointTypeCustom     RpInMageRecoveryPointType = "Custom"
	RpInMageRecoveryPointTypeLatestTag  RpInMageRecoveryPointType = "LatestTag"
	RpInMageRecoveryPointTypeLatestTime RpInMageRecoveryPointType = "LatestTime"
)

func PossibleValuesForRpInMageRecoveryPointType() []string {
	return []string{
		string(RpInMageRecoveryPointTypeCustom),
		string(RpInMageRecoveryPointTypeLatestTag),
		string(RpInMageRecoveryPointTypeLatestTime),
	}
}

type SourceSiteOperations string

const (
	SourceSiteOperationsNotRequired SourceSiteOperations = "NotRequired"
	SourceSiteOperationsRequired    SourceSiteOperations = "Required"
)

func PossibleValuesForSourceSiteOperations() []string {
	return []string{
		string(SourceSiteOperationsNotRequired),
		string(SourceSiteOperationsRequired),
	}
}
