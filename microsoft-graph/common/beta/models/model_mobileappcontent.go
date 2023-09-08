package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppContent struct {
	ContainedApps *[]MobileContainedApp   `json:"containedApps,omitempty"`
	Files         *[]MobileAppContentFile `json:"files,omitempty"`
	Id            *string                 `json:"id,omitempty"`
	ODataType     *string                 `json:"@odata.type,omitempty"`
}
