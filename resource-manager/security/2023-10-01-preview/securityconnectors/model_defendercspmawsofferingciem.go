package securityconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderCspmAwsOfferingCiem struct {
	CiemDiscovery *DefenderCspmAwsOfferingCiemCiemDiscovery `json:"ciemDiscovery,omitempty"`
	CiemOidc      *DefenderCspmAwsOfferingCiemCiemOidc      `json:"ciemOidc,omitempty"`
}
