package tasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MongoDbConnectionInfo struct {
	ConnectionString string  `json:"connectionString"`
	Password         *string `json:"password,omitempty"`
	Type             string  `json:"type"`
	UserName         *string `json:"userName,omitempty"`
}
