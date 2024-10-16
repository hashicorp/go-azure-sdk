package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidationThreshold struct {
	Grouping ValidationThresholdGrouping `json:"grouping"`
	Type     ValidationThresholdType     `json:"type"`
	Value    int64                       `json:"value"`
}
