package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VpnProxyServer = Windows10VpnProxyServer{}

type Windows10VpnProxyServer struct {
	// Bypass proxy server for local address.
	BypassProxyServerForLocalAddress *bool `json:"bypassProxyServerForLocalAddress,omitempty"`

	// Fields inherited from VpnProxyServer

	// Address.
	Address nullable.Type[string] `json:"address,omitempty"`

	// Proxy's automatic configuration script url.
	AutomaticConfigurationScriptUrl nullable.Type[string] `json:"automaticConfigurationScriptUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Port. Valid values 0 to 65535
	Port nullable.Type[int64] `json:"port,omitempty"`
}

func (s Windows10VpnProxyServer) VpnProxyServer() BaseVpnProxyServerImpl {
	return BaseVpnProxyServerImpl{
		Address:                         s.Address,
		AutomaticConfigurationScriptUrl: s.AutomaticConfigurationScriptUrl,
		ODataId:                         s.ODataId,
		ODataType:                       s.ODataType,
		Port:                            s.Port,
	}
}

var _ json.Marshaler = Windows10VpnProxyServer{}

func (s Windows10VpnProxyServer) MarshalJSON() ([]byte, error) {
	type wrapper Windows10VpnProxyServer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10VpnProxyServer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10VpnProxyServer: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windows10VpnProxyServer"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10VpnProxyServer: %+v", err)
	}

	return encoded, nil
}
