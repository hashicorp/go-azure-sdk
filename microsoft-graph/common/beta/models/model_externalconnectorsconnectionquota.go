package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsConnectionQuota struct {
	Id             *string `json:"id,omitempty"`
	ItemsRemaining *int64  `json:"itemsRemaining,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
}
