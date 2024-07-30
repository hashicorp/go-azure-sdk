package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobSubject struct {
	Links          *SynchronizationLinkedObjects `json:"links,omitempty"`
	ODataType      *string                       `json:"@odata.type,omitempty"`
	ObjectId       *string                       `json:"objectId,omitempty"`
	ObjectTypeName *string                       `json:"objectTypeName,omitempty"`
}
