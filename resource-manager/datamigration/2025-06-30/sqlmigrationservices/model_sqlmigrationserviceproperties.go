package sqlmigrationservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlMigrationServiceProperties struct {
	IntegrationRuntimeState *string `json:"integrationRuntimeState,omitempty"`
	ProvisioningState       *string `json:"provisioningState,omitempty"`
}
