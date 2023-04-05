package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BlockedTransformers string

const (
	BlockedTransformersCatTargetEncoder  BlockedTransformers = "CatTargetEncoder"
	BlockedTransformersCountVectorizer   BlockedTransformers = "CountVectorizer"
	BlockedTransformersHashOneHotEncoder BlockedTransformers = "HashOneHotEncoder"
	BlockedTransformersLabelEncoder      BlockedTransformers = "LabelEncoder"
	BlockedTransformersNaiveBayes        BlockedTransformers = "NaiveBayes"
	BlockedTransformersOneHotEncoder     BlockedTransformers = "OneHotEncoder"
	BlockedTransformersTextTargetEncoder BlockedTransformers = "TextTargetEncoder"
	BlockedTransformersTfIdf             BlockedTransformers = "TfIdf"
	BlockedTransformersWoETargetEncoder  BlockedTransformers = "WoETargetEncoder"
	BlockedTransformersWordEmbedding     BlockedTransformers = "WordEmbedding"
)

func PossibleValuesForBlockedTransformers() []string {
	return []string{
		string(BlockedTransformersCatTargetEncoder),
		string(BlockedTransformersCountVectorizer),
		string(BlockedTransformersHashOneHotEncoder),
		string(BlockedTransformersLabelEncoder),
		string(BlockedTransformersNaiveBayes),
		string(BlockedTransformersOneHotEncoder),
		string(BlockedTransformersTextTargetEncoder),
		string(BlockedTransformersTfIdf),
		string(BlockedTransformersWoETargetEncoder),
		string(BlockedTransformersWordEmbedding),
	}
}

type ClassificationModels string

const (
	ClassificationModelsBernoulliNaiveBayes   ClassificationModels = "BernoulliNaiveBayes"
	ClassificationModelsDecisionTree          ClassificationModels = "DecisionTree"
	ClassificationModelsExtremeRandomTrees    ClassificationModels = "ExtremeRandomTrees"
	ClassificationModelsGradientBoosting      ClassificationModels = "GradientBoosting"
	ClassificationModelsKNN                   ClassificationModels = "KNN"
	ClassificationModelsLightGBM              ClassificationModels = "LightGBM"
	ClassificationModelsLinearSVM             ClassificationModels = "LinearSVM"
	ClassificationModelsLogisticRegression    ClassificationModels = "LogisticRegression"
	ClassificationModelsMultinomialNaiveBayes ClassificationModels = "MultinomialNaiveBayes"
	ClassificationModelsRandomForest          ClassificationModels = "RandomForest"
	ClassificationModelsSGD                   ClassificationModels = "SGD"
	ClassificationModelsSVM                   ClassificationModels = "SVM"
	ClassificationModelsXGBoostClassifier     ClassificationModels = "XGBoostClassifier"
)

func PossibleValuesForClassificationModels() []string {
	return []string{
		string(ClassificationModelsBernoulliNaiveBayes),
		string(ClassificationModelsDecisionTree),
		string(ClassificationModelsExtremeRandomTrees),
		string(ClassificationModelsGradientBoosting),
		string(ClassificationModelsKNN),
		string(ClassificationModelsLightGBM),
		string(ClassificationModelsLinearSVM),
		string(ClassificationModelsLogisticRegression),
		string(ClassificationModelsMultinomialNaiveBayes),
		string(ClassificationModelsRandomForest),
		string(ClassificationModelsSGD),
		string(ClassificationModelsSVM),
		string(ClassificationModelsXGBoostClassifier),
	}
}

type ClassificationMultilabelPrimaryMetrics string

