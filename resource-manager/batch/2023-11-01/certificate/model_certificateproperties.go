package certificate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateProperties struct {
	DeleteCertificateError                  *DeleteCertificateError       `json:"deleteCertificateError,omitempty"`
	Format                                  *CertificateFormat            `json:"format,omitempty"`
	PreviousProvisioningState               *CertificateProvisioningState `json:"previousProvisioningState,omitempty"`
	PreviousProvisioningStateTransitionTime *string                       `json:"previousProvisioningStateTransitionTime,omitempty"`
	ProvisioningState                       *CertificateProvisioningState `json:"provisioningState,omitempty"`
	ProvisioningStateTransitionTime         *string                       `json:"provisioningStateTransitionTime,omitempty"`
	PublicData                              *string                       `json:"publicData,omitempty"`
	Thumbprint                              *string                       `json:"thumbprint,omitempty"`
	ThumbprintAlgorithm                     *string                       `json:"thumbprintAlgorithm,omitempty"`
}
