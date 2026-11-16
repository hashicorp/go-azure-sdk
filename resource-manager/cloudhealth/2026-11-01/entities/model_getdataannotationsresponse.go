package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDataAnnotationsResponse struct {
	Annotations []DataAnnotation `json:"annotations"`
	EntityName  string           `json:"entityName"`
	NextMarker  *string          `json:"nextMarker,omitempty"`
}
