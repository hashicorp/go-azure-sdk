package featuresetversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeaturesetVersionBackfillRequest struct {
	Description        *string                         `json:"description,omitempty"`
	DisplayName        *string                         `json:"displayName,omitempty"`
	FeatureWindow      *FeatureWindow                  `json:"featureWindow,omitempty"`
	Resource           *MaterializationComputeResource `json:"resource,omitempty"`
	SparkConfiguration *map[string]string              `json:"sparkConfiguration,omitempty"`
	Tags               *map[string]string              `json:"tags,omitempty"`
}
