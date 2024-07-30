package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonType struct {
	Class     *string `json:"class,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Subclass  *string `json:"subclass,omitempty"`
}
