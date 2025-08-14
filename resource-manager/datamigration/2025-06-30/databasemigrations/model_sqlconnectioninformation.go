package databasemigrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlConnectionInformation struct {
	Authentication         *string `json:"authentication,omitempty"`
	DataSource             *string `json:"dataSource,omitempty"`
	EncryptConnection      *bool   `json:"encryptConnection,omitempty"`
	Password               *string `json:"password,omitempty"`
	TrustServerCertificate *bool   `json:"trustServerCertificate,omitempty"`
	UserName               *string `json:"userName,omitempty"`
}
