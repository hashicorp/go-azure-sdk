package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlProperties struct {
	ContentTypes           []ContentType           `json:"contentTypes"`
	Description            *string                 `json:"description,omitempty"`
	DisplayName            string                  `json:"displayName"`
	Id                     *string                 `json:"id,omitempty"`
	LastDeploymentInfo     *DeploymentInfo         `json:"lastDeploymentInfo"`
	RepoType               RepoType                `json:"repoType"`
	Repository             Repository              `json:"repository"`
	RepositoryResourceInfo *RepositoryResourceInfo `json:"repositoryResourceInfo"`
	Version                *Version                `json:"version,omitempty"`
}
