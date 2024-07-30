package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileCardAnnotation struct {
	DisplayName   *string                    `json:"displayName,omitempty"`
	Localizations *[]DisplayNameLocalization `json:"localizations,omitempty"`
	ODataType     *string                    `json:"@odata.type,omitempty"`
}
