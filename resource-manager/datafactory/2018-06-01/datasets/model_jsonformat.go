package datasets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JsonFormat struct {
	Deserializer       *interface{} `json:"deserializer,omitempty"`
	EncodingName       *interface{} `json:"encodingName,omitempty"`
	FilePattern        *interface{} `json:"filePattern,omitempty"`
	JsonNodeReference  *interface{} `json:"jsonNodeReference,omitempty"`
	JsonPathDefinition *interface{} `json:"jsonPathDefinition,omitempty"`
	NestingSeparator   *interface{} `json:"nestingSeparator,omitempty"`
	Serializer         *interface{} `json:"serializer,omitempty"`
	Type               string       `json:"type"`
}
