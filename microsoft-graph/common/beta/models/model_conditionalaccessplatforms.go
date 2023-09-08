package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPlatforms struct {
	ExcludePlatforms *[]ConditionalAccessPlatformsExcludePlatforms `json:"excludePlatforms,omitempty"`
	IncludePlatforms *[]ConditionalAccessPlatformsIncludePlatforms `json:"includePlatforms,omitempty"`
	ODataType        *string                                       `json:"@odata.type,omitempty"`
}
