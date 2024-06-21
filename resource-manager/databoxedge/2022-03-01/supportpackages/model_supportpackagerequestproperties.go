package supportpackages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportPackageRequestProperties struct {
	Include          *string `json:"include,omitempty"`
	MaximumTimeStamp *string `json:"maximumTimeStamp,omitempty"`
	MinimumTimeStamp *string `json:"minimumTimeStamp,omitempty"`
}
