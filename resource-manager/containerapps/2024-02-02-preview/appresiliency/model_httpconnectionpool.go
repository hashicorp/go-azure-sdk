package appresiliency

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HTTPConnectionPool struct {
	HTTP1MaxPendingRequests *int64 `json:"http1MaxPendingRequests,omitempty"`
	HTTP2MaxRequests        *int64 `json:"http2MaxRequests,omitempty"`
}
