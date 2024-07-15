package containerappsbuilds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsBuildConfiguration struct {
	BaseOs               *string                `json:"baseOs,omitempty"`
	EnvironmentVariables *[]EnvironmentVariable `json:"environmentVariables,omitempty"`
	Platform             *string                `json:"platform,omitempty"`
	PlatformVersion      *string                `json:"platformVersion,omitempty"`
	PreBuildSteps        *[]PreBuildStep        `json:"preBuildSteps,omitempty"`
}
