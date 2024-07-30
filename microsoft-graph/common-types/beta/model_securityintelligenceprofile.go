package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIntelligenceProfile struct {
	Aliases                    *[]string                                             `json:"aliases,omitempty"`
	CountriesOrRegionsOfOrigin *[]SecurityIntelligenceProfileCountryOrRegionOfOrigin `json:"countriesOrRegionsOfOrigin,omitempty"`
	Description                *FormattedContent                                     `json:"description,omitempty"`
	FirstActiveDateTime        *string                                               `json:"firstActiveDateTime,omitempty"`
	Id                         *string                                               `json:"id,omitempty"`
	Indicators                 *[]SecurityIntelligenceProfileIndicator               `json:"indicators,omitempty"`
	Kind                       *SecurityIntelligenceProfileKind                      `json:"kind,omitempty"`
	ODataType                  *string                                               `json:"@odata.type,omitempty"`
	Summary                    *FormattedContent                                     `json:"summary,omitempty"`
	Targets                    *[]string                                             `json:"targets,omitempty"`
	Title                      *string                                               `json:"title,omitempty"`
	Tradecraft                 *SecurityFormattedContent                             `json:"tradecraft,omitempty"`
}
