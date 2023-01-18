package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiOperationAnnotation struct {
	Family   *string           `json:"family,omitempty"`
	Revision *int64            `json:"revision,omitempty"`
	Status   *StatusAnnotation `json:"status,omitempty"`
}
