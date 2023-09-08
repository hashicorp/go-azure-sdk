package mecalendarcalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
