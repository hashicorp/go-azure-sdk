package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoTaskList struct {
	DisplayName       *string                        `json:"displayName,omitempty"`
	Extensions        *[]Extension                   `json:"extensions,omitempty"`
	Id                *string                        `json:"id,omitempty"`
	IsOwner           *bool                          `json:"isOwner,omitempty"`
	IsShared          *bool                          `json:"isShared,omitempty"`
	ODataType         *string                        `json:"@odata.type,omitempty"`
	Tasks             *[]TodoTask                    `json:"tasks,omitempty"`
	WellknownListName *TodoTaskListWellknownListName `json:"wellknownListName,omitempty"`
}
