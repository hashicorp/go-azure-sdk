package dataproducts

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountSas struct {
	ExpiryTimeStamp string `json:"expiryTimeStamp"`
	IPAddress       string `json:"ipAddress"`
	StartTimeStamp  string `json:"startTimeStamp"`
}

func (o *AccountSas) GetExpiryTimeStampAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.ExpiryTimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *AccountSas) SetExpiryTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryTimeStamp = formatted
}

func (o *AccountSas) GetStartTimeStampAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *AccountSas) SetStartTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTimeStamp = formatted
}
