package inferencepool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartialMinimalTrackedResourceWithSkuAndIdentity struct {
	Identity *PartialManagedServiceIdentity `json:"identity,omitempty"`
	Sku      *PartialSku                    `json:"sku,omitempty"`
	Tags     *map[string]string             `json:"tags,omitempty"`
}
