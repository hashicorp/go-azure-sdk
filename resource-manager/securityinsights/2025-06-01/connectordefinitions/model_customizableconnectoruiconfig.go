package connectordefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomizableConnectorUiConfig struct {
	Availability                     *ConnectorDefinitionsAvailability `json:"availability,omitempty"`
	ConnectivityCriteria             []ConnectivityCriterion           `json:"connectivityCriteria"`
	DataTypes                        []ConnectorDataType               `json:"dataTypes"`
	DescriptionMarkdown              string                            `json:"descriptionMarkdown"`
	GraphQueries                     []GraphQuery                      `json:"graphQueries"`
	Id                               *string                           `json:"id,omitempty"`
	InstructionSteps                 []InstructionStep                 `json:"instructionSteps"`
	IsConnectivityCriteriasMatchSome *bool                             `json:"isConnectivityCriteriasMatchSome,omitempty"`
	Logo                             *string                           `json:"logo,omitempty"`
	Permissions                      ConnectorDefinitionsPermissions   `json:"permissions"`
	Publisher                        string                            `json:"publisher"`
	Title                            string                            `json:"title"`
}
