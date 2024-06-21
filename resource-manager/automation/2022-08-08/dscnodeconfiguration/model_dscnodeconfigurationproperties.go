package dscnodeconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscNodeConfigurationProperties struct {
	Configuration                   *DscConfigurationAssociationProperty `json:"configuration,omitempty"`
	CreationTime                    *string                              `json:"creationTime,omitempty"`
	IncrementNodeConfigurationBuild *bool                                `json:"incrementNodeConfigurationBuild,omitempty"`
	LastModifiedTime                *string                              `json:"lastModifiedTime,omitempty"`
	NodeCount                       *int64                               `json:"nodeCount,omitempty"`
	Source                          *string                              `json:"source,omitempty"`
}
