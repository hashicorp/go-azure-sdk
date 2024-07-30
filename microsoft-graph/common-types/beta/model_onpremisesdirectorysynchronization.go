package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesDirectorySynchronization struct {
	Configuration *OnPremisesDirectorySynchronizationConfiguration `json:"configuration,omitempty"`
	Features      *OnPremisesDirectorySynchronizationFeature       `json:"features,omitempty"`
	Id            *string                                          `json:"id,omitempty"`
	ODataType     *string                                          `json:"@odata.type,omitempty"`
}
