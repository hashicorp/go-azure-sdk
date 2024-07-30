package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookOperation struct {
	Error            *WorkbookOperationError  `json:"error,omitempty"`
	Id               *string                  `json:"id,omitempty"`
	ODataType        *string                  `json:"@odata.type,omitempty"`
	ResourceLocation *string                  `json:"resourceLocation,omitempty"`
	Status           *WorkbookOperationStatus `json:"status,omitempty"`
}
