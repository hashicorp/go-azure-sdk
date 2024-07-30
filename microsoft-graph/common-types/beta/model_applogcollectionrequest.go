package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppLogCollectionRequest struct {
	CompletedDateTime *string                        `json:"completedDateTime,omitempty"`
	CustomLogFolders  *[]string                      `json:"customLogFolders,omitempty"`
	ErrorMessage      *string                        `json:"errorMessage,omitempty"`
	Id                *string                        `json:"id,omitempty"`
	ODataType         *string                        `json:"@odata.type,omitempty"`
	Status            *AppLogCollectionRequestStatus `json:"status,omitempty"`
}
