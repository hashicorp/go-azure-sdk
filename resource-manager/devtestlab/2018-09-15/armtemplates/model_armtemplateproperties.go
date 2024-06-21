package armtemplates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArmTemplateProperties struct {
	Contents                 *interface{}               `json:"contents,omitempty"`
	CreatedDate              *string                    `json:"createdDate,omitempty"`
	Description              *string                    `json:"description,omitempty"`
	DisplayName              *string                    `json:"displayName,omitempty"`
	Enabled                  *bool                      `json:"enabled,omitempty"`
	Icon                     *string                    `json:"icon,omitempty"`
	ParametersValueFilesInfo *[]ParametersValueFileInfo `json:"parametersValueFilesInfo,omitempty"`
	Publisher                *string                    `json:"publisher,omitempty"`
}
