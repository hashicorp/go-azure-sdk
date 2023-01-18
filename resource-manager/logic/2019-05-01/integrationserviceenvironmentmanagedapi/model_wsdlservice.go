package integrationserviceenvironmentmanagedapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WsdlService struct {
	EndpointQualifiedNames *[]string `json:"EndpointQualifiedNames,omitempty"`
	QualifiedName          *string   `json:"qualifiedName,omitempty"`
}
