package serverrestart

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnableStatusEnum string

const (
	EnableStatusEnumDisabled EnableStatusEnum = "Disabled"
	EnableStatusEnumEnabled  EnableStatusEnum = "Enabled"
)

func PossibleValuesForEnableStatusEnum() []string {
	return []string{
		string(EnableStatusEnumDisabled),
		string(EnableStatusEnumEnabled),
	}
}
