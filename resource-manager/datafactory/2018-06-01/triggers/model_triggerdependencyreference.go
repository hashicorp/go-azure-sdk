package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerDependencyReference struct {
	ReferenceTrigger TriggerReference `json:"referenceTrigger"`
	Type             string           `json:"type"`
}
