package mecalendareventexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &MeCalendarEventExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
