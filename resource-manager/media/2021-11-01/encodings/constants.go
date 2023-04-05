package encodings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AacAudioProfile string

const (
	AacAudioProfileAacLc     AacAudioProfile = "AacLc"
	AacAudioProfileHeAacVOne AacAudioProfile = "HeAacV1"
	AacAudioProfileHeAacVTwo AacAudioProfile = "HeAacV2"
)

func PossibleValuesForAacAudioProfile() []string {
	return []string{
		string(AacAudioProfileAacLc),
		string(AacAudioProfileHeAacVOne),
		string(AacAudioProfileHeAacVTwo),
	}
}

func (s *AacAudioProfile) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAacAudioProfile() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AacAudioProfile(decoded)
	return nil
}

type AnalysisResolution string

const (
	AnalysisResolutionSourceResolution   AnalysisResolution = "SourceResolution"
	AnalysisResolutionStandardDefinition AnalysisResolution = "StandardDefinition"
)

func PossibleValuesForAnalysisResolution() []string {
	return []string{
		string(AnalysisResolutionSourceResolution),
		string(AnalysisResolutionStandardDefinition),
	}
}

func (s *AnalysisResolution) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAnalysisResolution() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AnalysisResolution(decoded)
	return nil
}

type AttributeFilter string

const (
	AttributeFilterAll         AttributeFilter = "All"
	AttributeFilterBottom      AttributeFilter = "Bottom"
	AttributeFilterTop         AttributeFilter = "Top"
	AttributeFilterValueEquals AttributeFilter = "ValueEquals"
)

func PossibleValuesForAttributeFilter() []string {
	return []string{
		string(AttributeFilterAll),
		string(AttributeFilterBottom),
		string(AttributeFilterTop),
		string(AttributeFilterValueEquals),
	}
}

func (s *AttributeFilter) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAttributeFilter() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AttributeFilter(decoded)
	return nil
}

type AudioAnalysisMode string

const (
	AudioAnalysisModeBasic    AudioAnalysisMode = "Basic"
	AudioAnalysisModeStandard AudioAnalysisMode = "Standard"
)

func PossibleValuesForAudioAnalysisMode() []string {
	return []string{
		string(AudioAnalysisModeBasic),
		string(AudioAnalysisModeStandard),
	}
}

func (s *AudioAnalysisMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAudioAnalysisMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AudioAnalysisMode(decoded)
	return nil
}

type BlurType string

const (
	BlurTypeBlack BlurType = "Black"
	BlurTypeBox   BlurType = "Box"
	BlurTypeHigh  BlurType = "High"
	BlurTypeLow   BlurType = "Low"
	BlurTypeMed   BlurType = "Med"
)

func PossibleValuesForBlurType() []string {
	return []string{
		string(BlurTypeBlack),
		string(BlurTypeBox),
		string(BlurTypeHigh),
		string(BlurTypeLow),
		string(BlurTypeMed),
	}
}

func (s *BlurType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForBlurType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = BlurType(decoded)
	return nil
}

type ChannelMapping string

const (
	ChannelMappingBackLeft            ChannelMapping = "BackLeft"
	ChannelMappingBackRight           ChannelMapping = "BackRight"
	ChannelMappingCenter              ChannelMapping = "Center"
	ChannelMappingFrontLeft           ChannelMapping = "FrontLeft"
	ChannelMappingFrontRight          ChannelMapping = "FrontRight"
	ChannelMappingLowFrequencyEffects ChannelMapping = "LowFrequencyEffects"
	ChannelMappingStereoLeft          ChannelMapping = "StereoLeft"
	ChannelMappingStereoRight         ChannelMapping = "StereoRight"
)

func PossibleValuesForChannelMapping() []string {
	return []string{
		string(ChannelMappingBackLeft),
		string(ChannelMappingBackRight),
		string(ChannelMappingCenter),
		string(ChannelMappingFrontLeft),
		string(ChannelMappingFrontRight),
		string(ChannelMappingLowFrequencyEffects),
		string(ChannelMappingStereoLeft),
		string(ChannelMappingStereoRight),
	}
}

func (s *ChannelMapping) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForChannelMapping() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ChannelMapping(decoded)
	return nil
}

type Complexity string

const (
	ComplexityBalanced Complexity = "Balanced"
	ComplexityQuality  Complexity = "Quality"
	ComplexitySpeed    Complexity = "Speed"
)

