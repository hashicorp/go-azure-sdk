package mecalendarcalendarviewexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &MeCalendarCalendarViewExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
