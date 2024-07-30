package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteOperation struct {
	CreatedDateTime    *string                 `json:"createdDateTime,omitempty"`
	Error              *OnenoteOperationError  `json:"error,omitempty"`
	Id                 *string                 `json:"id,omitempty"`
	LastActionDateTime *string                 `json:"lastActionDateTime,omitempty"`
	ODataType          *string                 `json:"@odata.type,omitempty"`
	PercentComplete    *string                 `json:"percentComplete,omitempty"`
	ResourceId         *string                 `json:"resourceId,omitempty"`
	ResourceLocation   *string                 `json:"resourceLocation,omitempty"`
	Status             *OnenoteOperationStatus `json:"status,omitempty"`
}
