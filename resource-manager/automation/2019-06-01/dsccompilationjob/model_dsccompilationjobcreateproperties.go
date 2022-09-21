package dsccompilationjob

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscCompilationJobCreateProperties struct {
	Configuration                   DscConfigurationAssociationProperty `json:"configuration"`
	IncrementNodeConfigurationBuild *bool                               `json:"incrementNodeConfigurationBuild,omitempty"`
	Parameters                      *map[string]string                  `json:"parameters,omitempty"`
}
