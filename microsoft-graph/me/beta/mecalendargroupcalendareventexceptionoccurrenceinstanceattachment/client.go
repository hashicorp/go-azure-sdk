package mecalendargroupcalendareventexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
