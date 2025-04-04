package operationalizationclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Image struct {
	Reference *string    `json:"reference,omitempty"`
	Type      *ImageType `json:"type,omitempty"`
	Version   *string    `json:"version,omitempty"`
}
