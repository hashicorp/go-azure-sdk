package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataPolicyOperation struct {
	CompletedDateTime *string                    `json:"completedDateTime,omitempty"`
	Id                *string                    `json:"id,omitempty"`
	ODataType         *string                    `json:"@odata.type,omitempty"`
	Status            *DataPolicyOperationStatus `json:"status,omitempty"`
	StorageLocation   *string                    `json:"storageLocation,omitempty"`
	SubmittedDateTime *string                    `json:"submittedDateTime,omitempty"`
	UserId            *string                    `json:"userId,omitempty"`
}
