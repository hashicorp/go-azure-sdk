package mecalendarviewexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &MeCalendarViewExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