const (
	ClassificationMultilabelPrimaryMetricsAUCWeighted                   ClassificationMultilabelPrimaryMetrics = "AUCWeighted"
	ClassificationMultilabelPrimaryMetricsAccuracy                      ClassificationMultilabelPrimaryMetrics = "Accuracy"
	ClassificationMultilabelPrimaryMetricsAveragePrecisionScoreWeighted ClassificationMultilabelPrimaryMetrics = "AveragePrecisionScoreWeighted"
	ClassificationMultilabelPrimaryMetricsIOU                           ClassificationMultilabelPrimaryMetrics = "IOU"
	ClassificationMultilabelPrimaryMetricsNormMacroRecall               ClassificationMultilabelPrimaryMetrics = "NormMacroRecall"
	ClassificationMultilabelPrimaryMetricsPrecisionScoreWeighted        ClassificationMultilabelPrimaryMetrics = "PrecisionScoreWeighted"
)

func PossibleValuesForClassificationMultilabelPrimaryMetrics() []string {
	return []string{
		string(ClassificationMultilabelPrimaryMetricsAUCWeighted),
		string(ClassificationMultilabelPrimaryMetricsAccuracy),
		string(ClassificationMultilabelPrimaryMetricsAveragePrecisionScoreWeighted),
		string(ClassificationMultilabelPrimaryMetricsIOU),
		string(ClassificationMultilabelPrimaryMetricsNormMacroRecall),
		string(ClassificationMultilabelPrimaryMetricsPrecisionScoreWeighted),
	}
}

type ClassificationPrimaryMetrics string

const (
	ClassificationPrimaryMetricsAUCWeighted                   ClassificationPrimaryMetrics = "AUCWeighted"
	ClassificationPrimaryMetricsAccuracy                      ClassificationPrimaryMetrics = "Accuracy"
	ClassificationPrimaryMetricsAveragePrecisionScoreWeighted ClassificationPrimaryMetrics = "AveragePrecisionScoreWeighted"
	ClassificationPrimaryMetricsNormMacroRecall               ClassificationPrimaryMetrics = "NormMacroRecall"
	ClassificationPrimaryMetricsPrecisionScoreWeighted        ClassificationPrimaryMetrics = "PrecisionScoreWeighted"
)

func PossibleValuesForClassificationPrimaryMetrics() []string {
	return []string{
		string(ClassificationPrimaryMetricsAUCWeighted),
		string(ClassificationPrimaryMetricsAccuracy),
		string(ClassificationPrimaryMetricsAveragePrecisionScoreWeighted),
		string(ClassificationPrimaryMetricsNormMacroRecall),
		string(ClassificationPrimaryMetricsPrecisionScoreWeighted),
	}
}

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

type FeatureLags string

const (
	FeatureLagsAuto FeatureLags = "Auto"
	FeatureLagsNone FeatureLags = "None"
)

func PossibleValuesForFeatureLags() []string {
	return []string{
		string(FeatureLagsAuto),
		string(FeatureLagsNone),
	}
}

type FeaturizationMode string

const (
	FeaturizationModeAuto   FeaturizationMode = "Auto"
	FeaturizationModeCustom FeaturizationMode = "Custom"
	FeaturizationModeOff    FeaturizationMode = "Off"
)

func PossibleValuesForFeaturizationMode() []string {
	return []string{
		string(FeaturizationModeAuto),
		string(FeaturizationModeCustom),
		string(FeaturizationModeOff),
	}
}

type ForecastHorizonMode string

const (
	ForecastHorizonModeAuto   ForecastHorizonMode = "Auto"
	ForecastHorizonModeCustom ForecastHorizonMode = "Custom"
)

func PossibleValuesForForecastHorizonMode() []string {
	return []string{
		string(ForecastHorizonModeAuto),
		string(ForecastHorizonModeCustom),
	}
}

type ForecastingModels string

