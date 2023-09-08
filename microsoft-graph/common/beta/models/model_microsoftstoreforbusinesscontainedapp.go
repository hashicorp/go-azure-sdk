package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftStoreForBusinessContainedApp struct {
	AppUserModelId *string `json:"appUserModelId,omitempty"`
	Id             *string `json:"id,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
}
