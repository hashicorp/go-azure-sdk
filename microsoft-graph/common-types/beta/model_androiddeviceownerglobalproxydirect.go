package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGlobalProxyDirect struct {
	ExcludedHosts *[]string `json:"excludedHosts,omitempty"`
	Host          *string   `json:"host,omitempty"`
	ODataType     *string   `json:"@odata.type,omitempty"`
	Port          *int64    `json:"port,omitempty"`
}
