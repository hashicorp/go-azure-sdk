package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Configuration struct {
	AuthorizedAppIds *[]string `json:"authorizedAppIds,omitempty"`
	AuthorizedApps   *[]string `json:"authorizedApps,omitempty"`
	ODataType        *string   `json:"@odata.type,omitempty"`
}
