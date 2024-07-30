package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationCsvDataProvider struct {
	Customizations *EducationSynchronizationCustomizations `json:"customizations,omitempty"`
	ODataType      *string                                 `json:"@odata.type,omitempty"`
}
