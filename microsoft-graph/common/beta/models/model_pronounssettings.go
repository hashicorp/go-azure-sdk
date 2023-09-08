package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PronounsSettings struct {
	Id                      *string `json:"id,omitempty"`
	IsEnabledInOrganization *bool   `json:"isEnabledInOrganization,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
}
