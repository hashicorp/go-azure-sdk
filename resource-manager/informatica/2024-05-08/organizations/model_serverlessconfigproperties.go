package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerlessConfigProperties struct {
	ApplicationTypes *[]ApplicationTypeMetadata `json:"applicationTypes,omitempty"`
	ComputeUnits     *[]ComputeUnitsMetadata    `json:"computeUnits,omitempty"`
	ExecutionTimeout *string                    `json:"executionTimeout,omitempty"`
	Platform         *PlatformType              `json:"platform,omitempty"`
	Regions          *[]RegionsMetadata         `json:"regions,omitempty"`
}
