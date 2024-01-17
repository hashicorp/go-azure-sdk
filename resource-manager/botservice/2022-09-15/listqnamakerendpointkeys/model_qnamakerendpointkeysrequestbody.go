package listqnamakerendpointkeys

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QnAMakerEndpointKeysRequestBody struct {
	Authkey  *string `json:"authkey,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
}