func PossibleValuesForComplexity() []string {
	return []string{
		string(ComplexityBalanced),
		string(ComplexityQuality),
		string(ComplexitySpeed),
	}
}

func (s *Complexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForComplexity() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = Complexity(decoded)
	return nil
}

type DeinterlaceMode string

const (
	DeinterlaceModeAutoPixelAdaptive DeinterlaceMode = "AutoPixelAdaptive"
	DeinterlaceModeOff               DeinterlaceMode = "Off"
)

func PossibleValuesForDeinterlaceMode() []string {
	return []string{
		string(DeinterlaceModeAutoPixelAdaptive),
		string(DeinterlaceModeOff),
	}
}

func (s *DeinterlaceMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForDeinterlaceMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = DeinterlaceMode(decoded)
	return nil
}

type DeinterlaceParity string

const (
	DeinterlaceParityAuto             DeinterlaceParity = "Auto"
	DeinterlaceParityBottomFieldFirst DeinterlaceParity = "BottomFieldFirst"
	DeinterlaceParityTopFieldFirst    DeinterlaceParity = "TopFieldFirst"
)

func PossibleValuesForDeinterlaceParity() []string {
	return []string{
		string(DeinterlaceParityAuto),
		string(DeinterlaceParityBottomFieldFirst),
		string(DeinterlaceParityTopFieldFirst),
	}
}

func (s *DeinterlaceParity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForDeinterlaceParity() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = DeinterlaceParity(decoded)
	return nil
}

type EncoderNamedPreset string

const (
	EncoderNamedPresetAACGoodQualityAudio                         EncoderNamedPreset = "AACGoodQualityAudio"
	EncoderNamedPresetAdaptiveStreaming                           EncoderNamedPreset = "AdaptiveStreaming"
	EncoderNamedPresetContentAwareEncoding                        EncoderNamedPreset = "ContentAwareEncoding"
	EncoderNamedPresetContentAwareEncodingExperimental            EncoderNamedPreset = "ContentAwareEncodingExperimental"
	EncoderNamedPresetCopyAllBitrateNonInterleaved                EncoderNamedPreset = "CopyAllBitrateNonInterleaved"
	EncoderNamedPresetHTwoSixFiveAdaptiveStreaming                EncoderNamedPreset = "H265AdaptiveStreaming"
	EncoderNamedPresetHTwoSixFiveContentAwareEncoding             EncoderNamedPreset = "H265ContentAwareEncoding"
	EncoderNamedPresetHTwoSixFiveSingleBitrateFourK               EncoderNamedPreset = "H265SingleBitrate4K"
	EncoderNamedPresetHTwoSixFiveSingleBitrateOneZeroEightZerop   EncoderNamedPreset = "H265SingleBitrate1080p"
	EncoderNamedPresetHTwoSixFiveSingleBitrateSevenTwoZerop       EncoderNamedPreset = "H265SingleBitrate720p"
	EncoderNamedPresetHTwoSixFourMultipleBitrateOneZeroEightZerop EncoderNamedPreset = "H264MultipleBitrate1080p"
	EncoderNamedPresetHTwoSixFourMultipleBitrateSD                EncoderNamedPreset = "H264MultipleBitrateSD"
	EncoderNamedPresetHTwoSixFourMultipleBitrateSevenTwoZerop     EncoderNamedPreset = "H264MultipleBitrate720p"
	EncoderNamedPresetHTwoSixFourSingleBitrateOneZeroEightZerop   EncoderNamedPreset = "H264SingleBitrate1080p"
	EncoderNamedPresetHTwoSixFourSingleBitrateSD                  EncoderNamedPreset = "H264SingleBitrateSD"
	EncoderNamedPresetHTwoSixFourSingleBitrateSevenTwoZerop       EncoderNamedPreset = "H264SingleBitrate720p"
)

