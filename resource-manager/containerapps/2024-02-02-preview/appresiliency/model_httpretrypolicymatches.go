package appresiliency

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HTTPRetryPolicyMatches struct {
	Errors          *[]string      `json:"errors,omitempty"`
	HTTPStatusCodes *[]int64       `json:"httpStatusCodes,omitempty"`
	Headers         *[]HeaderMatch `json:"headers,omitempty"`
}
