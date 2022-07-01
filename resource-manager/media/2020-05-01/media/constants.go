package media

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountEncryptionKeyType string

const (
	AccountEncryptionKeyTypeCustomerKey AccountEncryptionKeyType = "CustomerKey"
	AccountEncryptionKeyTypeSystemKey   AccountEncryptionKeyType = "SystemKey"
)

func PossibleValuesForAccountEncryptionKeyType() []string {
	return []string{
		string(AccountEncryptionKeyTypeCustomerKey),
		string(AccountEncryptionKeyTypeSystemKey),
	}
}

func parseAccountEncryptionKeyType(input string) (*AccountEncryptionKeyType, error) {
	vals := map[string]AccountEncryptionKeyType{
		"customerkey": AccountEncryptionKeyTypeCustomerKey,
		"systemkey":   AccountEncryptionKeyTypeSystemKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccountEncryptionKeyType(input)
	return &out, nil
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

func parseAnalysisResolution(input string) (*AnalysisResolution, error) {
	vals := map[string]AnalysisResolution{
		"sourceresolution":   AnalysisResolutionSourceResolution,
		"standarddefinition": AnalysisResolutionStandardDefinition,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AnalysisResolution(input)
	return &out, nil
}

type AssetContainerPermission string

const (
	AssetContainerPermissionRead            AssetContainerPermission = "Read"
	AssetContainerPermissionReadWrite       AssetContainerPermission = "ReadWrite"
	AssetContainerPermissionReadWriteDelete AssetContainerPermission = "ReadWriteDelete"
)

func PossibleValuesForAssetContainerPermission() []string {
	return []string{
		string(AssetContainerPermissionRead),
		string(AssetContainerPermissionReadWrite),
		string(AssetContainerPermissionReadWriteDelete),
	}
}

func parseAssetContainerPermission(input string) (*AssetContainerPermission, error) {
	vals := map[string]AssetContainerPermission{
		"read":            AssetContainerPermissionRead,
		"readwrite":       AssetContainerPermissionReadWrite,
		"readwritedelete": AssetContainerPermissionReadWriteDelete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssetContainerPermission(input)
	return &out, nil
}

type AssetStorageEncryptionFormat string

const (
	AssetStorageEncryptionFormatMediaStorageClientEncryption AssetStorageEncryptionFormat = "MediaStorageClientEncryption"
	AssetStorageEncryptionFormatNone                         AssetStorageEncryptionFormat = "None"
)

func PossibleValuesForAssetStorageEncryptionFormat() []string {
	return []string{
		string(AssetStorageEncryptionFormatMediaStorageClientEncryption),
		string(AssetStorageEncryptionFormatNone),
	}
}

func parseAssetStorageEncryptionFormat(input string) (*AssetStorageEncryptionFormat, error) {
	vals := map[string]AssetStorageEncryptionFormat{
		"mediastorageclientencryption": AssetStorageEncryptionFormatMediaStorageClientEncryption,
		"none":                         AssetStorageEncryptionFormatNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssetStorageEncryptionFormat(input)
	return &out, nil
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

func parseAudioAnalysisMode(input string) (*AudioAnalysisMode, error) {
	vals := map[string]AudioAnalysisMode{
		"basic":    AudioAnalysisModeBasic,
		"standard": AudioAnalysisModeStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AudioAnalysisMode(input)
	return &out, nil
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

func parseBlurType(input string) (*BlurType, error) {
	vals := map[string]BlurType{
		"black": BlurTypeBlack,
		"box":   BlurTypeBox,
		"high":  BlurTypeHigh,
		"low":   BlurTypeLow,
		"med":   BlurTypeMed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BlurType(input)
	return &out, nil
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

func parseChannelMapping(input string) (*ChannelMapping, error) {
	vals := map[string]ChannelMapping{
		"backleft":            ChannelMappingBackLeft,
		"backright":           ChannelMappingBackRight,
		"center":              ChannelMappingCenter,
		"frontleft":           ChannelMappingFrontLeft,
		"frontright":          ChannelMappingFrontRight,
		"lowfrequencyeffects": ChannelMappingLowFrequencyEffects,
		"stereoleft":          ChannelMappingStereoLeft,
		"stereoright":         ChannelMappingStereoRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChannelMapping(input)
	return &out, nil
}

type ContentKeyPolicyFairPlayRentalAndLeaseKeyType string

const (
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeDualExpiry          ContentKeyPolicyFairPlayRentalAndLeaseKeyType = "DualExpiry"
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentLimited   ContentKeyPolicyFairPlayRentalAndLeaseKeyType = "PersistentLimited"
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentUnlimited ContentKeyPolicyFairPlayRentalAndLeaseKeyType = "PersistentUnlimited"
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUndefined           ContentKeyPolicyFairPlayRentalAndLeaseKeyType = "Undefined"
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUnknown             ContentKeyPolicyFairPlayRentalAndLeaseKeyType = "Unknown"
)

func PossibleValuesForContentKeyPolicyFairPlayRentalAndLeaseKeyType() []string {
	return []string{
		string(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeDualExpiry),
		string(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentLimited),
		string(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentUnlimited),
		string(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUndefined),
		string(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUnknown),
	}
}

func parseContentKeyPolicyFairPlayRentalAndLeaseKeyType(input string) (*ContentKeyPolicyFairPlayRentalAndLeaseKeyType, error) {
	vals := map[string]ContentKeyPolicyFairPlayRentalAndLeaseKeyType{
		"dualexpiry":          ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeDualExpiry,
		"persistentlimited":   ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentLimited,
		"persistentunlimited": ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentUnlimited,
		"undefined":           ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUndefined,
		"unknown":             ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentKeyPolicyFairPlayRentalAndLeaseKeyType(input)
	return &out, nil
}

type ContentKeyPolicyPlayReadyContentType string

const (
	ContentKeyPolicyPlayReadyContentTypeUltraVioletDownload  ContentKeyPolicyPlayReadyContentType = "UltraVioletDownload"
	ContentKeyPolicyPlayReadyContentTypeUltraVioletStreaming ContentKeyPolicyPlayReadyContentType = "UltraVioletStreaming"
	ContentKeyPolicyPlayReadyContentTypeUnknown              ContentKeyPolicyPlayReadyContentType = "Unknown"
	ContentKeyPolicyPlayReadyContentTypeUnspecified          ContentKeyPolicyPlayReadyContentType = "Unspecified"
)

func PossibleValuesForContentKeyPolicyPlayReadyContentType() []string {
	return []string{
		string(ContentKeyPolicyPlayReadyContentTypeUltraVioletDownload),
		string(ContentKeyPolicyPlayReadyContentTypeUltraVioletStreaming),
		string(ContentKeyPolicyPlayReadyContentTypeUnknown),
		string(ContentKeyPolicyPlayReadyContentTypeUnspecified),
	}
}

func parseContentKeyPolicyPlayReadyContentType(input string) (*ContentKeyPolicyPlayReadyContentType, error) {
	vals := map[string]ContentKeyPolicyPlayReadyContentType{
		"ultravioletdownload":  ContentKeyPolicyPlayReadyContentTypeUltraVioletDownload,
		"ultravioletstreaming": ContentKeyPolicyPlayReadyContentTypeUltraVioletStreaming,
		"unknown":              ContentKeyPolicyPlayReadyContentTypeUnknown,
		"unspecified":          ContentKeyPolicyPlayReadyContentTypeUnspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentKeyPolicyPlayReadyContentType(input)
	return &out, nil
}

type ContentKeyPolicyPlayReadyLicenseType string

const (
	ContentKeyPolicyPlayReadyLicenseTypeNonPersistent ContentKeyPolicyPlayReadyLicenseType = "NonPersistent"
	ContentKeyPolicyPlayReadyLicenseTypePersistent    ContentKeyPolicyPlayReadyLicenseType = "Persistent"
	ContentKeyPolicyPlayReadyLicenseTypeUnknown       ContentKeyPolicyPlayReadyLicenseType = "Unknown"
)

func PossibleValuesForContentKeyPolicyPlayReadyLicenseType() []string {
	return []string{
		string(ContentKeyPolicyPlayReadyLicenseTypeNonPersistent),
		string(ContentKeyPolicyPlayReadyLicenseTypePersistent),
		string(ContentKeyPolicyPlayReadyLicenseTypeUnknown),
	}
}

func parseContentKeyPolicyPlayReadyLicenseType(input string) (*ContentKeyPolicyPlayReadyLicenseType, error) {
	vals := map[string]ContentKeyPolicyPlayReadyLicenseType{
		"nonpersistent": ContentKeyPolicyPlayReadyLicenseTypeNonPersistent,
		"persistent":    ContentKeyPolicyPlayReadyLicenseTypePersistent,
		"unknown":       ContentKeyPolicyPlayReadyLicenseTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentKeyPolicyPlayReadyLicenseType(input)
	return &out, nil
}

type ContentKeyPolicyPlayReadyUnknownOutputPassingOption string

const (
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowed                      ContentKeyPolicyPlayReadyUnknownOutputPassingOption = "Allowed"
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowedWithVideoConstriction ContentKeyPolicyPlayReadyUnknownOutputPassingOption = "AllowedWithVideoConstriction"
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionNotAllowed                   ContentKeyPolicyPlayReadyUnknownOutputPassingOption = "NotAllowed"
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionUnknown                      ContentKeyPolicyPlayReadyUnknownOutputPassingOption = "Unknown"
)

func PossibleValuesForContentKeyPolicyPlayReadyUnknownOutputPassingOption() []string {
	return []string{
		string(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowed),
		string(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowedWithVideoConstriction),
		string(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionNotAllowed),
		string(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionUnknown),
	}
}

func parseContentKeyPolicyPlayReadyUnknownOutputPassingOption(input string) (*ContentKeyPolicyPlayReadyUnknownOutputPassingOption, error) {
	vals := map[string]ContentKeyPolicyPlayReadyUnknownOutputPassingOption{
		"allowed":                      ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowed,
		"allowedwithvideoconstriction": ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowedWithVideoConstriction,
		"notallowed":                   ContentKeyPolicyPlayReadyUnknownOutputPassingOptionNotAllowed,
		"unknown":                      ContentKeyPolicyPlayReadyUnknownOutputPassingOptionUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentKeyPolicyPlayReadyUnknownOutputPassingOption(input)
	return &out, nil
}

type ContentKeyPolicyRestrictionTokenType string

const (
	ContentKeyPolicyRestrictionTokenTypeJwt     ContentKeyPolicyRestrictionTokenType = "Jwt"
	ContentKeyPolicyRestrictionTokenTypeSwt     ContentKeyPolicyRestrictionTokenType = "Swt"
	ContentKeyPolicyRestrictionTokenTypeUnknown ContentKeyPolicyRestrictionTokenType = "Unknown"
)

func PossibleValuesForContentKeyPolicyRestrictionTokenType() []string {
	return []string{
		string(ContentKeyPolicyRestrictionTokenTypeJwt),
		string(ContentKeyPolicyRestrictionTokenTypeSwt),
		string(ContentKeyPolicyRestrictionTokenTypeUnknown),
	}
}

func parseContentKeyPolicyRestrictionTokenType(input string) (*ContentKeyPolicyRestrictionTokenType, error) {
	vals := map[string]ContentKeyPolicyRestrictionTokenType{
		"jwt":     ContentKeyPolicyRestrictionTokenTypeJwt,
		"swt":     ContentKeyPolicyRestrictionTokenTypeSwt,
		"unknown": ContentKeyPolicyRestrictionTokenTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentKeyPolicyRestrictionTokenType(input)
	return &out, nil
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

func parseDeinterlaceMode(input string) (*DeinterlaceMode, error) {
	vals := map[string]DeinterlaceMode{
		"autopixeladaptive": DeinterlaceModeAutoPixelAdaptive,
		"off":               DeinterlaceModeOff,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeinterlaceMode(input)
	return &out, nil
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

func parseDeinterlaceParity(input string) (*DeinterlaceParity, error) {
	vals := map[string]DeinterlaceParity{
		"auto":             DeinterlaceParityAuto,
		"bottomfieldfirst": DeinterlaceParityBottomFieldFirst,
		"topfieldfirst":    DeinterlaceParityTopFieldFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeinterlaceParity(input)
	return &out, nil
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

func parseEncoderNamedPreset(input string) (*EncoderNamedPreset, error) {
	vals := map[string]EncoderNamedPreset{
		"aacgoodqualityaudio":              EncoderNamedPresetAACGoodQualityAudio,
		"adaptivestreaming":                EncoderNamedPresetAdaptiveStreaming,
		"contentawareencoding":             EncoderNamedPresetContentAwareEncoding,
		"contentawareencodingexperimental": EncoderNamedPresetContentAwareEncodingExperimental,
		"copyallbitratenoninterleaved":     EncoderNamedPresetCopyAllBitrateNonInterleaved,
		"h265adaptivestreaming":            EncoderNamedPresetHTwoSixFiveAdaptiveStreaming,
		"h265contentawareencoding":         EncoderNamedPresetHTwoSixFiveContentAwareEncoding,
		"h265singlebitrate4k":              EncoderNamedPresetHTwoSixFiveSingleBitrateFourK,
		"h265singlebitrate1080p":           EncoderNamedPresetHTwoSixFiveSingleBitrateOneZeroEightZerop,
		"h265singlebitrate720p":            EncoderNamedPresetHTwoSixFiveSingleBitrateSevenTwoZerop,
		"h264multiplebitrate1080p":         EncoderNamedPresetHTwoSixFourMultipleBitrateOneZeroEightZerop,
		"h264multiplebitratesd":            EncoderNamedPresetHTwoSixFourMultipleBitrateSD,
		"h264multiplebitrate720p":          EncoderNamedPresetHTwoSixFourMultipleBitrateSevenTwoZerop,
		"h264singlebitrate1080p":           EncoderNamedPresetHTwoSixFourSingleBitrateOneZeroEightZerop,
		"h264singlebitratesd":              EncoderNamedPresetHTwoSixFourSingleBitrateSD,
		"h264singlebitrate720p":            EncoderNamedPresetHTwoSixFourSingleBitrateSevenTwoZerop,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncoderNamedPreset(input)
	return &out, nil
}

type EncryptionScheme string

const (
	EncryptionSchemeCommonEncryptionCbcs EncryptionScheme = "CommonEncryptionCbcs"
	EncryptionSchemeCommonEncryptionCenc EncryptionScheme = "CommonEncryptionCenc"
	EncryptionSchemeEnvelopeEncryption   EncryptionScheme = "EnvelopeEncryption"
	EncryptionSchemeNoEncryption         EncryptionScheme = "NoEncryption"
)

func PossibleValuesForEncryptionScheme() []string {
	return []string{
		string(EncryptionSchemeCommonEncryptionCbcs),
		string(EncryptionSchemeCommonEncryptionCenc),
		string(EncryptionSchemeEnvelopeEncryption),
		string(EncryptionSchemeNoEncryption),
	}
}

func parseEncryptionScheme(input string) (*EncryptionScheme, error) {
	vals := map[string]EncryptionScheme{
		"commonencryptioncbcs": EncryptionSchemeCommonEncryptionCbcs,
		"commonencryptioncenc": EncryptionSchemeCommonEncryptionCenc,
		"envelopeencryption":   EncryptionSchemeEnvelopeEncryption,
		"noencryption":         EncryptionSchemeNoEncryption,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionScheme(input)
	return &out, nil
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

func parseFaceRedactorMode(input string) (*FaceRedactorMode, error) {
	vals := map[string]FaceRedactorMode{
		"analyze":  FaceRedactorModeAnalyze,
		"combined": FaceRedactorModeCombined,
		"redact":   FaceRedactorModeRedact,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FaceRedactorMode(input)
	return &out, nil
}

type FilterTrackPropertyCompareOperation string

const (
	FilterTrackPropertyCompareOperationEqual    FilterTrackPropertyCompareOperation = "Equal"
	FilterTrackPropertyCompareOperationNotEqual FilterTrackPropertyCompareOperation = "NotEqual"
)

func PossibleValuesForFilterTrackPropertyCompareOperation() []string {
	return []string{
		string(FilterTrackPropertyCompareOperationEqual),
		string(FilterTrackPropertyCompareOperationNotEqual),
	}
}

func parseFilterTrackPropertyCompareOperation(input string) (*FilterTrackPropertyCompareOperation, error) {
	vals := map[string]FilterTrackPropertyCompareOperation{
		"equal":    FilterTrackPropertyCompareOperationEqual,
		"notequal": FilterTrackPropertyCompareOperationNotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FilterTrackPropertyCompareOperation(input)
	return &out, nil
}

type FilterTrackPropertyType string

const (
	FilterTrackPropertyTypeBitrate  FilterTrackPropertyType = "Bitrate"
	FilterTrackPropertyTypeFourCC   FilterTrackPropertyType = "FourCC"
	FilterTrackPropertyTypeLanguage FilterTrackPropertyType = "Language"
	FilterTrackPropertyTypeName     FilterTrackPropertyType = "Name"
	FilterTrackPropertyTypeType     FilterTrackPropertyType = "Type"
	FilterTrackPropertyTypeUnknown  FilterTrackPropertyType = "Unknown"
)

func PossibleValuesForFilterTrackPropertyType() []string {
	return []string{
		string(FilterTrackPropertyTypeBitrate),
		string(FilterTrackPropertyTypeFourCC),
		string(FilterTrackPropertyTypeLanguage),
		string(FilterTrackPropertyTypeName),
		string(FilterTrackPropertyTypeType),
		string(FilterTrackPropertyTypeUnknown),
	}
}

func parseFilterTrackPropertyType(input string) (*FilterTrackPropertyType, error) {
	vals := map[string]FilterTrackPropertyType{
		"bitrate":  FilterTrackPropertyTypeBitrate,
		"fourcc":   FilterTrackPropertyTypeFourCC,
		"language": FilterTrackPropertyTypeLanguage,
		"name":     FilterTrackPropertyTypeName,
		"type":     FilterTrackPropertyTypeType,
		"unknown":  FilterTrackPropertyTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FilterTrackPropertyType(input)
	return &out, nil
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

func parseJobErrorCategory(input string) (*JobErrorCategory, error) {
	vals := map[string]JobErrorCategory{
		"configuration": JobErrorCategoryConfiguration,
		"content":       JobErrorCategoryContent,
		"download":      JobErrorCategoryDownload,
		"service":       JobErrorCategoryService,
		"upload":        JobErrorCategoryUpload,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobErrorCategory(input)
	return &out, nil
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

func parseJobErrorCode(input string) (*JobErrorCode, error) {
	vals := map[string]JobErrorCode{
		"configurationunsupported": JobErrorCodeConfigurationUnsupported,
		"contentmalformed":         JobErrorCodeContentMalformed,
		"contentunsupported":       JobErrorCodeContentUnsupported,
		"downloadnotaccessible":    JobErrorCodeDownloadNotAccessible,
		"downloadtransienterror":   JobErrorCodeDownloadTransientError,
		"serviceerror":             JobErrorCodeServiceError,
		"servicetransienterror":    JobErrorCodeServiceTransientError,
		"uploadnotaccessible":      JobErrorCodeUploadNotAccessible,
		"uploadtransienterror":     JobErrorCodeUploadTransientError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobErrorCode(input)
	return &out, nil
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

func parseJobRetry(input string) (*JobRetry, error) {
	vals := map[string]JobRetry{
		"donotretry": JobRetryDoNotRetry,
		"mayretry":   JobRetryMayRetry,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobRetry(input)
	return &out, nil
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

func parseJobState(input string) (*JobState, error) {
	vals := map[string]JobState{
		"canceled":   JobStateCanceled,
		"canceling":  JobStateCanceling,
		"error":      JobStateError,
		"finished":   JobStateFinished,
		"processing": JobStateProcessing,
		"queued":     JobStateQueued,
		"scheduled":  JobStateScheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobState(input)
	return &out, nil
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

func parseOnErrorType(input string) (*OnErrorType, error) {
	vals := map[string]OnErrorType{
		"continuejob":       OnErrorTypeContinueJob,
		"stopprocessingjob": OnErrorTypeStopProcessingJob,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnErrorType(input)
	return &out, nil
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

func parsePriority(input string) (*Priority, error) {
	vals := map[string]Priority{
		"high":   PriorityHigh,
		"low":    PriorityLow,
		"normal": PriorityNormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Priority(input)
	return &out, nil
}

type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

func PossibleValuesForPrivateEndpointConnectionProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProvisioningStateCreating),
		string(PrivateEndpointConnectionProvisioningStateDeleting),
		string(PrivateEndpointConnectionProvisioningStateFailed),
		string(PrivateEndpointConnectionProvisioningStateSucceeded),
	}
}

func parsePrivateEndpointConnectionProvisioningState(input string) (*PrivateEndpointConnectionProvisioningState, error) {
	vals := map[string]PrivateEndpointConnectionProvisioningState{
		"creating":  PrivateEndpointConnectionProvisioningStateCreating,
		"deleting":  PrivateEndpointConnectionProvisioningStateDeleting,
		"failed":    PrivateEndpointConnectionProvisioningStateFailed,
		"succeeded": PrivateEndpointConnectionProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateEndpointConnectionProvisioningState(input)
	return &out, nil
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointServiceConnectionStatus() []string {
	return []string{
		string(PrivateEndpointServiceConnectionStatusApproved),
		string(PrivateEndpointServiceConnectionStatusPending),
		string(PrivateEndpointServiceConnectionStatusRejected),
	}
}

func parsePrivateEndpointServiceConnectionStatus(input string) (*PrivateEndpointServiceConnectionStatus, error) {
	vals := map[string]PrivateEndpointServiceConnectionStatus{
		"approved": PrivateEndpointServiceConnectionStatusApproved,
		"pending":  PrivateEndpointServiceConnectionStatusPending,
		"rejected": PrivateEndpointServiceConnectionStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateEndpointServiceConnectionStatus(input)
	return &out, nil
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

func parseRotation(input string) (*Rotation, error) {
	vals := map[string]Rotation{
		"auto":      RotationAuto,
		"none":      RotationNone,
		"rotate90":  RotationRotateNineZero,
		"rotate180": RotationRotateOneEightZero,
		"rotate270": RotationRotateTwoSevenZero,
		"rotate0":   RotationRotateZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Rotation(input)
	return &out, nil
}

type StorageAccountType string

const (
	StorageAccountTypePrimary   StorageAccountType = "Primary"
	StorageAccountTypeSecondary StorageAccountType = "Secondary"
)

func PossibleValuesForStorageAccountType() []string {
	return []string{
		string(StorageAccountTypePrimary),
		string(StorageAccountTypeSecondary),
	}
}

func parseStorageAccountType(input string) (*StorageAccountType, error) {
	vals := map[string]StorageAccountType{
		"primary":   StorageAccountTypePrimary,
		"secondary": StorageAccountTypeSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAccountType(input)
	return &out, nil
}

type StorageAuthentication string

const (
	StorageAuthenticationManagedIdentity StorageAuthentication = "ManagedIdentity"
	StorageAuthenticationSystem          StorageAuthentication = "System"
)

func PossibleValuesForStorageAuthentication() []string {
	return []string{
		string(StorageAuthenticationManagedIdentity),
		string(StorageAuthenticationSystem),
	}
}

func parseStorageAuthentication(input string) (*StorageAuthentication, error) {
	vals := map[string]StorageAuthentication{
		"managedidentity": StorageAuthenticationManagedIdentity,
		"system":          StorageAuthenticationSystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAuthentication(input)
	return &out, nil
}

type StreamingLocatorContentKeyType string

const (
	StreamingLocatorContentKeyTypeCommonEncryptionCbcs StreamingLocatorContentKeyType = "CommonEncryptionCbcs"
	StreamingLocatorContentKeyTypeCommonEncryptionCenc StreamingLocatorContentKeyType = "CommonEncryptionCenc"
	StreamingLocatorContentKeyTypeEnvelopeEncryption   StreamingLocatorContentKeyType = "EnvelopeEncryption"
)

func PossibleValuesForStreamingLocatorContentKeyType() []string {
	return []string{
		string(StreamingLocatorContentKeyTypeCommonEncryptionCbcs),
		string(StreamingLocatorContentKeyTypeCommonEncryptionCenc),
		string(StreamingLocatorContentKeyTypeEnvelopeEncryption),
	}
}

func parseStreamingLocatorContentKeyType(input string) (*StreamingLocatorContentKeyType, error) {
	vals := map[string]StreamingLocatorContentKeyType{
		"commonencryptioncbcs": StreamingLocatorContentKeyTypeCommonEncryptionCbcs,
		"commonencryptioncenc": StreamingLocatorContentKeyTypeCommonEncryptionCenc,
		"envelopeencryption":   StreamingLocatorContentKeyTypeEnvelopeEncryption,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StreamingLocatorContentKeyType(input)
	return &out, nil
}

type StreamingPolicyStreamingProtocol string

const (
	StreamingPolicyStreamingProtocolDash            StreamingPolicyStreamingProtocol = "Dash"
	StreamingPolicyStreamingProtocolDownload        StreamingPolicyStreamingProtocol = "Download"
	StreamingPolicyStreamingProtocolHls             StreamingPolicyStreamingProtocol = "Hls"
	StreamingPolicyStreamingProtocolSmoothStreaming StreamingPolicyStreamingProtocol = "SmoothStreaming"
)

func PossibleValuesForStreamingPolicyStreamingProtocol() []string {
	return []string{
		string(StreamingPolicyStreamingProtocolDash),
		string(StreamingPolicyStreamingProtocolDownload),
		string(StreamingPolicyStreamingProtocolHls),
		string(StreamingPolicyStreamingProtocolSmoothStreaming),
	}
}

func parseStreamingPolicyStreamingProtocol(input string) (*StreamingPolicyStreamingProtocol, error) {
	vals := map[string]StreamingPolicyStreamingProtocol{
		"dash":            StreamingPolicyStreamingProtocolDash,
		"download":        StreamingPolicyStreamingProtocolDownload,
		"hls":             StreamingPolicyStreamingProtocolHls,
		"smoothstreaming": StreamingPolicyStreamingProtocolSmoothStreaming,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StreamingPolicyStreamingProtocol(input)
	return &out, nil
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

func parseStretchMode(input string) (*StretchMode, error) {
	vals := map[string]StretchMode{
		"autofit":  StretchModeAutoFit,
		"autosize": StretchModeAutoSize,
		"none":     StretchModeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StretchMode(input)
	return &out, nil
}

type TrackPropertyCompareOperation string

const (
	TrackPropertyCompareOperationEqual   TrackPropertyCompareOperation = "Equal"
	TrackPropertyCompareOperationUnknown TrackPropertyCompareOperation = "Unknown"
)

func PossibleValuesForTrackPropertyCompareOperation() []string {
	return []string{
		string(TrackPropertyCompareOperationEqual),
		string(TrackPropertyCompareOperationUnknown),
	}
}

func parseTrackPropertyCompareOperation(input string) (*TrackPropertyCompareOperation, error) {
	vals := map[string]TrackPropertyCompareOperation{
		"equal":   TrackPropertyCompareOperationEqual,
		"unknown": TrackPropertyCompareOperationUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrackPropertyCompareOperation(input)
	return &out, nil
}

type TrackPropertyType string

const (
	TrackPropertyTypeFourCC  TrackPropertyType = "FourCC"
	TrackPropertyTypeUnknown TrackPropertyType = "Unknown"
)

func PossibleValuesForTrackPropertyType() []string {
	return []string{
		string(TrackPropertyTypeFourCC),
		string(TrackPropertyTypeUnknown),
	}
}

func parseTrackPropertyType(input string) (*TrackPropertyType, error) {
	vals := map[string]TrackPropertyType{
		"fourcc":  TrackPropertyTypeFourCC,
		"unknown": TrackPropertyTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrackPropertyType(input)
	return &out, nil
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

func parseVideoSyncMode(input string) (*VideoSyncMode, error) {
	vals := map[string]VideoSyncMode{
		"auto":        VideoSyncModeAuto,
		"cfr":         VideoSyncModeCfr,
		"passthrough": VideoSyncModePassthrough,
		"vfr":         VideoSyncModeVfr,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VideoSyncMode(input)
	return &out, nil
}
