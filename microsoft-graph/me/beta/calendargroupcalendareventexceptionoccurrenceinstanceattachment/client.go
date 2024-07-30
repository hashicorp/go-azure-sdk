package calendargroupcalendareventexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendareventexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
