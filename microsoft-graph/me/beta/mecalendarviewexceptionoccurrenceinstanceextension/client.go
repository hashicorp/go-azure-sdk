package mecalendarviewexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &MeCalendarViewExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
