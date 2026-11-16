package logfiles

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogFileProperties struct {
	CreatedAtUtc          string      `json:"createdAtUtc"`
	IsMutable             bool        `json:"isMutable"`
	LastModifiedAtUtc     string      `json:"lastModifiedAtUtc"`
	LogType               LogFileType `json:"logType"`
	Node                  string      `json:"node"`
	Pool                  string      `json:"pool"`
	RetentionExpiresAtUtc string      `json:"retentionExpiresAtUtc"`
	Role                  NodeRole    `json:"role"`
	SizeInKilobytes       int64       `json:"sizeInKilobytes"`
}

func (o *LogFileProperties) GetCreatedAtUtcAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.CreatedAtUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *LogFileProperties) SetCreatedAtUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedAtUtc = formatted
}

func (o *LogFileProperties) GetLastModifiedAtUtcAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.LastModifiedAtUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *LogFileProperties) SetLastModifiedAtUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedAtUtc = formatted
}

func (o *LogFileProperties) GetRetentionExpiresAtUtcAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.RetentionExpiresAtUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *LogFileProperties) SetRetentionExpiresAtUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RetentionExpiresAtUtc = formatted
}
