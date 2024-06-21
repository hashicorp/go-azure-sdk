package distributedavailabilitygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateInfo struct {
	CertificateId   *string `json:"certificateId,omitempty"`
	CertificateName *string `json:"certificateName,omitempty"`
	ExpiryDate      *string `json:"expiryDate,omitempty"`
}
