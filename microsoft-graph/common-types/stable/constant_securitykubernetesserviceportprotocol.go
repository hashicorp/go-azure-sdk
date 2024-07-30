package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServicePortProtocol string

const (
	SecurityKubernetesServicePortProtocol_Sctp SecurityKubernetesServicePortProtocol = "sctp"
	SecurityKubernetesServicePortProtocol_Tcp  SecurityKubernetesServicePortProtocol = "tcp"
	SecurityKubernetesServicePortProtocol_Udp  SecurityKubernetesServicePortProtocol = "udp"
)

func PossibleValuesForSecurityKubernetesServicePortProtocol() []string {
	return []string{
		string(SecurityKubernetesServicePortProtocol_Sctp),
		string(SecurityKubernetesServicePortProtocol_Tcp),
		string(SecurityKubernetesServicePortProtocol_Udp),
	}
}

func (s *SecurityKubernetesServicePortProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesServicePortProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesServicePortProtocol(input string) (*SecurityKubernetesServicePortProtocol, error) {
	vals := map[string]SecurityKubernetesServicePortProtocol{
		"sctp": SecurityKubernetesServicePortProtocol_Sctp,
		"tcp":  SecurityKubernetesServicePortProtocol_Tcp,
		"udp":  SecurityKubernetesServicePortProtocol_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServicePortProtocol(input)
	return &out, nil
}
