package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TumblingWindowTriggerDependencyReference struct {
	Offset           *string          `json:"offset,omitempty"`
	ReferenceTrigger TriggerReference `json:"referenceTrigger"`
	Size             *string          `json:"size,omitempty"`
	Type             string           `json:"type"`
}
