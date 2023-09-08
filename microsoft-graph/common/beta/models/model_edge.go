package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Edge struct {
	Id                   *string               `json:"id,omitempty"`
	InternetExplorerMode *InternetExplorerMode `json:"internetExplorerMode,omitempty"`
	ODataType            *string               `json:"@odata.type,omitempty"`
}
