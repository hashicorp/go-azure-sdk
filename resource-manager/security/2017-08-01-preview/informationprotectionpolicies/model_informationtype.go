package informationprotectionpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationType struct {
	Custom             *bool                           `json:"custom,omitempty"`
	Description        *string                         `json:"description,omitempty"`
	DisplayName        *string                         `json:"displayName,omitempty"`
	Enabled            *bool                           `json:"enabled,omitempty"`
	Keywords           *[]InformationProtectionKeyword `json:"keywords,omitempty"`
	Order              *int64                          `json:"order,omitempty"`
	RecommendedLabelId *string                         `json:"recommendedLabelId,omitempty"`
}
