package componentannotationsapis

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *Annotation) GetEventTimeAsTime() (*time.Time, error) {
	if o.EventTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EventTime, "2006-01-02T15:04:05Z07:00")
}

func (o *Annotation) SetEventTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EventTime = &formatted
}
