package mecalendareventexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &MeCalendarEventExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