func PossibleValuesForEncoderNamedPreset() []string {
	return []string{
		string(EncoderNamedPresetAACGoodQualityAudio),
		string(EncoderNamedPresetAdaptiveStreaming),
		string(EncoderNamedPresetContentAwareEncoding),
		string(EncoderNamedPresetContentAwareEncodingExperimental),
		string(EncoderNamedPresetCopyAllBitrateNonInterleaved),
		string(EncoderNamedPresetHTwoSixFiveAdaptiveStreaming),
		string(EncoderNamedPresetHTwoSixFiveContentAwareEncoding),
		string(EncoderNamedPresetHTwoSixFiveSingleBitrateFourK),
		string(EncoderNamedPresetHTwoSixFiveSingleBitrateOneZeroEightZerop),
		string(EncoderNamedPresetHTwoSixFiveSingleBitrateSevenTwoZerop),
		string(EncoderNamedPresetHTwoSixFourMultipleBitrateOneZeroEightZerop),
		string(EncoderNamedPresetHTwoSixFourMultipleBitrateSD),
		string(EncoderNamedPresetHTwoSixFourMultipleBitrateSevenTwoZerop),
		string(EncoderNamedPresetHTwoSixFourSingleBitrateOneZeroEightZerop),
		string(EncoderNamedPresetHTwoSixFourSingleBitrateSD),
		string(EncoderNamedPresetHTwoSixFourSingleBitrateSevenTwoZerop),
	}
}

func (s *EncoderNamedPreset) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForEncoderNamedPreset() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = EncoderNamedPreset(decoded)
	return nil
}

type EntropyMode string

const (
	EntropyModeCabac EntropyMode = "Cabac"
	EntropyModeCavlc EntropyMode = "Cavlc"
)

func PossibleValuesForEntropyMode() []string {
	return []string{
		string(EntropyModeCabac),
		string(EntropyModeCavlc),
	}
}

func (s *EntropyMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForEntropyMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = EntropyMode(decoded)
	return nil
}

type FaceRedactorMode string

const (
	FaceRedactorModeAnalyze  FaceRedactorMode = "Analyze"
	FaceRedactorModeCombined FaceRedactorMode = "Combined"
	FaceRedactorModeRedact   FaceRedactorMode = "Redact"
)

func PossibleValuesForFaceRedactorMode() []string {
	return []string{
		string(FaceRedactorModeAnalyze),
		string(FaceRedactorModeCombined),
		string(FaceRedactorModeRedact),
	}
}

func (s *FaceRedactorMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForFaceRedactorMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = FaceRedactorMode(decoded)
	return nil
}

type H264Complexity string

const (
	H264ComplexityBalanced H264Complexity = "Balanced"
	H264ComplexityQuality  H264Complexity = "Quality"
	H264ComplexitySpeed    H264Complexity = "Speed"
)

func PossibleValuesForH264Complexity() []string {
	return []string{
		string(H264ComplexityBalanced),
		string(H264ComplexityQuality),
		string(H264ComplexitySpeed),
	}
}

func (s *H264Complexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForH264Complexity() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = H264Complexity(decoded)
	return nil
}

type H264RateControlMode string

const (
	H264RateControlModeABR H264RateControlMode = "ABR"
	H264RateControlModeCBR H264RateControlMode = "CBR"
	H264RateControlModeCRF H264RateControlMode = "CRF"
)

func PossibleValuesForH264RateControlMode() []string {
	return []string{
		string(H264RateControlModeABR),
		string(H264RateControlModeCBR),
		string(H264RateControlModeCRF),
	}
}

func (s *H264RateControlMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForH264RateControlMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = H264RateControlMode(decoded)
	return nil
}

type H264VideoProfile string

const (
	H264VideoProfileAuto             H264VideoProfile = "Auto"
	H264VideoProfileBaseline         H264VideoProfile = "Baseline"
	H264VideoProfileHigh             H264VideoProfile = "High"
	H264VideoProfileHighFourFourFour H264VideoProfile = "High444"
	H264VideoProfileHighFourTwoTwo   H264VideoProfile = "High422"
	H264VideoProfileMain             H264VideoProfile = "Main"
)

func PossibleValuesForH264VideoProfile() []string {
	return []string{
		string(H264VideoProfileAuto),
		string(H264VideoProfileBaseline),
		string(H264VideoProfileHigh),
		string(H264VideoProfileHighFourFourFour),
		string(H264VideoProfileHighFourTwoTwo),
		string(H264VideoProfileMain),
	}
}

func (s *H264VideoProfile) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForH264VideoProfile() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = H264VideoProfile(decoded)
	return nil
}

type H265Complexity string

const (
	H265ComplexityBalanced H265Complexity = "Balanced"
	H265ComplexityQuality  H265Complexity = "Quality"
	H265ComplexitySpeed    H265Complexity = "Speed"
)

