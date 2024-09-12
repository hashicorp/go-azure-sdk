package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteArchivalDetails struct {
	// Represents the current archive status of the site collection. Returned only on $select. The possible values are:
	// recentlyArchived, fullyArchived, reactivating, unknownFutureValue.
	ArchiveStatus *SiteArchiveStatus `json:"archiveStatus,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
