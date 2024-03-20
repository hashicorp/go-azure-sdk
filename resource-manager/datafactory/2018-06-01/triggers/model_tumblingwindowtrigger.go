package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TumblingWindowTrigger struct {
	Annotations    *[]interface{}                      `json:"annotations,omitempty"`
	Description    *string                             `json:"description,omitempty"`
	Pipeline       TriggerPipelineReference            `json:"pipeline"`
	RuntimeState   *TriggerRuntimeState                `json:"runtimeState,omitempty"`
	Type           string                              `json:"type"`
	TypeProperties TumblingWindowTriggerTypeProperties `json:"typeProperties"`
}
