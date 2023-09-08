package mecalendargroupcalendareventexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
