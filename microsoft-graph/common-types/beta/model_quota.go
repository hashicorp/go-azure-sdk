package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Quota struct {
	Deleted                *int64                  `json:"deleted,omitempty"`
	ODataType              *string                 `json:"@odata.type,omitempty"`
	Remaining              *int64                  `json:"remaining,omitempty"`
	State                  *string                 `json:"state,omitempty"`
	StoragePlanInformation *StoragePlanInformation `json:"storagePlanInformation,omitempty"`
	Total                  *int64                  `json:"total,omitempty"`
	Used                   *int64                  `json:"used,omitempty"`
}
