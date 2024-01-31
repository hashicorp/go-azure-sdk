package tasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectToTargetAzureDbForMySqlTaskInput struct {
	SourceConnectionInfo MySqlConnectionInfo `json:"sourceConnectionInfo"`
	TargetConnectionInfo MySqlConnectionInfo `json:"targetConnectionInfo"`
}
