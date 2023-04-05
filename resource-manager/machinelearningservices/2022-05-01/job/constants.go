package job

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DistributionType string

const (
	DistributionTypeMpi        DistributionType = "Mpi"
	DistributionTypePyTorch    DistributionType = "PyTorch"
	DistributionTypeTensorFlow DistributionType = "TensorFlow"
)

func PossibleValuesForDistributionType() []string {
	return []string{
		string(DistributionTypeMpi),
		string(DistributionTypePyTorch),
		string(DistributionTypeTensorFlow),
	}
}

type EarlyTerminationPolicyType string

const (
	EarlyTerminationPolicyTypeBandit              EarlyTerminationPolicyType = "Bandit"
	EarlyTerminationPolicyTypeMedianStopping      EarlyTerminationPolicyType = "MedianStopping"
	EarlyTerminationPolicyTypeTruncationSelection EarlyTerminationPolicyType = "TruncationSelection"
)

func PossibleValuesForEarlyTerminationPolicyType() []string {
	return []string{
		string(EarlyTerminationPolicyTypeBandit),
		string(EarlyTerminationPolicyTypeMedianStopping),
		string(EarlyTerminationPolicyTypeTruncationSelection),
	}
}

type Goal string

const (
	GoalMaximize Goal = "Maximize"
	GoalMinimize Goal = "Minimize"
)

func PossibleValuesForGoal() []string {
	return []string{
		string(GoalMaximize),
		string(GoalMinimize),
	}
}

type IdentityConfigurationType string

const (
	IdentityConfigurationTypeAMLToken     IdentityConfigurationType = "AMLToken"
	IdentityConfigurationTypeManaged      IdentityConfigurationType = "Managed"
	IdentityConfigurationTypeUserIdentity IdentityConfigurationType = "UserIdentity"
)

func PossibleValuesForIdentityConfigurationType() []string {
	return []string{
		string(IdentityConfigurationTypeAMLToken),
		string(IdentityConfigurationTypeManaged),
		string(IdentityConfigurationTypeUserIdentity),
	}
}

type InputDeliveryMode string

const (
	InputDeliveryModeDirect         InputDeliveryMode = "Direct"
	InputDeliveryModeDownload       InputDeliveryMode = "Download"
	InputDeliveryModeEvalDownload   InputDeliveryMode = "EvalDownload"
	InputDeliveryModeEvalMount      InputDeliveryMode = "EvalMount"
	InputDeliveryModeReadOnlyMount  InputDeliveryMode = "ReadOnlyMount"
	InputDeliveryModeReadWriteMount InputDeliveryMode = "ReadWriteMount"
)

func PossibleValuesForInputDeliveryMode() []string {
	return []string{
		string(InputDeliveryModeDirect),
		string(InputDeliveryModeDownload),
		string(InputDeliveryModeEvalDownload),
		string(InputDeliveryModeEvalMount),
		string(InputDeliveryModeReadOnlyMount),
		string(InputDeliveryModeReadWriteMount),
	}
}

type JobInputType string

const (
	JobInputTypeCustomModel JobInputType = "custom_model"
	JobInputTypeLiteral     JobInputType = "literal"
	JobInputTypeMlflowModel JobInputType = "mlflow_model"
	JobInputTypeMltable     JobInputType = "mltable"
	JobInputTypeTritonModel JobInputType = "triton_model"
	JobInputTypeUriFile     JobInputType = "uri_file"
	JobInputTypeUriFolder   JobInputType = "uri_folder"
)

func PossibleValuesForJobInputType() []string {
	return []string{
		string(JobInputTypeCustomModel),
		string(JobInputTypeLiteral),
		string(JobInputTypeMlflowModel),
		string(JobInputTypeMltable),
		string(JobInputTypeTritonModel),
		string(JobInputTypeUriFile),
		string(JobInputTypeUriFolder),
	}
}

type JobLimitsType string

const (
	JobLimitsTypeCommand JobLimitsType = "Command"
	JobLimitsTypeSweep   JobLimitsType = "Sweep"
)

func PossibleValuesForJobLimitsType() []string {
	return []string{
		string(JobLimitsTypeCommand),
		string(JobLimitsTypeSweep),
	}
}

type JobOutputType string

