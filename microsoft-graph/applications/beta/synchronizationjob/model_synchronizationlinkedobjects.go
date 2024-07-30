package synchronizationjob

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationLinkedObjects struct {
	Manager   *SynchronizationJobSubject   `json:"manager,omitempty"`
	Members   *[]SynchronizationJobSubject `json:"members,omitempty"`
	ODataType *string                      `json:"@odata.type,omitempty"`
	Owners    *[]SynchronizationJobSubject `json:"owners,omitempty"`
}
