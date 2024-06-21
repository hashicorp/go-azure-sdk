package connectionmonitors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionStateSnapshot struct {
	AvgLatencyInMs  *int64             `json:"avgLatencyInMs,omitempty"`
	ConnectionState *ConnectionState   `json:"connectionState,omitempty"`
	EndTime         *string            `json:"endTime,omitempty"`
	EvaluationState *EvaluationState   `json:"evaluationState,omitempty"`
	Hops            *[]ConnectivityHop `json:"hops,omitempty"`
	MaxLatencyInMs  *int64             `json:"maxLatencyInMs,omitempty"`
	MinLatencyInMs  *int64             `json:"minLatencyInMs,omitempty"`
	ProbesFailed    *int64             `json:"probesFailed,omitempty"`
	ProbesSent      *int64             `json:"probesSent,omitempty"`
	StartTime       *string            `json:"startTime,omitempty"`
}
