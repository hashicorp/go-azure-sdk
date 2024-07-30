package conversationthreadpostinreplyto

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiValueLegacyExtendedProperty struct {
	Id        *string   `json:"id,omitempty"`
	ODataType *string   `json:"@odata.type,omitempty"`
	Value     *[]string `json:"value,omitempty"`
}
