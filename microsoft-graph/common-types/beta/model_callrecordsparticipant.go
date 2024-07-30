package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsParticipant struct {
	Id        *string                  `json:"id,omitempty"`
	Identity  *CallRecordsUserIdentity `json:"identity,omitempty"`
	ODataType *string                  `json:"@odata.type,omitempty"`
}
