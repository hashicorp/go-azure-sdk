package calendargroupcalendareventinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendareventinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
