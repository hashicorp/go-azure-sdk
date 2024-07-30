package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRemoteNetworkHealthEventStatus string

const (
	NetworkaccessRemoteNetworkHealthEventStatus_BgpConnected       NetworkaccessRemoteNetworkHealthEventStatus = "bgpConnected"
	NetworkaccessRemoteNetworkHealthEventStatus_BgpDisconnected    NetworkaccessRemoteNetworkHealthEventStatus = "bgpDisconnected"
	NetworkaccessRemoteNetworkHealthEventStatus_RemoteNetworkAlive NetworkaccessRemoteNetworkHealthEventStatus = "remoteNetworkAlive"
	NetworkaccessRemoteNetworkHealthEventStatus_TunnelConnected    NetworkaccessRemoteNetworkHealthEventStatus = "tunnelConnected"
	NetworkaccessRemoteNetworkHealthEventStatus_TunnelDisconnected NetworkaccessRemoteNetworkHealthEventStatus = "tunnelDisconnected"
)

func PossibleValuesForNetworkaccessRemoteNetworkHealthEventStatus() []string {
	return []string{
		string(NetworkaccessRemoteNetworkHealthEventStatus_BgpConnected),
		string(NetworkaccessRemoteNetworkHealthEventStatus_BgpDisconnected),
		string(NetworkaccessRemoteNetworkHealthEventStatus_RemoteNetworkAlive),
		string(NetworkaccessRemoteNetworkHealthEventStatus_TunnelConnected),
		string(NetworkaccessRemoteNetworkHealthEventStatus_TunnelDisconnected),
	}
}

func (s *NetworkaccessRemoteNetworkHealthEventStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessRemoteNetworkHealthEventStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessRemoteNetworkHealthEventStatus(input string) (*NetworkaccessRemoteNetworkHealthEventStatus, error) {
	vals := map[string]NetworkaccessRemoteNetworkHealthEventStatus{
		"bgpconnected":       NetworkaccessRemoteNetworkHealthEventStatus_BgpConnected,
		"bgpdisconnected":    NetworkaccessRemoteNetworkHealthEventStatus_BgpDisconnected,
		"remotenetworkalive": NetworkaccessRemoteNetworkHealthEventStatus_RemoteNetworkAlive,
		"tunnelconnected":    NetworkaccessRemoteNetworkHealthEventStatus_TunnelConnected,
		"tunneldisconnected": NetworkaccessRemoteNetworkHealthEventStatus_TunnelDisconnected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessRemoteNetworkHealthEventStatus(input)
	return &out, nil
}
