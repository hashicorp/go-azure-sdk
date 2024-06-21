package componentannotationsapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Annotation struct {
	AnnotationName    *string `json:"AnnotationName,omitempty"`
	Category          *string `json:"Category,omitempty"`
	EventTime         *string `json:"EventTime,omitempty"`
	Id                *string `json:"Id,omitempty"`
	Properties        *string `json:"Properties,omitempty"`
	RelatedAnnotation *string `json:"RelatedAnnotation,omitempty"`
}
