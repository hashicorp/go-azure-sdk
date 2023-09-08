package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesClusterEvidencePlatform string

const (
	SecurityKubernetesClusterEvidencePlatformaks     SecurityKubernetesClusterEvidencePlatform = "Aks"
	SecurityKubernetesClusterEvidencePlatformarc     SecurityKubernetesClusterEvidencePlatform = "Arc"
	SecurityKubernetesClusterEvidencePlatformeks     SecurityKubernetesClusterEvidencePlatform = "Eks"
	SecurityKubernetesClusterEvidencePlatformgke     SecurityKubernetesClusterEvidencePlatform = "Gke"
	SecurityKubernetesClusterEvidencePlatformunknown SecurityKubernetesClusterEvidencePlatform = "Unknown"
)

func PossibleValuesForSecurityKubernetesClusterEvidencePlatform() []string {
	return []string{
		string(SecurityKubernetesClusterEvidencePlatformaks),
		string(SecurityKubernetesClusterEvidencePlatformarc),
		string(SecurityKubernetesClusterEvidencePlatformeks),
		string(SecurityKubernetesClusterEvidencePlatformgke),
		string(SecurityKubernetesClusterEvidencePlatformunknown),
	}
}

func parseSecurityKubernetesClusterEvidencePlatform(input string) (*SecurityKubernetesClusterEvidencePlatform, error) {
	vals := map[string]SecurityKubernetesClusterEvidencePlatform{
		"aks":     SecurityKubernetesClusterEvidencePlatformaks,
		"arc":     SecurityKubernetesClusterEvidencePlatformarc,
		"eks":     SecurityKubernetesClusterEvidencePlatformeks,
		"gke":     SecurityKubernetesClusterEvidencePlatformgke,
		"unknown": SecurityKubernetesClusterEvidencePlatformunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesClusterEvidencePlatform(input)
	return &out, nil
}
