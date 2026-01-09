package calendareventattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarEventAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventAttachmentClient: %+v", err)
	}

	return &CalendarEventAttachmentClient{
		Client: client,
	}, nil
}
