package deploymentinfo

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticDeploymentStatus string

const (
	ElasticDeploymentStatusHealthy   ElasticDeploymentStatus = "Healthy"
	ElasticDeploymentStatusUnhealthy ElasticDeploymentStatus = "Unhealthy"
)

func PossibleValuesForElasticDeploymentStatus() []string {
	return []string{
		string(ElasticDeploymentStatusHealthy),
		string(ElasticDeploymentStatusUnhealthy),
	}
}

func parseElasticDeploymentStatus(input string) (*ElasticDeploymentStatus, error) {
	vals := map[string]ElasticDeploymentStatus{
		"healthy":   ElasticDeploymentStatusHealthy,
		"unhealthy": ElasticDeploymentStatusUnhealthy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ElasticDeploymentStatus(input)
	return &out, nil
}
