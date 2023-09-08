package mecalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
