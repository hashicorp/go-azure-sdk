package connectordefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstructionStepDetails struct {
	Parameters interface{} `json:"parameters"`
	Type       string      `json:"type"`
}
