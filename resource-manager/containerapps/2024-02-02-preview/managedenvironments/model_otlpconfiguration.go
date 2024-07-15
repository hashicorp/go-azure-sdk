package managedenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OtlpConfiguration struct {
	Endpoint *string   `json:"endpoint,omitempty"`
	Headers  *[]Header `json:"headers,omitempty"`
	Insecure *bool     `json:"insecure,omitempty"`
	Name     *string   `json:"name,omitempty"`
}
