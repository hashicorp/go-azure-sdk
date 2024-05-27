package dataproductscatalogs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataProductsCatalogProperties struct {
	ProvisioningState *ProvisioningState     `json:"provisioningState,omitempty"`
	Publishers        []PublisherInformation `json:"publishers"`
}
