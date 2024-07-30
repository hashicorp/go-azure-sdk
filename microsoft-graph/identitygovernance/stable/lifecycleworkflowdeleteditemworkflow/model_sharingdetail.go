package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingDetail struct {
	ODataType        *string            `json:"@odata.type,omitempty"`
	SharedBy         *InsightIdentity   `json:"sharedBy,omitempty"`
	SharedDateTime   *string            `json:"sharedDateTime,omitempty"`
	SharingReference *ResourceReference `json:"sharingReference,omitempty"`
	SharingSubject   *string            `json:"sharingSubject,omitempty"`
	SharingType      *string            `json:"sharingType,omitempty"`
}
