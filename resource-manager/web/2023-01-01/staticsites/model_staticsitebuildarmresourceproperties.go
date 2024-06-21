package staticsites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSiteBuildARMResourceProperties struct {
	BuildId                  *string                              `json:"buildId,omitempty"`
	CreatedTimeUtc           *string                              `json:"createdTimeUtc,omitempty"`
	DatabaseConnections      *[]DatabaseConnectionOverview        `json:"databaseConnections,omitempty"`
	Hostname                 *string                              `json:"hostname,omitempty"`
	LastUpdatedOn            *string                              `json:"lastUpdatedOn,omitempty"`
	LinkedBackends           *[]StaticSiteLinkedBackend           `json:"linkedBackends,omitempty"`
	PullRequestTitle         *string                              `json:"pullRequestTitle,omitempty"`
	SourceBranch             *string                              `json:"sourceBranch,omitempty"`
	Status                   *BuildStatus                         `json:"status,omitempty"`
	UserProvidedFunctionApps *[]StaticSiteUserProvidedFunctionApp `json:"userProvidedFunctionApps,omitempty"`
}
