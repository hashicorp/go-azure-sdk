package clusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterUpdate struct {
	Properties *ClusterUpdateProperties `json:"properties,omitempty"`
	Sku        *Sku                     `json:"sku,omitempty"`
}
