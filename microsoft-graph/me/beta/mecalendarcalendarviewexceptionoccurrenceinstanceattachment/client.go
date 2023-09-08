package mecalendarcalendarviewexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