const (
	ForecastingModelsArimax               ForecastingModels = "Arimax"
	ForecastingModelsAutoArima            ForecastingModels = "AutoArima"
	ForecastingModelsAverage              ForecastingModels = "Average"
	ForecastingModelsDecisionTree         ForecastingModels = "DecisionTree"
	ForecastingModelsElasticNet           ForecastingModels = "ElasticNet"
	ForecastingModelsExponentialSmoothing ForecastingModels = "ExponentialSmoothing"
	ForecastingModelsExtremeRandomTrees   ForecastingModels = "ExtremeRandomTrees"
	ForecastingModelsGradientBoosting     ForecastingModels = "GradientBoosting"
	ForecastingModelsKNN                  ForecastingModels = "KNN"
	ForecastingModelsLassoLars            ForecastingModels = "LassoLars"
	ForecastingModelsLightGBM             ForecastingModels = "LightGBM"
	ForecastingModelsNaive                ForecastingModels = "Naive"
	ForecastingModelsProphet              ForecastingModels = "Prophet"
	ForecastingModelsRandomForest         ForecastingModels = "RandomForest"
	ForecastingModelsSGD                  ForecastingModels = "SGD"
	ForecastingModelsSeasonalAverage      ForecastingModels = "SeasonalAverage"
	ForecastingModelsSeasonalNaive        ForecastingModels = "SeasonalNaive"
	ForecastingModelsTCNForecaster        ForecastingModels = "TCNForecaster"
	ForecastingModelsXGBoostRegressor     ForecastingModels = "XGBoostRegressor"
)

func PossibleValuesForForecastingModels() []string {
	return []string{
		string(ForecastingModelsArimax),
		string(ForecastingModelsAutoArima),
		string(ForecastingModelsAverage),
		string(ForecastingModelsDecisionTree),
		string(ForecastingModelsElasticNet),
		string(ForecastingModelsExponentialSmoothing),
		string(ForecastingModelsExtremeRandomTrees),
		string(ForecastingModelsGradientBoosting),
		string(ForecastingModelsKNN),
		string(ForecastingModelsLassoLars),
		string(ForecastingModelsLightGBM),
		string(ForecastingModelsNaive),
		string(ForecastingModelsProphet),
		string(ForecastingModelsRandomForest),
		string(ForecastingModelsSGD),
		string(ForecastingModelsSeasonalAverage),
		string(ForecastingModelsSeasonalNaive),
		string(ForecastingModelsTCNForecaster),
		string(ForecastingModelsXGBoostRegressor),
	}
}

type ForecastingPrimaryMetrics string

const (
	ForecastingPrimaryMetricsNormalizedMeanAbsoluteError    ForecastingPrimaryMetrics = "NormalizedMeanAbsoluteError"
	ForecastingPrimaryMetricsNormalizedRootMeanSquaredError ForecastingPrimaryMetrics = "NormalizedRootMeanSquaredError"
	ForecastingPrimaryMetricsRTwoScore                      ForecastingPrimaryMetrics = "R2Score"
	ForecastingPrimaryMetricsSpearmanCorrelation            ForecastingPrimaryMetrics = "SpearmanCorrelation"
)

