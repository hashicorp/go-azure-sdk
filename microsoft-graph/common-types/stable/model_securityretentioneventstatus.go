package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionEventStatus struct {
	// The error if the status isn't successful.
	Error *PublicError `json:"error,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The status of the distribution. The possible values are: pending, error, success, notAvaliable.
	Status *SecurityEventStatusType `json:"status,omitempty"`
}
