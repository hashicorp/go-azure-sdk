package job

import "strings"

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

func parseDistributionType(input string) (*DistributionType, error) {
	vals := map[string]DistributionType{
		"mpi":        DistributionTypeMpi,
		"pytorch":    DistributionTypePyTorch,
		"tensorflow": DistributionTypeTensorFlow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DistributionType(input)
	return &out, nil
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

func parseEarlyTerminationPolicyType(input string) (*EarlyTerminationPolicyType, error) {
	vals := map[string]EarlyTerminationPolicyType{
		"bandit":              EarlyTerminationPolicyTypeBandit,
		"medianstopping":      EarlyTerminationPolicyTypeMedianStopping,
		"truncationselection": EarlyTerminationPolicyTypeTruncationSelection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EarlyTerminationPolicyType(input)
	return &out, nil
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

func parseGoal(input string) (*Goal, error) {
	vals := map[string]Goal{
		"maximize": GoalMaximize,
		"minimize": GoalMinimize,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Goal(input)
	return &out, nil
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

func parseIdentityConfigurationType(input string) (*IdentityConfigurationType, error) {
	vals := map[string]IdentityConfigurationType{
		"amltoken":     IdentityConfigurationTypeAMLToken,
		"managed":      IdentityConfigurationTypeManaged,
		"useridentity": IdentityConfigurationTypeUserIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityConfigurationType(input)
	return &out, nil
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

func parseInputDeliveryMode(input string) (*InputDeliveryMode, error) {
	vals := map[string]InputDeliveryMode{
		"direct":         InputDeliveryModeDirect,
		"download":       InputDeliveryModeDownload,
		"evaldownload":   InputDeliveryModeEvalDownload,
		"evalmount":      InputDeliveryModeEvalMount,
		"readonlymount":  InputDeliveryModeReadOnlyMount,
		"readwritemount": InputDeliveryModeReadWriteMount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InputDeliveryMode(input)
	return &out, nil
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

func parseJobInputType(input string) (*JobInputType, error) {
	vals := map[string]JobInputType{
		"custom_model": JobInputTypeCustomModel,
		"literal":      JobInputTypeLiteral,
		"mlflow_model": JobInputTypeMlflowModel,
		"mltable":      JobInputTypeMltable,
		"triton_model": JobInputTypeTritonModel,
		"uri_file":     JobInputTypeUriFile,
		"uri_folder":   JobInputTypeUriFolder,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobInputType(input)
	return &out, nil
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

func parseJobLimitsType(input string) (*JobLimitsType, error) {
	vals := map[string]JobLimitsType{
		"command": JobLimitsTypeCommand,
		"sweep":   JobLimitsTypeSweep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobLimitsType(input)
	return &out, nil
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

func parseJobOutputType(input string) (*JobOutputType, error) {
	vals := map[string]JobOutputType{
		"custom_model": JobOutputTypeCustomModel,
		"mlflow_model": JobOutputTypeMlflowModel,
		"mltable":      JobOutputTypeMltable,
		"triton_model": JobOutputTypeTritonModel,
		"uri_file":     JobOutputTypeUriFile,
		"uri_folder":   JobOutputTypeUriFolder,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobOutputType(input)
	return &out, nil
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

func parseJobStatus(input string) (*JobStatus, error) {
	vals := map[string]JobStatus{
		"cancelrequested": JobStatusCancelRequested,
		"canceled":        JobStatusCanceled,
		"completed":       JobStatusCompleted,
		"failed":          JobStatusFailed,
		"finalizing":      JobStatusFinalizing,
		"notresponding":   JobStatusNotResponding,
		"notstarted":      JobStatusNotStarted,
		"paused":          JobStatusPaused,
		"preparing":       JobStatusPreparing,
		"provisioning":    JobStatusProvisioning,
		"queued":          JobStatusQueued,
		"running":         JobStatusRunning,
		"starting":        JobStatusStarting,
		"unknown":         JobStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobStatus(input)
	return &out, nil
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

func parseJobType(input string) (*JobType, error) {
	vals := map[string]JobType{
		"command":  JobTypeCommand,
		"pipeline": JobTypePipeline,
		"sweep":    JobTypeSweep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobType(input)
	return &out, nil
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

func parseListViewType(input string) (*ListViewType, error) {
	vals := map[string]ListViewType{
		"activeonly":   ListViewTypeActiveOnly,
		"all":          ListViewTypeAll,
		"archivedonly": ListViewTypeArchivedOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ListViewType(input)
	return &out, nil
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

func parseOutputDeliveryMode(input string) (*OutputDeliveryMode, error) {
	vals := map[string]OutputDeliveryMode{
		"readwritemount": OutputDeliveryModeReadWriteMount,
		"upload":         OutputDeliveryModeUpload,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutputDeliveryMode(input)
	return &out, nil
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

func parseRandomSamplingAlgorithmRule(input string) (*RandomSamplingAlgorithmRule, error) {
	vals := map[string]RandomSamplingAlgorithmRule{
		"random": RandomSamplingAlgorithmRuleRandom,
		"sobol":  RandomSamplingAlgorithmRuleSobol,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RandomSamplingAlgorithmRule(input)
	return &out, nil
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

func parseSamplingAlgorithmType(input string) (*SamplingAlgorithmType, error) {
	vals := map[string]SamplingAlgorithmType{
		"bayesian": SamplingAlgorithmTypeBayesian,
		"grid":     SamplingAlgorithmTypeGrid,
		"random":   SamplingAlgorithmTypeRandom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SamplingAlgorithmType(input)
	return &out, nil
}
