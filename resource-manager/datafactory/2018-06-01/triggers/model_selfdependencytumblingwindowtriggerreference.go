package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelfDependencyTumblingWindowTriggerReference struct {
	Offset string  `json:"offset"`
	Size   *string `json:"size,omitempty"`
	Type   string  `json:"type"`
}
