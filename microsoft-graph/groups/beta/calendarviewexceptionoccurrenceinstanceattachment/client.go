package calendarviewexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &CalendarViewExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
