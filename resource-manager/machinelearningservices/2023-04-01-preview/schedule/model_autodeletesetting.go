package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoDeleteSetting struct {
	Condition *AutoDeleteCondition `json:"condition,omitempty"`
	Value     *string              `json:"value,omitempty"`
}
