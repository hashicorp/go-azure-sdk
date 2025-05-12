package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentParameter struct {
	Expression *string                     `json:"expression,omitempty"`
	Reference  *KeyVaultParameterReference `json:"reference,omitempty"`
	Value      *interface{}                `json:"value,omitempty"`
}
