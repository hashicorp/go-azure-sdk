package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatIntelligence struct {
	ArticleIndicators             *[]SecurityArticleIndicator             `json:"articleIndicators,omitempty"`
	Articles                      *[]SecurityArticle                      `json:"articles,omitempty"`
	HostComponents                *[]SecurityHostComponent                `json:"hostComponents,omitempty"`
	HostCookies                   *[]SecurityHostCookie                   `json:"hostCookies,omitempty"`
	HostPairs                     *[]SecurityHostPair                     `json:"hostPairs,omitempty"`
	HostPorts                     *[]SecurityHostPort                     `json:"hostPorts,omitempty"`
	HostSslCertificates           *[]SecurityHostSslCertificate           `json:"hostSslCertificates,omitempty"`
	HostTrackers                  *[]SecurityHostTracker                  `json:"hostTrackers,omitempty"`
	Hosts                         *[]SecurityHost                         `json:"hosts,omitempty"`
	Id                            *string                                 `json:"id,omitempty"`
	IntelProfiles                 *[]SecurityIntelligenceProfile          `json:"intelProfiles,omitempty"`
	IntelligenceProfileIndicators *[]SecurityIntelligenceProfileIndicator `json:"intelligenceProfileIndicators,omitempty"`
	ODataType                     *string                                 `json:"@odata.type,omitempty"`
	PassiveDnsRecords             *[]SecurityPassiveDnsRecord             `json:"passiveDnsRecords,omitempty"`
	SslCertificates               *[]SecuritySslCertificate               `json:"sslCertificates,omitempty"`
	Subdomains                    *[]SecuritySubdomain                    `json:"subdomains,omitempty"`
	Vulnerabilities               *[]SecurityVulnerability                `json:"vulnerabilities,omitempty"`
	WhoisHistoryRecords           *[]SecurityWhoisHistoryRecord           `json:"whoisHistoryRecords,omitempty"`
	WhoisRecords                  *[]SecurityWhoisRecord                  `json:"whoisRecords,omitempty"`
}
