package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentSetContent struct {
	ContentType *ContentTypeInfo `json:"contentType,omitempty"`
	FileName    *string          `json:"fileName,omitempty"`
	FolderName  *string          `json:"folderName,omitempty"`
	ODataType   *string          `json:"@odata.type,omitempty"`
}
