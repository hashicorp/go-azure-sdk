package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationError struct {
	Code      *string `json:"code,omitempty"`
	Message   *string `json:"message,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
