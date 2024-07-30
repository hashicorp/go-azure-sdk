package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEduCertificateSettings struct {
	CertFileName                   *string                                                  `json:"certFileName,omitempty"`
	CertificateTemplateName        *string                                                  `json:"certificateTemplateName,omitempty"`
	CertificateValidityPeriodScale *IosEduCertificateSettingsCertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
	CertificateValidityPeriodValue *int64                                                   `json:"certificateValidityPeriodValue,omitempty"`
	CertificationAuthority         *string                                                  `json:"certificationAuthority,omitempty"`
	CertificationAuthorityName     *string                                                  `json:"certificationAuthorityName,omitempty"`
	ODataType                      *string                                                  `json:"@odata.type,omitempty"`
	RenewalThresholdPercentage     *int64                                                   `json:"renewalThresholdPercentage,omitempty"`
	TrustedRootCertificate         *string                                                  `json:"trustedRootCertificate,omitempty"`
}
