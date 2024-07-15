package appresiliency

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppResiliencyProperties struct {
	CircuitBreakerPolicy *CircuitBreakerPolicy `json:"circuitBreakerPolicy,omitempty"`
	HTTPConnectionPool   *HTTPConnectionPool   `json:"httpConnectionPool,omitempty"`
	HTTPRetryPolicy      *HTTPRetryPolicy      `json:"httpRetryPolicy,omitempty"`
	TcpConnectionPool    *TcpConnectionPool    `json:"tcpConnectionPool,omitempty"`
	TcpRetryPolicy       *TcpRetryPolicy       `json:"tcpRetryPolicy,omitempty"`
	TimeoutPolicy        *TimeoutPolicy        `json:"timeoutPolicy,omitempty"`
}
