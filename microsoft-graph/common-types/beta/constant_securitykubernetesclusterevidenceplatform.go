package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesClusterEvidencePlatform string

const (
	SecurityKubernetesClusterEvidencePlatform_Aks     SecurityKubernetesClusterEvidencePlatform = "aks"
	SecurityKubernetesClusterEvidencePlatform_Arc     SecurityKubernetesClusterEvidencePlatform = "arc"
	SecurityKubernetesClusterEvidencePlatform_Eks     SecurityKubernetesClusterEvidencePlatform = "eks"
	SecurityKubernetesClusterEvidencePlatform_Gke     SecurityKubernetesClusterEvidencePlatform = "gke"
	SecurityKubernetesClusterEvidencePlatform_Unknown SecurityKubernetesClusterEvidencePlatform = "unknown"
)

func PossibleValuesForSecurityKubernetesClusterEvidencePlatform() []string {
	return []string{
		string(SecurityKubernetesClusterEvidencePlatform_Aks),
		string(SecurityKubernetesClusterEvidencePlatform_Arc),
		string(SecurityKubernetesClusterEvidencePlatform_Eks),
		string(SecurityKubernetesClusterEvidencePlatform_Gke),
		string(SecurityKubernetesClusterEvidencePlatform_Unknown),
	}
}

func (s *SecurityKubernetesClusterEvidencePlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesClusterEvidencePlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesClusterEvidencePlatform(input string) (*SecurityKubernetesClusterEvidencePlatform, error) {
	vals := map[string]SecurityKubernetesClusterEvidencePlatform{
		"aks":     SecurityKubernetesClusterEvidencePlatform_Aks,
		"arc":     SecurityKubernetesClusterEvidencePlatform_Arc,
		"eks":     SecurityKubernetesClusterEvidencePlatform_Eks,
		"gke":     SecurityKubernetesClusterEvidencePlatform_Gke,
		"unknown": SecurityKubernetesClusterEvidencePlatform_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesClusterEvidencePlatform(input)
	return &out, nil
}
