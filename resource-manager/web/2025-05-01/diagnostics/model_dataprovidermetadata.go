package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataProviderMetadata struct {
	PropertyBag  *[]KeyValuePairStringObject `json:"propertyBag,omitempty"`
	ProviderName *string                     `json:"providerName,omitempty"`
}
