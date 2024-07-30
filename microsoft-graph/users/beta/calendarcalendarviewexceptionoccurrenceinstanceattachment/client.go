package calendarcalendarviewexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
