package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintTaskStatus struct {
	Description *string               `json:"description,omitempty"`
	ODataType   *string               `json:"@odata.type,omitempty"`
	State       *PrintTaskStatusState `json:"state,omitempty"`
}
