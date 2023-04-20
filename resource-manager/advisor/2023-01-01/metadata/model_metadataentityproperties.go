package metadata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetadataEntityProperties struct {
	ApplicableScenarios *[]Scenario                     `json:"applicableScenarios,omitempty"`
	DependsOn           *[]string                       `json:"dependsOn,omitempty"`
	DisplayName         *string                         `json:"displayName,omitempty"`
	SupportedValues     *[]MetadataSupportedValueDetail `json:"supportedValues,omitempty"`
}
