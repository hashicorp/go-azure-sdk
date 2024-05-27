package dataproducts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsumptionEndpointsProperties struct {
	FileAccessResourceId *string `json:"fileAccessResourceId,omitempty"`
	FileAccessUrl        *string `json:"fileAccessUrl,omitempty"`
	IngestionResourceId  *string `json:"ingestionResourceId,omitempty"`
	IngestionUrl         *string `json:"ingestionUrl,omitempty"`
	QueryResourceId      *string `json:"queryResourceId,omitempty"`
	QueryUrl             *string `json:"queryUrl,omitempty"`
}
