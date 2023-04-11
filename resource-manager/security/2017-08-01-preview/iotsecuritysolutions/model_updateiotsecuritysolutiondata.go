package iotsecuritysolutions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateIotSecuritySolutionData struct {
	RecommendationsConfiguration *[]RecommendationConfigurationProperties `json:"recommendationsConfiguration,omitempty"`
	Tags                         *map[string]string                       `json:"tags,omitempty"`
	UserDefinedResources         *UserDefinedResourcesProperties          `json:"userDefinedResources,omitempty"`
}
