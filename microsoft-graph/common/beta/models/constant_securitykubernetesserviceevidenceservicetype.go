package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceEvidenceServiceType string

const (
	SecurityKubernetesServiceEvidenceServiceTypeclusterIP    SecurityKubernetesServiceEvidenceServiceType = "ClusterIP"
	SecurityKubernetesServiceEvidenceServiceTypeexternalName SecurityKubernetesServiceEvidenceServiceType = "ExternalName"
	SecurityKubernetesServiceEvidenceServiceTypeloadBalancer SecurityKubernetesServiceEvidenceServiceType = "LoadBalancer"
	SecurityKubernetesServiceEvidenceServiceTypenodePort     SecurityKubernetesServiceEvidenceServiceType = "NodePort"
	SecurityKubernetesServiceEvidenceServiceTypeunknown      SecurityKubernetesServiceEvidenceServiceType = "Unknown"
)

func PossibleValuesForSecurityKubernetesServiceEvidenceServiceType() []string {
	return []string{
		string(SecurityKubernetesServiceEvidenceServiceTypeclusterIP),
		string(SecurityKubernetesServiceEvidenceServiceTypeexternalName),
		string(SecurityKubernetesServiceEvidenceServiceTypeloadBalancer),
		string(SecurityKubernetesServiceEvidenceServiceTypenodePort),
		string(SecurityKubernetesServiceEvidenceServiceTypeunknown),
	}
}

func parseSecurityKubernetesServiceEvidenceServiceType(input string) (*SecurityKubernetesServiceEvidenceServiceType, error) {
	vals := map[string]SecurityKubernetesServiceEvidenceServiceType{
		"clusterip":    SecurityKubernetesServiceEvidenceServiceTypeclusterIP,
		"externalname": SecurityKubernetesServiceEvidenceServiceTypeexternalName,
		"loadbalancer": SecurityKubernetesServiceEvidenceServiceTypeloadBalancer,
		"nodeport":     SecurityKubernetesServiceEvidenceServiceTypenodePort,
		"unknown":      SecurityKubernetesServiceEvidenceServiceTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceEvidenceServiceType(input)
	return &out, nil
}
