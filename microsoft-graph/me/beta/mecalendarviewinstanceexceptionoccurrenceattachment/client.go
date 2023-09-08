package mecalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &MeCalendarViewInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
