package deploymentinfo

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
