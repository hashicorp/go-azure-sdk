package mecalendareventinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &MeCalendarEventInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
