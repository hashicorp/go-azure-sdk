package standbycontainergrouppoolruntimeviews

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerGroupInstanceCountSummary struct {
	InstanceCountsByState []PoolContainerGroupStateCount `json:"instanceCountsByState"`
	Zone                  *int64                         `json:"zone,omitempty"`
}
