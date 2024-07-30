package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThumbnailSet struct {
	Id        *string    `json:"id,omitempty"`
	Large     *Thumbnail `json:"large,omitempty"`
	Medium    *Thumbnail `json:"medium,omitempty"`
	ODataType *string    `json:"@odata.type,omitempty"`
	Small     *Thumbnail `json:"small,omitempty"`
	Source    *Thumbnail `json:"source,omitempty"`
}
