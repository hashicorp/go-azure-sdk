package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GCPAuthProperties struct {
	ProjectNumber              string `json:"projectNumber"`
	ServiceAccountEmail        string `json:"serviceAccountEmail"`
	WorkloadIdentityProviderId string `json:"workloadIdentityProviderId"`
}
