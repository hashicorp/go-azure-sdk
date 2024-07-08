package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerlessMetadataResponse struct {
	ServerlessConfigProperties        *ServerlessConfigProperties        `json:"serverlessConfigProperties,omitempty"`
	ServerlessRuntimeConfigProperties *ServerlessRuntimeConfigProperties `json:"serverlessRuntimeConfigProperties,omitempty"`
	Type                              *RuntimeType                       `json:"type,omitempty"`
}
