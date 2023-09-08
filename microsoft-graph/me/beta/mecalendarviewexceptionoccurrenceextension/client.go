package mecalendarviewexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &MeCalendarViewExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