func PossibleValuesForH265Complexity() []string {
	return []string{
		string(H265ComplexityBalanced),
		string(H265ComplexityQuality),
		string(H265ComplexitySpeed),
	}
}

func (s *H265Complexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForH265Complexity() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = H265Complexity(decoded)
	return nil
}

type H265VideoProfile string

const (
	H265VideoProfileAuto        H265VideoProfile = "Auto"
	H265VideoProfileMain        H265VideoProfile = "Main"
	H265VideoProfileMainOneZero H265VideoProfile = "Main10"
)

func PossibleValuesForH265VideoProfile() []string {
	return []string{
		string(H265VideoProfileAuto),
		string(H265VideoProfileMain),
		string(H265VideoProfileMainOneZero),
	}
}

func (s *H265VideoProfile) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForH265VideoProfile() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = H265VideoProfile(decoded)
	return nil
}

type InsightsType string

const (
	InsightsTypeAllInsights       InsightsType = "AllInsights"
	InsightsTypeAudioInsightsOnly InsightsType = "AudioInsightsOnly"
	InsightsTypeVideoInsightsOnly InsightsType = "VideoInsightsOnly"
)

func PossibleValuesForInsightsType() []string {
	return []string{
		string(InsightsTypeAllInsights),
		string(InsightsTypeAudioInsightsOnly),
		string(InsightsTypeVideoInsightsOnly),
	}
}

func (s *InsightsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForInsightsType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = InsightsType(decoded)
	return nil
}

type InterleaveOutput string

const (
	InterleaveOutputInterleavedOutput    InterleaveOutput = "InterleavedOutput"
	InterleaveOutputNonInterleavedOutput InterleaveOutput = "NonInterleavedOutput"
)

func PossibleValuesForInterleaveOutput() []string {
	return []string{
		string(InterleaveOutputInterleavedOutput),
		string(InterleaveOutputNonInterleavedOutput),
	}
}

func (s *InterleaveOutput) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForInterleaveOutput() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = InterleaveOutput(decoded)
	return nil
}

type JobErrorCategory string

const (
	JobErrorCategoryConfiguration JobErrorCategory = "Configuration"
	JobErrorCategoryContent       JobErrorCategory = "Content"
	JobErrorCategoryDownload      JobErrorCategory = "Download"
	JobErrorCategoryService       JobErrorCategory = "Service"
	JobErrorCategoryUpload        JobErrorCategory = "Upload"
)

func PossibleValuesForJobErrorCategory() []string {
	return []string{
		string(JobErrorCategoryConfiguration),
		string(JobErrorCategoryContent),
		string(JobErrorCategoryDownload),
		string(JobErrorCategoryService),
		string(JobErrorCategoryUpload),
	}
}

func (s *JobErrorCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForJobErrorCategory() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = JobErrorCategory(decoded)
	return nil
}

type JobErrorCode string

const (
	JobErrorCodeConfigurationUnsupported JobErrorCode = "ConfigurationUnsupported"
	JobErrorCodeContentMalformed         JobErrorCode = "ContentMalformed"
	JobErrorCodeContentUnsupported       JobErrorCode = "ContentUnsupported"
	JobErrorCodeDownloadNotAccessible    JobErrorCode = "DownloadNotAccessible"
	JobErrorCodeDownloadTransientError   JobErrorCode = "DownloadTransientError"
	JobErrorCodeServiceError             JobErrorCode = "ServiceError"
	JobErrorCodeServiceTransientError    JobErrorCode = "ServiceTransientError"
	JobErrorCodeUploadNotAccessible      JobErrorCode = "UploadNotAccessible"
	JobErrorCodeUploadTransientError     JobErrorCode = "UploadTransientError"
)

func PossibleValuesForJobErrorCode() []string {
	return []string{
		string(JobErrorCodeConfigurationUnsupported),
		string(JobErrorCodeContentMalformed),
		string(JobErrorCodeContentUnsupported),
		string(JobErrorCodeDownloadNotAccessible),
		string(JobErrorCodeDownloadTransientError),
		string(JobErrorCodeServiceError),
		string(JobErrorCodeServiceTransientError),
		string(JobErrorCodeUploadNotAccessible),
		string(JobErrorCodeUploadTransientError),
	}
}

func (s *JobErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForJobErrorCode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = JobErrorCode(decoded)
	return nil
}

type JobRetry string

