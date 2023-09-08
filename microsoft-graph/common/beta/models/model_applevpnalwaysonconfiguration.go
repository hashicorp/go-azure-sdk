package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnAlwaysOnConfiguration struct {
	AirPrintExceptionAction       *AppleVpnAlwaysOnConfigurationAirPrintExceptionAction  `json:"airPrintExceptionAction,omitempty"`
	AllowAllCaptiveNetworkPlugins *bool                                                  `json:"allowAllCaptiveNetworkPlugins,omitempty"`
	AllowCaptiveWebSheet          *bool                                                  `json:"allowCaptiveWebSheet,omitempty"`
	AllowedCaptiveNetworkPlugins  *SpecifiedCaptiveNetworkPlugins                        `json:"allowedCaptiveNetworkPlugins,omitempty"`
	CellularExceptionAction       *AppleVpnAlwaysOnConfigurationCellularExceptionAction  `json:"cellularExceptionAction,omitempty"`
	NatKeepAliveIntervalInSeconds *int64                                                 `json:"natKeepAliveIntervalInSeconds,omitempty"`
	NatKeepAliveOffloadEnable     *bool                                                  `json:"natKeepAliveOffloadEnable,omitempty"`
	ODataType                     *string                                                `json:"@odata.type,omitempty"`
	TunnelConfiguration           *AppleVpnAlwaysOnConfigurationTunnelConfiguration      `json:"tunnelConfiguration,omitempty"`
	UserToggleEnabled             *bool                                                  `json:"userToggleEnabled,omitempty"`
	VoicemailExceptionAction      *AppleVpnAlwaysOnConfigurationVoicemailExceptionAction `json:"voicemailExceptionAction,omitempty"`
}
