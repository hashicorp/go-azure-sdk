package integrationruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIntegrationRuntimeError struct {
	Code       *string   `json:"code,omitempty"`
	Message    *string   `json:"message,omitempty"`
	Parameters *[]string `json:"parameters,omitempty"`
	Time       *string   `json:"time,omitempty"`
}
