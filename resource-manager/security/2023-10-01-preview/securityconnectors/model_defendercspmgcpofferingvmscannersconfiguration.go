package securityconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderCspmGcpOfferingVMScannersConfiguration struct {
	ExclusionTags *map[string]string `json:"exclusionTags,omitempty"`
	ScanningMode  *ScanningMode      `json:"scanningMode,omitempty"`
}
