package appserviceenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CsmUsageQuota struct {
	CurrentValue  *int64             `json:"currentValue,omitempty"`
	Limit         *int64             `json:"limit,omitempty"`
	Name          *LocalizableString `json:"name,omitempty"`
	NextResetTime *string            `json:"nextResetTime,omitempty"`
	Unit          *string            `json:"unit,omitempty"`
}
