package usercalendargroupcalendarcalendarviewexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
