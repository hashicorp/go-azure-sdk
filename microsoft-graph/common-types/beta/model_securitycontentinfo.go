package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContentInfo struct {
	ContentFormat *string                   `json:"contentFormat,omitempty"`
	Identifier    *string                   `json:"identifier,omitempty"`
	Metadata      *[]SecurityKeyValuePair   `json:"metadata,omitempty"`
	ODataType     *string                   `json:"@odata.type,omitempty"`
	State         *SecurityContentInfoState `json:"state,omitempty"`
}
