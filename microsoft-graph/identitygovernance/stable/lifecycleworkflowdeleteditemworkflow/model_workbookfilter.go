package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookFilter struct {
	Criteria  *WorkbookFilterCriteria `json:"criteria,omitempty"`
	Id        *string                 `json:"id,omitempty"`
	ODataType *string                 `json:"@odata.type,omitempty"`
}
