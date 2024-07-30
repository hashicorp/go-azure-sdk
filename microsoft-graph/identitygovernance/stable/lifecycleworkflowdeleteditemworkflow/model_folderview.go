package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FolderView struct {
	ODataType *string `json:"@odata.type,omitempty"`
	SortBy    *string `json:"sortBy,omitempty"`
	SortOrder *string `json:"sortOrder,omitempty"`
	ViewType  *string `json:"viewType,omitempty"`
}
