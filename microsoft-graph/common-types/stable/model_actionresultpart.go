package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionResultPart struct {
	Error     *PublicError `json:"error,omitempty"`
	ODataType *string      `json:"@odata.type,omitempty"`
}
