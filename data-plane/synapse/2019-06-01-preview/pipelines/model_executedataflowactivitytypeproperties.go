package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExecuteDataFlowActivityTypeProperties struct {
	Compute            *ExecuteDataFlowActivityTypePropertiesCompute `json:"compute,omitempty"`
	ContinueOnError    *interface{}                                  `json:"continueOnError,omitempty"`
	Dataflow           DataFlowReference                             `json:"dataflow"`
	IntegrationRuntime *IntegrationRuntimeReference                  `json:"integrationRuntime,omitempty"`
	RunConcurrently    *interface{}                                  `json:"runConcurrently,omitempty"`
	Staging            *DataFlowStagingInfo                          `json:"staging,omitempty"`
	TraceLevel         *interface{}                                  `json:"traceLevel,omitempty"`
}
