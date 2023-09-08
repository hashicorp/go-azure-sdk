package mecalendarcalendarviewexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &MeCalendarCalendarViewExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
