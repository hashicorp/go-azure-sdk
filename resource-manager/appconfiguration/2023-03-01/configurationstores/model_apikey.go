package configurationstores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiKey struct {
	ConnectionString *string `json:"connectionString,omitempty"`
	Id               *string `json:"id,omitempty"`
	LastModified     *string `json:"lastModified,omitempty"`
	Name             *string `json:"name,omitempty"`
	ReadOnly         *bool   `json:"readOnly,omitempty"`
	Value            *string `json:"value,omitempty"`
}
