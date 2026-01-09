package calendareventextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarEventExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventExtensionClient: %+v", err)
	}

	return &CalendarEventExtensionClient{
		Client: client,
	}, nil
}
