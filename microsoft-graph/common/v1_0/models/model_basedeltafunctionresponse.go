package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BaseDeltaFunctionResponse struct {
	ODataDeltaLink *string `json:"@odata.deltaLink,omitempty"`
	ODataNextLink  *string `json:"@odata.nextLink,omitempty"`
}
