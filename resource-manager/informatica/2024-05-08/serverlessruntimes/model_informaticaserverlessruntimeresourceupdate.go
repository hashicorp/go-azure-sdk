package serverlessruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformaticaServerlessRuntimeResourceUpdate struct {
	Identity   *AzureResourceManagerCommonTypesManagedServiceIdentityUpdate `json:"identity,omitempty"`
	Properties *ServerlessRuntimePropertiesCustomUpdate                     `json:"properties,omitempty"`
}
