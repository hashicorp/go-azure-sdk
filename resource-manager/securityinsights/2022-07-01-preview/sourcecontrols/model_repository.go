package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Repository struct {
	Branch            *string           `json:"branch,omitempty"`
	DeploymentLogsUrl *string           `json:"deploymentLogsUrl,omitempty"`
	DisplayUrl        *string           `json:"displayUrl,omitempty"`
	PathMapping       *[]ContentPathMap `json:"pathMapping,omitempty"`
	Url               *string           `json:"url,omitempty"`
}
