package appresiliency

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HeaderMatch struct {
	Header *string           `json:"header,omitempty"`
	Match  *HeaderMatchMatch `json:"match,omitempty"`
}
