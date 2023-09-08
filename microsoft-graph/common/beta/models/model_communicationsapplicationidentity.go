package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunicationsApplicationIdentity struct {
	ApplicationType *string `json:"applicationType,omitempty"`
	DisplayName     *string `json:"displayName,omitempty"`
	Hidden          *bool   `json:"hidden,omitempty"`
	Id              *string `json:"id,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
}
