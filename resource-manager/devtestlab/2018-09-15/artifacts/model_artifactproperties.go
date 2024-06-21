package artifacts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArtifactProperties struct {
	CreatedDate  *string      `json:"createdDate,omitempty"`
	Description  *string      `json:"description,omitempty"`
	FilePath     *string      `json:"filePath,omitempty"`
	Icon         *string      `json:"icon,omitempty"`
	Parameters   *interface{} `json:"parameters,omitempty"`
	Publisher    *string      `json:"publisher,omitempty"`
	TargetOsType *string      `json:"targetOsType,omitempty"`
	Title        *string      `json:"title,omitempty"`
}
