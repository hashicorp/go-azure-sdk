package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAllDeviceCertificateState struct {
	CertificateExpirationDateTime             *string                                                  `json:"certificateExpirationDateTime,omitempty"`
	CertificateExtendedKeyUsages              *string                                                  `json:"certificateExtendedKeyUsages,omitempty"`
	CertificateIssuanceDateTime               *string                                                  `json:"certificateIssuanceDateTime,omitempty"`
	CertificateIssuerName                     *string                                                  `json:"certificateIssuerName,omitempty"`
	CertificateKeyUsages                      *int64                                                   `json:"certificateKeyUsages,omitempty"`
	CertificateRevokeStatus                   *ManagedAllDeviceCertificateStateCertificateRevokeStatus `json:"certificateRevokeStatus,omitempty"`
	CertificateRevokeStatusLastChangeDateTime *string                                                  `json:"certificateRevokeStatusLastChangeDateTime,omitempty"`
	CertificateSerialNumber                   *string                                                  `json:"certificateSerialNumber,omitempty"`
	CertificateSubjectName                    *string                                                  `json:"certificateSubjectName,omitempty"`
	CertificateThumbprint                     *string                                                  `json:"certificateThumbprint,omitempty"`
	Id                                        *string                                                  `json:"id,omitempty"`
	ManagedDeviceDisplayName                  *string                                                  `json:"managedDeviceDisplayName,omitempty"`
	ODataType                                 *string                                                  `json:"@odata.type,omitempty"`
	UserPrincipalName                         *string                                                  `json:"userPrincipalName,omitempty"`
}
