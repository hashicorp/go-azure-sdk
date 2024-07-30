package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintTask struct {
	Definition *PrintTaskDefinition `json:"definition,omitempty"`
	Id         *string              `json:"id,omitempty"`
	ODataType  *string              `json:"@odata.type,omitempty"`
	ParentUrl  *string              `json:"parentUrl,omitempty"`
	Status     *PrintTaskStatus     `json:"status,omitempty"`
	Trigger    *PrintTaskTrigger    `json:"trigger,omitempty"`
}
