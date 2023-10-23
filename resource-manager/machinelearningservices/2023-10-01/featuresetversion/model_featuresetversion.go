package featuresetversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeaturesetVersion struct {
	Description             *string                  `json:"description,omitempty"`
	Entities                *[]string                `json:"entities,omitempty"`
	IsAnonymous             *bool                    `json:"isAnonymous,omitempty"`
	IsArchived              *bool                    `json:"isArchived,omitempty"`
	MaterializationSettings *MaterializationSettings `json:"materializationSettings,omitempty"`
	Properties              *map[string]string       `json:"properties,omitempty"`
	ProvisioningState       *AssetProvisioningState  `json:"provisioningState,omitempty"`
	Specification           *FeaturesetSpecification `json:"specification,omitempty"`
	Stage                   *string                  `json:"stage,omitempty"`
	Tags                    *map[string]string       `json:"tags,omitempty"`
}
