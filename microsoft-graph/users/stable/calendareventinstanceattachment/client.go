package calendareventinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarEventInstanceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventInstanceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventInstanceAttachmentClient: %+v", err)
	}

	return &CalendarEventInstanceAttachmentClient{
		Client: client,
	}, nil
}
