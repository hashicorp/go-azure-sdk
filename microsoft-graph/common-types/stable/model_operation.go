package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Operation struct {
	CreatedDateTime    *string          `json:"createdDateTime,omitempty"`
	Id                 *string          `json:"id,omitempty"`
	LastActionDateTime *string          `json:"lastActionDateTime,omitempty"`
	ODataType          *string          `json:"@odata.type,omitempty"`
	Status             *OperationStatus `json:"status,omitempty"`
}
