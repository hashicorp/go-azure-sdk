package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusBase struct {
	ODataType *string           `json:"@odata.type,omitempty"`
	Status    *StatusBaseStatus `json:"status,omitempty"`
}
