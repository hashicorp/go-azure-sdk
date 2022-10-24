package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightQueryItemPropertiesDefaultTimeRange struct {
	AfterRange  *string `json:"afterRange,omitempty"`
	BeforeRange *string `json:"beforeRange,omitempty"`
}
