package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppIcon struct {
	HostedContent *TeamworkHostedContent `json:"hostedContent,omitempty"`
	Id            *string                `json:"id,omitempty"`
	ODataType     *string                `json:"@odata.type,omitempty"`
	WebUrl        *string                `json:"webUrl,omitempty"`
}
