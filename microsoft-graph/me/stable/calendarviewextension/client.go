package calendarviewextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarViewExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewExtensionClient: %+v", err)
	}

	return &CalendarViewExtensionClient{
		Client: client,
	}, nil
}
