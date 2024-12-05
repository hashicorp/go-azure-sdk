package endpoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerlessEndpointInferenceEndpoint struct {
	Headers *map[string]string `json:"headers,omitempty"`
	Uri     string             `json:"uri"`
}
