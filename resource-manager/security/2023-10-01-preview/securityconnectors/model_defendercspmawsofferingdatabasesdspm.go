package securityconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderCspmAwsOfferingDatabasesDspm struct {
	CloudRoleArn *string `json:"cloudRoleArn,omitempty"`
	Enabled      *bool   `json:"enabled,omitempty"`
}
