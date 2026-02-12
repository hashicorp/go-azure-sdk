package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowDebugPackage struct {
	DataFlow       *DataFlowDebugResource             `json:"dataFlow,omitempty"`
	DataFlows      *[]DataFlowDebugResource           `json:"dataFlows,omitempty"`
	Datasets       *[]DatasetDebugResource            `json:"datasets,omitempty"`
	DebugSettings  *DataFlowDebugPackageDebugSettings `json:"debugSettings,omitempty"`
	LinkedServices *[]LinkedServiceDebugResource      `json:"linkedServices,omitempty"`
	SessionId      *string                            `json:"sessionId,omitempty"`
	Staging        *DataFlowStagingInfo               `json:"staging,omitempty"`
}
