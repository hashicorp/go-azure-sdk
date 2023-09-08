package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServicePortProtocol string

const (
	SecurityKubernetesServicePortProtocolsctp SecurityKubernetesServicePortProtocol = "Sctp"
	SecurityKubernetesServicePortProtocoltcp  SecurityKubernetesServicePortProtocol = "Tcp"
	SecurityKubernetesServicePortProtocoludp  SecurityKubernetesServicePortProtocol = "Udp"
)

func PossibleValuesForSecurityKubernetesServicePortProtocol() []string {
	return []string{
		string(SecurityKubernetesServicePortProtocolsctp),
		string(SecurityKubernetesServicePortProtocoltcp),
		string(SecurityKubernetesServicePortProtocoludp),
	}
}

func parseSecurityKubernetesServicePortProtocol(input string) (*SecurityKubernetesServicePortProtocol, error) {
	vals := map[string]SecurityKubernetesServicePortProtocol{
		"sctp": SecurityKubernetesServicePortProtocolsctp,
		"tcp":  SecurityKubernetesServicePortProtocoltcp,
		"udp":  SecurityKubernetesServicePortProtocoludp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServicePortProtocol(input)
	return &out, nil
}
