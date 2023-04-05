package dppfeaturesupport

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureSupportStatus string

const (
	FeatureSupportStatusAlphaPreview       FeatureSupportStatus = "AlphaPreview"
	FeatureSupportStatusGenerallyAvailable FeatureSupportStatus = "GenerallyAvailable"
	FeatureSupportStatusInvalid            FeatureSupportStatus = "Invalid"
	FeatureSupportStatusNotSupported       FeatureSupportStatus = "NotSupported"
	FeatureSupportStatusPrivatePreview     FeatureSupportStatus = "PrivatePreview"
	FeatureSupportStatusPublicPreview      FeatureSupportStatus = "PublicPreview"
)

func PossibleValuesForFeatureSupportStatus() []string {
	return []string{
		string(FeatureSupportStatusAlphaPreview),
		string(FeatureSupportStatusGenerallyAvailable),
		string(FeatureSupportStatusInvalid),
		string(FeatureSupportStatusNotSupported),
		string(FeatureSupportStatusPrivatePreview),
		string(FeatureSupportStatusPublicPreview),
	}
}

type FeatureType string

const (
	FeatureTypeDataSourceType FeatureType = "DataSourceType"
	FeatureTypeInvalid        FeatureType = "Invalid"
)

func PossibleValuesForFeatureType() []string {
	return []string{
		string(FeatureTypeDataSourceType),
		string(FeatureTypeInvalid),
	}
}
