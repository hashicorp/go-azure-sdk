package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildStageProperties struct {
	Name   *string                           `json:"name,omitempty"`
	Status *KPackBuildStageProvisioningState `json:"status,omitempty"`
}
