package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnProxyServer struct {
	Address                         *string `json:"address,omitempty"`
	AutomaticConfigurationScriptUrl *string `json:"automaticConfigurationScriptUrl,omitempty"`
	ODataType                       *string `json:"@odata.type,omitempty"`
	Port                            *int64  `json:"port,omitempty"`
}
