package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResponseStatus struct {
	ODataType *string                 `json:"@odata.type,omitempty"`
	Response  *ResponseStatusResponse `json:"response,omitempty"`
	Time      *string                 `json:"time,omitempty"`
}
