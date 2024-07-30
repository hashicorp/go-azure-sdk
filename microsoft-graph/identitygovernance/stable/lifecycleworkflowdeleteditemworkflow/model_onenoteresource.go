package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteResource struct {
	Content    *string `json:"content,omitempty"`
	ContentUrl *string `json:"contentUrl,omitempty"`
	Id         *string `json:"id,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
	Self       *string `json:"self,omitempty"`
}
