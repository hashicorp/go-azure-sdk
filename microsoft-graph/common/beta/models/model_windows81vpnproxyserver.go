package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81VpnProxyServer struct {
	Address                          *string `json:"address,omitempty"`
	AutomaticConfigurationScriptUrl  *string `json:"automaticConfigurationScriptUrl,omitempty"`
	AutomaticallyDetectProxySettings *bool   `json:"automaticallyDetectProxySettings,omitempty"`
	BypassProxyServerForLocalAddress *bool   `json:"bypassProxyServerForLocalAddress,omitempty"`
	ODataType                        *string `json:"@odata.type,omitempty"`
	Port                             *int64  `json:"port,omitempty"`
}
