package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkTeamsClientConfiguration struct {
	AccountConfiguration  *TeamworkAccountConfiguration  `json:"accountConfiguration,omitempty"`
	FeaturesConfiguration *TeamworkFeaturesConfiguration `json:"featuresConfiguration,omitempty"`
	ODataType             *string                        `json:"@odata.type,omitempty"`
}
