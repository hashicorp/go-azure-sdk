package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Repository struct {
	Branch            *string           `json:"branch,omitempty"`
	DeploymentLogsURL *string           `json:"deploymentLogsUrl,omitempty"`
	DisplayURL        *string           `json:"displayUrl,omitempty"`
	PathMapping       *[]ContentPathMap `json:"pathMapping,omitempty"`
	Url               *string           `json:"url,omitempty"`
}
