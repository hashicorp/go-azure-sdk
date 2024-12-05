package serverlessendpoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerlessInferenceEndpoint struct {
	Headers *map[string]string `json:"headers,omitempty"`
	Uri     string             `json:"uri"`
}
