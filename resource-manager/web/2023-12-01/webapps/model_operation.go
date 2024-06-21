package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Operation struct {
	CreatedTime          *string          `json:"createdTime,omitempty"`
	Errors               *[]ErrorEntity   `json:"errors,omitempty"`
	ExpirationTime       *string          `json:"expirationTime,omitempty"`
	GeoMasterOperationId *string          `json:"geoMasterOperationId,omitempty"`
	Id                   *string          `json:"id,omitempty"`
	ModifiedTime         *string          `json:"modifiedTime,omitempty"`
	Name                 *string          `json:"name,omitempty"`
	Status               *OperationStatus `json:"status,omitempty"`
}
