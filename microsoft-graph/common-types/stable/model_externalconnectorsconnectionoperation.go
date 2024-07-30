package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsConnectionOperation struct {
	Error     *PublicError                                 `json:"error,omitempty"`
	Id        *string                                      `json:"id,omitempty"`
	ODataType *string                                      `json:"@odata.type,omitempty"`
	Status    *ExternalConnectorsConnectionOperationStatus `json:"status,omitempty"`
}
