package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementsDataCollection struct {
	LastCollectionDateTime            *string                                                      `json:"lastCollectionDateTime,omitempty"`
	ODataType                         *string                                                      `json:"@odata.type,omitempty"`
	PermissionsModificationCapability *EntitlementsDataCollectionPermissionsModificationCapability `json:"permissionsModificationCapability,omitempty"`
	Status                            *EntitlementsDataCollectionStatus                            `json:"status,omitempty"`
}
