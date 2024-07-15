package appresiliency

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TcpConnectionPool struct {
	MaxConnections *int64 `json:"maxConnections,omitempty"`
}
