package querypacks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogAnalyticsQueryPackProperties struct {
	ProvisioningState *string `json:"provisioningState,omitempty"`
	QueryPackId       *string `json:"queryPackId,omitempty"`
	TimeCreated       *string `json:"timeCreated,omitempty"`
	TimeModified      *string `json:"timeModified,omitempty"`
}
