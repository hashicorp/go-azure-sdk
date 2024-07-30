package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceEvidenceServiceType string

const (
	SecurityKubernetesServiceEvidenceServiceType_ClusterIP    SecurityKubernetesServiceEvidenceServiceType = "clusterIP"
	SecurityKubernetesServiceEvidenceServiceType_ExternalName SecurityKubernetesServiceEvidenceServiceType = "externalName"
	SecurityKubernetesServiceEvidenceServiceType_LoadBalancer SecurityKubernetesServiceEvidenceServiceType = "loadBalancer"
	SecurityKubernetesServiceEvidenceServiceType_NodePort     SecurityKubernetesServiceEvidenceServiceType = "nodePort"
	SecurityKubernetesServiceEvidenceServiceType_Unknown      SecurityKubernetesServiceEvidenceServiceType = "unknown"
)

func PossibleValuesForSecurityKubernetesServiceEvidenceServiceType() []string {
	return []string{
		string(SecurityKubernetesServiceEvidenceServiceType_ClusterIP),
		string(SecurityKubernetesServiceEvidenceServiceType_ExternalName),
		string(SecurityKubernetesServiceEvidenceServiceType_LoadBalancer),
		string(SecurityKubernetesServiceEvidenceServiceType_NodePort),
		string(SecurityKubernetesServiceEvidenceServiceType_Unknown),
	}
}

func (s *SecurityKubernetesServiceEvidenceServiceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesServiceEvidenceServiceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesServiceEvidenceServiceType(input string) (*SecurityKubernetesServiceEvidenceServiceType, error) {
	vals := map[string]SecurityKubernetesServiceEvidenceServiceType{
		"clusterip":    SecurityKubernetesServiceEvidenceServiceType_ClusterIP,
		"externalname": SecurityKubernetesServiceEvidenceServiceType_ExternalName,
		"loadbalancer": SecurityKubernetesServiceEvidenceServiceType_LoadBalancer,
		"nodeport":     SecurityKubernetesServiceEvidenceServiceType_NodePort,
		"unknown":      SecurityKubernetesServiceEvidenceServiceType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceEvidenceServiceType(input)
	return &out, nil
}
