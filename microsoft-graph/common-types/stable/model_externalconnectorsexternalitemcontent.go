package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalItemContent struct {
	ODataType *string                                    `json:"@odata.type,omitempty"`
	Type      *ExternalConnectorsExternalItemContentType `json:"type,omitempty"`
	Value     *string                                    `json:"value,omitempty"`
}
