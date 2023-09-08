package mecalendargroupcalendarcalendarviewexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
