package databasemigrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupConfiguration struct {
	SourceLocation *SourceLocation `json:"sourceLocation,omitempty"`
	TargetLocation *TargetLocation `json:"targetLocation,omitempty"`
}
