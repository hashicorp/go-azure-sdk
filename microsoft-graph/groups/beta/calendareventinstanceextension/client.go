package calendareventinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarEventInstanceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventInstanceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventInstanceExtensionClient: %+v", err)
	}

	return &CalendarEventInstanceExtensionClient{
		Client: client,
	}, nil
}
