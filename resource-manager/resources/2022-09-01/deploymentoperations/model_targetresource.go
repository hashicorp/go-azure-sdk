package deploymentoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetResource struct {
	Id           *string `json:"id,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}