func PossibleValuesForForecastingPrimaryMetrics() []string {
	return []string{
		string(ForecastingPrimaryMetricsNormalizedMeanAbsoluteError),
		string(ForecastingPrimaryMetricsNormalizedRootMeanSquaredError),
		string(ForecastingPrimaryMetricsRTwoScore),
		string(ForecastingPrimaryMetricsSpearmanCorrelation),
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

type InstanceSegmentationPrimaryMetrics string

const (
	InstanceSegmentationPrimaryMetricsMeanAveragePrecision InstanceSegmentationPrimaryMetrics = "MeanAveragePrecision"
)

func PossibleValuesForInstanceSegmentationPrimaryMetrics() []string {
	return []string{
		string(InstanceSegmentationPrimaryMetricsMeanAveragePrecision),
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
	JobTypeAutoML   JobType = "AutoML"
	JobTypeCommand  JobType = "Command"
	JobTypePipeline JobType = "Pipeline"
	JobTypeSweep    JobType = "Sweep"
)

func PossibleValuesForJobType() []string {
	return []string{
		string(JobTypeAutoML),
		string(JobTypeCommand),
		string(JobTypePipeline),
		string(JobTypeSweep),
	}
}

type LearningRateScheduler string

const (
	LearningRateSchedulerNone         LearningRateScheduler = "None"
	LearningRateSchedulerStep         LearningRateScheduler = "Step"
	LearningRateSchedulerWarmupCosine LearningRateScheduler = "WarmupCosine"
)

func PossibleValuesForLearningRateScheduler() []string {
	return []string{
		string(LearningRateSchedulerNone),
		string(LearningRateSchedulerStep),
		string(LearningRateSchedulerWarmupCosine),
	}
}

type LogVerbosity string

const (
	LogVerbosityCritical LogVerbosity = "Critical"
	LogVerbosityDebug    LogVerbosity = "Debug"
	LogVerbosityError    LogVerbosity = "Error"
	LogVerbosityInfo     LogVerbosity = "Info"
	LogVerbosityNotSet   LogVerbosity = "NotSet"
	LogVerbosityWarning  LogVerbosity = "Warning"
)

func PossibleValuesForLogVerbosity() []string {
	return []string{
		string(LogVerbosityCritical),
		string(LogVerbosityDebug),
		string(LogVerbosityError),
		string(LogVerbosityInfo),
		string(LogVerbosityNotSet),
		string(LogVerbosityWarning),
	}
}

type ModelSize string

const (
	ModelSizeExtraLarge ModelSize = "ExtraLarge"
	ModelSizeLarge      ModelSize = "Large"
	ModelSizeMedium     ModelSize = "Medium"
	ModelSizeNone       ModelSize = "None"
	ModelSizeSmall      ModelSize = "Small"
)

func PossibleValuesForModelSize() []string {
	return []string{
		string(ModelSizeExtraLarge),
		string(ModelSizeLarge),
		string(ModelSizeMedium),
		string(ModelSizeNone),
		string(ModelSizeSmall),
	}
}

type NCrossValidationsMode string

const (
	NCrossValidationsModeAuto   NCrossValidationsMode = "Auto"
	NCrossValidationsModeCustom NCrossValidationsMode = "Custom"
)

func PossibleValuesForNCrossValidationsMode() []string {
	return []string{
		string(NCrossValidationsModeAuto),
		string(NCrossValidationsModeCustom),
	}
}

type ObjectDetectionPrimaryMetrics string

const (
	ObjectDetectionPrimaryMetricsMeanAveragePrecision ObjectDetectionPrimaryMetrics = "MeanAveragePrecision"
)

func PossibleValuesForObjectDetectionPrimaryMetrics() []string {
	return []string{
		string(ObjectDetectionPrimaryMetricsMeanAveragePrecision),
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

type RecurrenceFrequency string

const (
	RecurrenceFrequencyDay    RecurrenceFrequency = "Day"
	RecurrenceFrequencyHour   RecurrenceFrequency = "Hour"
	RecurrenceFrequencyMinute RecurrenceFrequency = "Minute"
	RecurrenceFrequencyMonth  RecurrenceFrequency = "Month"
	RecurrenceFrequencyWeek   RecurrenceFrequency = "Week"
)

func PossibleValuesForRecurrenceFrequency() []string {
	return []string{
		string(RecurrenceFrequencyDay),
		string(RecurrenceFrequencyHour),
		string(RecurrenceFrequencyMinute),
		string(RecurrenceFrequencyMonth),
		string(RecurrenceFrequencyWeek),
	}
}

type RegressionModels string

const (
	RegressionModelsDecisionTree       RegressionModels = "DecisionTree"
	RegressionModelsElasticNet         RegressionModels = "ElasticNet"
	RegressionModelsExtremeRandomTrees RegressionModels = "ExtremeRandomTrees"
	RegressionModelsGradientBoosting   RegressionModels = "GradientBoosting"
	RegressionModelsKNN                RegressionModels = "KNN"
	RegressionModelsLassoLars          RegressionModels = "LassoLars"
	RegressionModelsLightGBM           RegressionModels = "LightGBM"
	RegressionModelsRandomForest       RegressionModels = "RandomForest"
	RegressionModelsSGD                RegressionModels = "SGD"
	RegressionModelsXGBoostRegressor   RegressionModels = "XGBoostRegressor"
)

func PossibleValuesForRegressionModels() []string {
	return []string{
		string(RegressionModelsDecisionTree),
		string(RegressionModelsElasticNet),
		string(RegressionModelsExtremeRandomTrees),
		string(RegressionModelsGradientBoosting),
		string(RegressionModelsKNN),
		string(RegressionModelsLassoLars),
		string(RegressionModelsLightGBM),
		string(RegressionModelsRandomForest),
		string(RegressionModelsSGD),
		string(RegressionModelsXGBoostRegressor),
	}
}

type RegressionPrimaryMetrics string

const (
	RegressionPrimaryMetricsNormalizedMeanAbsoluteError    RegressionPrimaryMetrics = "NormalizedMeanAbsoluteError"
	RegressionPrimaryMetricsNormalizedRootMeanSquaredError RegressionPrimaryMetrics = "NormalizedRootMeanSquaredError"
	RegressionPrimaryMetricsRTwoScore                      RegressionPrimaryMetrics = "R2Score"
	RegressionPrimaryMetricsSpearmanCorrelation            RegressionPrimaryMetrics = "SpearmanCorrelation"
)

func PossibleValuesForRegressionPrimaryMetrics() []string {
	return []string{
		string(RegressionPrimaryMetricsNormalizedMeanAbsoluteError),
		string(RegressionPrimaryMetricsNormalizedRootMeanSquaredError),
		string(RegressionPrimaryMetricsRTwoScore),
		string(RegressionPrimaryMetricsSpearmanCorrelation),
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

type ScheduleActionType string

const (
	ScheduleActionTypeCreateJob           ScheduleActionType = "CreateJob"
	ScheduleActionTypeInvokeBatchEndpoint ScheduleActionType = "InvokeBatchEndpoint"
)

func PossibleValuesForScheduleActionType() []string {
	return []string{
		string(ScheduleActionTypeCreateJob),
		string(ScheduleActionTypeInvokeBatchEndpoint),
	}
}

type ScheduleListViewType string

const (
	ScheduleListViewTypeAll          ScheduleListViewType = "All"
	ScheduleListViewTypeDisabledOnly ScheduleListViewType = "DisabledOnly"
	ScheduleListViewTypeEnabledOnly  ScheduleListViewType = "EnabledOnly"
)

func PossibleValuesForScheduleListViewType() []string {
	return []string{
		string(ScheduleListViewTypeAll),
		string(ScheduleListViewTypeDisabledOnly),
		string(ScheduleListViewTypeEnabledOnly),
	}
}

type ScheduleProvisioningStatus string

const (
	ScheduleProvisioningStatusCanceled  ScheduleProvisioningStatus = "Canceled"
	ScheduleProvisioningStatusCreating  ScheduleProvisioningStatus = "Creating"
	ScheduleProvisioningStatusDeleting  ScheduleProvisioningStatus = "Deleting"
	ScheduleProvisioningStatusFailed    ScheduleProvisioningStatus = "Failed"
	ScheduleProvisioningStatusSucceeded ScheduleProvisioningStatus = "Succeeded"
	ScheduleProvisioningStatusUpdating  ScheduleProvisioningStatus = "Updating"
)

func PossibleValuesForScheduleProvisioningStatus() []string {
	return []string{
		string(ScheduleProvisioningStatusCanceled),
		string(ScheduleProvisioningStatusCreating),
		string(ScheduleProvisioningStatusDeleting),
		string(ScheduleProvisioningStatusFailed),
		string(ScheduleProvisioningStatusSucceeded),
		string(ScheduleProvisioningStatusUpdating),
	}
}

type SeasonalityMode string

const (
	SeasonalityModeAuto   SeasonalityMode = "Auto"
	SeasonalityModeCustom SeasonalityMode = "Custom"
)

func PossibleValuesForSeasonalityMode() []string {
	return []string{
		string(SeasonalityModeAuto),
		string(SeasonalityModeCustom),
	}
}

type ShortSeriesHandlingConfiguration string

const (
	ShortSeriesHandlingConfigurationAuto ShortSeriesHandlingConfiguration = "Auto"
	ShortSeriesHandlingConfigurationDrop ShortSeriesHandlingConfiguration = "Drop"
	ShortSeriesHandlingConfigurationNone ShortSeriesHandlingConfiguration = "None"
	ShortSeriesHandlingConfigurationPad  ShortSeriesHandlingConfiguration = "Pad"
)

func PossibleValuesForShortSeriesHandlingConfiguration() []string {
	return []string{
		string(ShortSeriesHandlingConfigurationAuto),
		string(ShortSeriesHandlingConfigurationDrop),
		string(ShortSeriesHandlingConfigurationNone),
		string(ShortSeriesHandlingConfigurationPad),
	}
}

type StackMetaLearnerType string

const (
	StackMetaLearnerTypeElasticNet           StackMetaLearnerType = "ElasticNet"
	StackMetaLearnerTypeElasticNetCV         StackMetaLearnerType = "ElasticNetCV"
	StackMetaLearnerTypeLightGBMClassifier   StackMetaLearnerType = "LightGBMClassifier"
	StackMetaLearnerTypeLightGBMRegressor    StackMetaLearnerType = "LightGBMRegressor"
	StackMetaLearnerTypeLinearRegression     StackMetaLearnerType = "LinearRegression"
	StackMetaLearnerTypeLogisticRegression   StackMetaLearnerType = "LogisticRegression"
	StackMetaLearnerTypeLogisticRegressionCV StackMetaLearnerType = "LogisticRegressionCV"
	StackMetaLearnerTypeNone                 StackMetaLearnerType = "None"
)

func PossibleValuesForStackMetaLearnerType() []string {
	return []string{
		string(StackMetaLearnerTypeElasticNet),
		string(StackMetaLearnerTypeElasticNetCV),
		string(StackMetaLearnerTypeLightGBMClassifier),
		string(StackMetaLearnerTypeLightGBMRegressor),
		string(StackMetaLearnerTypeLinearRegression),
		string(StackMetaLearnerTypeLogisticRegression),
		string(StackMetaLearnerTypeLogisticRegressionCV),
		string(StackMetaLearnerTypeNone),
	}
}

type StochasticOptimizer string

const (
	StochasticOptimizerAdam  StochasticOptimizer = "Adam"
	StochasticOptimizerAdamw StochasticOptimizer = "Adamw"
	StochasticOptimizerNone  StochasticOptimizer = "None"
	StochasticOptimizerSgd   StochasticOptimizer = "Sgd"
)

func PossibleValuesForStochasticOptimizer() []string {
	return []string{
		string(StochasticOptimizerAdam),
		string(StochasticOptimizerAdamw),
		string(StochasticOptimizerNone),
		string(StochasticOptimizerSgd),
	}
}

type TargetAggregationFunction string

const (
	TargetAggregationFunctionMax  TargetAggregationFunction = "Max"
	TargetAggregationFunctionMean TargetAggregationFunction = "Mean"
	TargetAggregationFunctionMin  TargetAggregationFunction = "Min"
	TargetAggregationFunctionNone TargetAggregationFunction = "None"
	TargetAggregationFunctionSum  TargetAggregationFunction = "Sum"
)

func PossibleValuesForTargetAggregationFunction() []string {
	return []string{
		string(TargetAggregationFunctionMax),
		string(TargetAggregationFunctionMean),
		string(TargetAggregationFunctionMin),
		string(TargetAggregationFunctionNone),
		string(TargetAggregationFunctionSum),
	}
}

type TargetLagsMode string

const (
	TargetLagsModeAuto   TargetLagsMode = "Auto"
	TargetLagsModeCustom TargetLagsMode = "Custom"
)

func PossibleValuesForTargetLagsMode() []string {
	return []string{
		string(TargetLagsModeAuto),
		string(TargetLagsModeCustom),
	}
}

type TargetRollingWindowSizeMode string

const (
	TargetRollingWindowSizeModeAuto   TargetRollingWindowSizeMode = "Auto"
	TargetRollingWindowSizeModeCustom TargetRollingWindowSizeMode = "Custom"
)

func PossibleValuesForTargetRollingWindowSizeMode() []string {
	return []string{
		string(TargetRollingWindowSizeModeAuto),
		string(TargetRollingWindowSizeModeCustom),
	}
}

type TaskType string

const (
	TaskTypeClassification                TaskType = "Classification"
	TaskTypeForecasting                   TaskType = "Forecasting"
	TaskTypeImageClassification           TaskType = "ImageClassification"
	TaskTypeImageClassificationMultilabel TaskType = "ImageClassificationMultilabel"
	TaskTypeImageInstanceSegmentation     TaskType = "ImageInstanceSegmentation"
	TaskTypeImageObjectDetection          TaskType = "ImageObjectDetection"
	TaskTypeRegression                    TaskType = "Regression"
	TaskTypeTextClassification            TaskType = "TextClassification"
	TaskTypeTextClassificationMultilabel  TaskType = "TextClassificationMultilabel"
	TaskTypeTextNER                       TaskType = "TextNER"
)

func PossibleValuesForTaskType() []string {
	return []string{
		string(TaskTypeClassification),
		string(TaskTypeForecasting),
		string(TaskTypeImageClassification),
		string(TaskTypeImageClassificationMultilabel),
		string(TaskTypeImageInstanceSegmentation),
		string(TaskTypeImageObjectDetection),
		string(TaskTypeRegression),
		string(TaskTypeTextClassification),
		string(TaskTypeTextClassificationMultilabel),
		string(TaskTypeTextNER),
	}
}

type TriggerType string

const (
	TriggerTypeCron       TriggerType = "Cron"
	TriggerTypeRecurrence TriggerType = "Recurrence"
)

func PossibleValuesForTriggerType() []string {
	return []string{
		string(TriggerTypeCron),
		string(TriggerTypeRecurrence),
	}
}

type UseStl string

const (
	UseStlNone        UseStl = "None"
	UseStlSeason      UseStl = "Season"
	UseStlSeasonTrend UseStl = "SeasonTrend"
)

func PossibleValuesForUseStl() []string {
	return []string{
		string(UseStlNone),
		string(UseStlSeason),
		string(UseStlSeasonTrend),
	}
}

type ValidationMetricType string

const (
	ValidationMetricTypeCoco    ValidationMetricType = "Coco"
	ValidationMetricTypeCocoVoc ValidationMetricType = "CocoVoc"
	ValidationMetricTypeNone    ValidationMetricType = "None"
	ValidationMetricTypeVoc     ValidationMetricType = "Voc"
)

func PossibleValuesForValidationMetricType() []string {
	return []string{
		string(ValidationMetricTypeCoco),
		string(ValidationMetricTypeCocoVoc),
		string(ValidationMetricTypeNone),
		string(ValidationMetricTypeVoc),
	}
}

type WeekDay string

const (
	WeekDayFriday    WeekDay = "Friday"
	WeekDayMonday    WeekDay = "Monday"
	WeekDaySaturday  WeekDay = "Saturday"
	WeekDaySunday    WeekDay = "Sunday"
	WeekDayThursday  WeekDay = "Thursday"
	WeekDayTuesday   WeekDay = "Tuesday"
	WeekDayWednesday WeekDay = "Wednesday"
)

func PossibleValuesForWeekDay() []string {
	return []string{
		string(WeekDayFriday),
		string(WeekDayMonday),
		string(WeekDaySaturday),
		string(WeekDaySunday),
		string(WeekDayThursday),
		string(WeekDayTuesday),
		string(WeekDayWednesday),
	}
}
