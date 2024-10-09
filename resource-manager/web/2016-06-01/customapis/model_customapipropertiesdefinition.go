package customapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomApiPropertiesDefinition struct {
	ApiDefinitions       *ApiResourceDefinitions         `json:"apiDefinitions,omitempty"`
	ApiType              *ApiType                        `json:"apiType,omitempty"`
	BackendService       *ApiResourceBackendService      `json:"backendService,omitempty"`
	BrandColor           *string                         `json:"brandColor,omitempty"`
	Capabilities         *[]string                       `json:"capabilities,omitempty"`
	ConnectionParameters *map[string]ConnectionParameter `json:"connectionParameters,omitempty"`
	Description          *string                         `json:"description,omitempty"`
	DisplayName          *string                         `json:"displayName,omitempty"`
	IconUri              *string                         `json:"iconUri,omitempty"`
	RuntimeURLs          *[]string                       `json:"runtimeUrls,omitempty"`
	Swagger              *interface{}                    `json:"swagger,omitempty"`
	WsdlDefinition       *WsdlDefinition                 `json:"wsdlDefinition,omitempty"`
}
