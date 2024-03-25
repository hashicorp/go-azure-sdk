package raicontentfilters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiContentFilter struct {
	Description *string               `json:"description,omitempty"`
	FilterType  *RaiContentFilterType `json:"filterType,omitempty"`
	PolicyName  *string               `json:"policyName,omitempty"`
}