const (
	JobOutputTypeCustomModel JobOutputType = "custom_model"
	JobOutputTypeMlflowModel JobOutputType = "mlflow_model"
	JobOutputTypeMltable     JobOutputType = "mltable"
	JobOutputTypeTritonModel JobOutputType = "triton_model"
	JobOutputTypeUriFile     JobOutputType = "uri_file"
	JobOutputTypeUriFolder   JobOutputType = "uri_folder"
)

func PossibleValuesForJobOutputType() []string {
	return []string{
		string(JobOutputTypeCustomModel),
		string(JobOutputTypeMlflowModel),
		string(JobOutputTypeMltable),
		string(JobOutputTypeTritonModel),
		string(JobOutputTypeUriFile),
		string(JobOutputTypeUriFolder),
	}
}

type JobStatus string

const (
	JobStatusCancelRequested JobStatus = "CancelRequested"
	JobStatusCanceled        JobStatus = "Canceled"
	JobStatusCompleted       JobStatus = "Completed"
	JobStatusFailed          JobStatus = "Failed"
	JobStatusFinalizing      JobStatus = "Finalizing"
	JobStatusNotResponding   JobStatus = "NotResponding"
	JobStatusNotStarted      JobStatus = "NotStarted"
	JobStatusPaused          JobStatus = "Paused"
	JobStatusPreparing       JobStatus = "Preparing"
	JobStatusProvisioning    JobStatus = "Provisioning"
	JobStatusQueued          JobStatus = "Queued"
	JobStatusRunning         JobStatus = "Running"
	JobStatusStarting        JobStatus = "Starting"
	JobStatusUnknown         JobStatus = "Unknown"
)

func PossibleValuesForJobStatus() []string {
	return []string{
		string(JobStatusCancelRequested),
		string(JobStatusCanceled),
		string(JobStatusCompleted),
		string(JobStatusFailed),
		string(JobStatusFinalizing),
		string(JobStatusNotResponding),
		string(JobStatusNotStarted),
		string(JobStatusPaused),
		string(JobStatusPreparing),
		string(JobStatusProvisioning),
		string(JobStatusQueued),
		string(JobStatusRunning),
		string(JobStatusStarting),
		string(JobStatusUnknown),
	}
}

type JobType string

const (
	JobTypeCommand  JobType = "Command"
	JobTypePipeline JobType = "Pipeline"
	JobTypeSweep    JobType = "Sweep"
)

func PossibleValuesForJobType() []string {
	return []string{
		string(JobTypeCommand),
		string(JobTypePipeline),
		string(JobTypeSweep),
	}
}

type ListViewType string

const (
	ListViewTypeActiveOnly   ListViewType = "ActiveOnly"
	ListViewTypeAll          ListViewType = "All"
	ListViewTypeArchivedOnly ListViewType = "ArchivedOnly"
)

func PossibleValuesForListViewType() []string {
	return []string{
		string(ListViewTypeActiveOnly),
		string(ListViewTypeAll),
		string(ListViewTypeArchivedOnly),
	}
}

type OutputDeliveryMode string

const (
	OutputDeliveryModeReadWriteMount OutputDeliveryMode = "ReadWriteMount"
	OutputDeliveryModeUpload         OutputDeliveryMode = "Upload"
)

func PossibleValuesForOutputDeliveryMode() []string {
	return []string{
		string(OutputDeliveryModeReadWriteMount),
		string(OutputDeliveryModeUpload),
	}
}

type RandomSamplingAlgorithmRule string

const (
	RandomSamplingAlgorithmRuleRandom RandomSamplingAlgorithmRule = "Random"
	RandomSamplingAlgorithmRuleSobol  RandomSamplingAlgorithmRule = "Sobol"
)

func PossibleValuesForRandomSamplingAlgorithmRule() []string {
	return []string{
		string(RandomSamplingAlgorithmRuleRandom),
		string(RandomSamplingAlgorithmRuleSobol),
	}
}

type SamplingAlgorithmType string

const (
	SamplingAlgorithmTypeBayesian SamplingAlgorithmType = "Bayesian"
	SamplingAlgorithmTypeGrid     SamplingAlgorithmType = "Grid"
	SamplingAlgorithmTypeRandom   SamplingAlgorithmType = "Random"
)

func PossibleValuesForSamplingAlgorithmType() []string {
	return []string{
		string(SamplingAlgorithmTypeBayesian),
		string(SamplingAlgorithmTypeGrid),
		string(SamplingAlgorithmTypeRandom),
	}
}
