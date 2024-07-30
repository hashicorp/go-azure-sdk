package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostPort struct {
	Banners                  *[]SecurityHostPortBanner    `json:"banners,omitempty"`
	FirstSeenDateTime        *string                      `json:"firstSeenDateTime,omitempty"`
	Host                     *SecurityHost                `json:"host,omitempty"`
	Id                       *string                      `json:"id,omitempty"`
	LastScanDateTime         *string                      `json:"lastScanDateTime,omitempty"`
	LastSeenDateTime         *string                      `json:"lastSeenDateTime,omitempty"`
	MostRecentSslCertificate *SecuritySslCertificate      `json:"mostRecentSslCertificate,omitempty"`
	ODataType                *string                      `json:"@odata.type,omitempty"`
	Port                     *int64                       `json:"port,omitempty"`
	Protocol                 *SecurityHostPortProtocol    `json:"protocol,omitempty"`
	Services                 *[]SecurityHostPortComponent `json:"services,omitempty"`
	Status                   *SecurityHostPortStatus      `json:"status,omitempty"`
	TimesObserved            *int64                       `json:"timesObserved,omitempty"`
}
