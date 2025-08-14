package migrationservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MongoConnectionInformation struct {
	ConnectionString *string `json:"connectionString,omitempty"`
	Host             *string `json:"host,omitempty"`
	Password         *string `json:"password,omitempty"`
	Port             *int64  `json:"port,omitempty"`
	UseSsl           *bool   `json:"useSsl,omitempty"`
	UserName         *string `json:"userName,omitempty"`
}
