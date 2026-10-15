package fleetupdatestrategies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateGroup struct {
	AfterGates     *[]GateConfiguration `json:"afterGates,omitempty"`
	BeforeGates    *[]GateConfiguration `json:"beforeGates,omitempty"`
	MaxConcurrency *string              `json:"maxConcurrency,omitempty"`
	Name           string               `json:"name"`
}
