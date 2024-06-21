package dscnodeconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscNodeConfiguration struct {
	Configuration    *DscConfigurationAssociationProperty `json:"configuration,omitempty"`
	CreationTime     *string                              `json:"creationTime,omitempty"`
	Id               *string                              `json:"id,omitempty"`
	LastModifiedTime *string                              `json:"lastModifiedTime,omitempty"`
	Name             *string                              `json:"name,omitempty"`
	Type             *string                              `json:"type,omitempty"`
}
