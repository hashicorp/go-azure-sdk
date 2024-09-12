package calendargroupcalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarviewinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &CalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
