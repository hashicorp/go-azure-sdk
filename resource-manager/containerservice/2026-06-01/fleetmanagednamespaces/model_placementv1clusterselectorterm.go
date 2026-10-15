package fleetmanagednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementV1ClusterSelectorTerm struct {
	LabelSelector    *MetaV1LabelSelector         `json:"labelSelector,omitempty"`
	PropertySelector *PlacementV1PropertySelector `json:"propertySelector,omitempty"`
}