const (
	JobRetryDoNotRetry JobRetry = "DoNotRetry"
	JobRetryMayRetry   JobRetry = "MayRetry"
)

func PossibleValuesForJobRetry() []string {
	return []string{
		string(JobRetryDoNotRetry),
		string(JobRetryMayRetry),
	}
}

func (s *JobRetry) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForJobRetry() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = JobRetry(decoded)
	return nil
}

type JobState string

const (
	JobStateCanceled   JobState = "Canceled"
	JobStateCanceling  JobState = "Canceling"
	JobStateError      JobState = "Error"
	JobStateFinished   JobState = "Finished"
	JobStateProcessing JobState = "Processing"
	JobStateQueued     JobState = "Queued"
	JobStateScheduled  JobState = "Scheduled"
)

func PossibleValuesForJobState() []string {
	return []string{
		string(JobStateCanceled),
		string(JobStateCanceling),
		string(JobStateError),
		string(JobStateFinished),
		string(JobStateProcessing),
		string(JobStateQueued),
		string(JobStateScheduled),
	}
}

func (s *JobState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForJobState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = JobState(decoded)
	return nil
}

type OnErrorType string

const (
	OnErrorTypeContinueJob       OnErrorType = "ContinueJob"
	OnErrorTypeStopProcessingJob OnErrorType = "StopProcessingJob"
)

func PossibleValuesForOnErrorType() []string {
	return []string{
		string(OnErrorTypeContinueJob),
		string(OnErrorTypeStopProcessingJob),
	}
}

func (s *OnErrorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForOnErrorType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = OnErrorType(decoded)
	return nil
}

type Priority string

const (
	PriorityHigh   Priority = "High"
	PriorityLow    Priority = "Low"
	PriorityNormal Priority = "Normal"
)

func PossibleValuesForPriority() []string {
	return []string{
		string(PriorityHigh),
		string(PriorityLow),
		string(PriorityNormal),
	}
}

func (s *Priority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPriority() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = Priority(decoded)
	return nil
}

type Rotation string

const (
	RotationAuto               Rotation = "Auto"
	RotationNone               Rotation = "None"
	RotationRotateNineZero     Rotation = "Rotate90"
	RotationRotateOneEightZero Rotation = "Rotate180"
	RotationRotateTwoSevenZero Rotation = "Rotate270"
	RotationRotateZero         Rotation = "Rotate0"
)

func PossibleValuesForRotation() []string {
	return []string{
		string(RotationAuto),
		string(RotationNone),
		string(RotationRotateNineZero),
		string(RotationRotateOneEightZero),
		string(RotationRotateTwoSevenZero),
		string(RotationRotateZero),
	}
}

func (s *Rotation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForRotation() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = Rotation(decoded)
	return nil
}

type StretchMode string

const (
	StretchModeAutoFit  StretchMode = "AutoFit"
	StretchModeAutoSize StretchMode = "AutoSize"
	StretchModeNone     StretchMode = "None"
)

func PossibleValuesForStretchMode() []string {
	return []string{
		string(StretchModeAutoFit),
		string(StretchModeAutoSize),
		string(StretchModeNone),
	}
}

func (s *StretchMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForStretchMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = StretchMode(decoded)
	return nil
}

type TrackAttribute string

const (
	TrackAttributeBitrate  TrackAttribute = "Bitrate"
	TrackAttributeLanguage TrackAttribute = "Language"
)

func PossibleValuesForTrackAttribute() []string {
	return []string{
		string(TrackAttributeBitrate),
		string(TrackAttributeLanguage),
	}
}

func (s *TrackAttribute) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTrackAttribute() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TrackAttribute(decoded)
	return nil
}

type VideoSyncMode string

const (
	VideoSyncModeAuto        VideoSyncMode = "Auto"
	VideoSyncModeCfr         VideoSyncMode = "Cfr"
	VideoSyncModePassthrough VideoSyncMode = "Passthrough"
	VideoSyncModeVfr         VideoSyncMode = "Vfr"
)

func PossibleValuesForVideoSyncMode() []string {
	return []string{
		string(VideoSyncModeAuto),
		string(VideoSyncModeCfr),
		string(VideoSyncModePassthrough),
		string(VideoSyncModeVfr),
	}
}

func (s *VideoSyncMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForVideoSyncMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = VideoSyncMode(decoded)
	return nil
}
