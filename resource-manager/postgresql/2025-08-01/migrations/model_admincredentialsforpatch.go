package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminCredentialsForPatch struct {
	SourceServerPassword *string `json:"sourceServerPassword,omitempty"`
	TargetServerPassword *string `json:"targetServerPassword,omitempty"`
}
