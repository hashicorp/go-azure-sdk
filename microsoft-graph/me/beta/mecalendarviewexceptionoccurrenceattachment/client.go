package mecalendarviewexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &MeCalendarViewExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
