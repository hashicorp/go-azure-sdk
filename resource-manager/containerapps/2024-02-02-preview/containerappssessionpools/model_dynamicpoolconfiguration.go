package containerappssessionpools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DynamicPoolConfiguration struct {
	CooldownPeriodInSeconds *int64         `json:"cooldownPeriodInSeconds,omitempty"`
	ExecutionType           *ExecutionType `json:"executionType,omitempty"`
}
