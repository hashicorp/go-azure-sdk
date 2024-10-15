package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnvVar struct {
	Name        *string `json:"name,omitempty"`
	SecretValue *string `json:"secretValue,omitempty"`
	Value       *string `json:"value,omitempty"`
}
