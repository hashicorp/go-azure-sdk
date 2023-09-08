package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Initiator struct {
	DisplayName   *string                 `json:"displayName,omitempty"`
	Id            *string                 `json:"id,omitempty"`
	InitiatorType *InitiatorInitiatorType `json:"initiatorType,omitempty"`
	ODataType     *string                 `json:"@odata.type,omitempty"`
}
