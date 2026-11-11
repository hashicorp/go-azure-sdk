package dataprotections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportedFeature struct {
	ExposureControlledFeatures *[]string             `json:"exposureControlledFeatures,omitempty"`
	FeatureName                *string               `json:"featureName,omitempty"`
	SupportStatus              *FeatureSupportStatus `json:"supportStatus,omitempty"`
}
