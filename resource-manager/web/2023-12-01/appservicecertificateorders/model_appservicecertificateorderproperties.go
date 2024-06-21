package appservicecertificateorders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppServiceCertificateOrderProperties struct {
	AppServiceCertificateNotRenewableReasons *[]ResourceNotRenewableReason     `json:"appServiceCertificateNotRenewableReasons,omitempty"`
	AutoRenew                                *bool                             `json:"autoRenew,omitempty"`
	Certificates                             *map[string]AppServiceCertificate `json:"certificates,omitempty"`
	Contact                                  *CertificateOrderContact          `json:"contact,omitempty"`
	Csr                                      *string                           `json:"csr,omitempty"`
	DistinguishedName                        *string                           `json:"distinguishedName,omitempty"`
	DomainVerificationToken                  *string                           `json:"domainVerificationToken,omitempty"`
	ExpirationTime                           *string                           `json:"expirationTime,omitempty"`
	Intermediate                             *CertificateDetails               `json:"intermediate,omitempty"`
	IsPrivateKeyExternal                     *bool                             `json:"isPrivateKeyExternal,omitempty"`
	KeySize                                  *int64                            `json:"keySize,omitempty"`
	LastCertificateIssuanceTime              *string                           `json:"lastCertificateIssuanceTime,omitempty"`
	NextAutoRenewalTimeStamp                 *string                           `json:"nextAutoRenewalTimeStamp,omitempty"`
	ProductType                              CertificateProductType            `json:"productType"`
	ProvisioningState                        *ProvisioningState                `json:"provisioningState,omitempty"`
	Root                                     *CertificateDetails               `json:"root,omitempty"`
	SerialNumber                             *string                           `json:"serialNumber,omitempty"`
	SignedCertificate                        *CertificateDetails               `json:"signedCertificate,omitempty"`
	Status                                   *CertificateOrderStatus           `json:"status,omitempty"`
	ValidityInYears                          *int64                            `json:"validityInYears,omitempty"`
}
