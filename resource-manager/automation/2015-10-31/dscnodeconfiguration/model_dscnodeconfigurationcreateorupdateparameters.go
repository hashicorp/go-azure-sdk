package dscnodeconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscNodeConfigurationCreateOrUpdateParameters struct {
	Configuration                   DscConfigurationAssociationProperty `json:"configuration"`
	IncrementNodeConfigurationBuild *bool                               `json:"incrementNodeConfigurationBuild,omitempty"`
	Name                            string                              `json:"name"`
	Source                          ContentSource                       `json:"source"`
}
