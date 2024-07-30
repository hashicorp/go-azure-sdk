package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintTaskTrigger struct {
	Definition *PrintTaskDefinition   `json:"definition,omitempty"`
	Event      *PrintTaskTriggerEvent `json:"event,omitempty"`
	Id         *string                `json:"id,omitempty"`
	ODataType  *string                `json:"@odata.type,omitempty"`
}
