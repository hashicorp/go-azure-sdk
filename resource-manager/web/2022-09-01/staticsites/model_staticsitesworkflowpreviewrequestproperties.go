package staticsites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesWorkflowPreviewRequestProperties struct {
	Branch          *string                    `json:"branch,omitempty"`
	BuildProperties *StaticSiteBuildProperties `json:"buildProperties,omitempty"`
	RepositoryURL   *string                    `json:"repositoryUrl,omitempty"`
}
