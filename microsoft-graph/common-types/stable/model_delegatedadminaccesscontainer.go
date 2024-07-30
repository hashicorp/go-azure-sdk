package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminAccessContainer struct {
	AccessContainerId   *string                                           `json:"accessContainerId,omitempty"`
	AccessContainerType *DelegatedAdminAccessContainerAccessContainerType `json:"accessContainerType,omitempty"`
	ODataType           *string                                           `json:"@odata.type,omitempty"`
}
