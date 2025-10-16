package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationSecretParametersForPatch struct {
	AdminCredentials     *AdminCredentialsForPatch `json:"adminCredentials,omitempty"`
	SourceServerUsername *string                   `json:"sourceServerUsername,omitempty"`
	TargetServerUsername *string                   `json:"targetServerUsername,omitempty"`
}
