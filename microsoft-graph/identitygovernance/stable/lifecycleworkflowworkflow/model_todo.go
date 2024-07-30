package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Todo struct {
	Id        *string         `json:"id,omitempty"`
	Lists     *[]TodoTaskList `json:"lists,omitempty"`
	ODataType *string         `json:"@odata.type,omitempty"`
}
