package integrationserviceenvironmentmanagedapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiDeploymentParameterMetadataSet struct {
	PackageContentLink         *ApiDeploymentParameterMetadata `json:"packageContentLink,omitempty"`
	RedisCacheConnectionString *ApiDeploymentParameterMetadata `json:"redisCacheConnectionString,omitempty"`
}
