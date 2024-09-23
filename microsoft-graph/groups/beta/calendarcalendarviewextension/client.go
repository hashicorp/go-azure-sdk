package calendarcalendarviewextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewExtensionClient: %+v", err)
	}

	return &CalendarCalendarViewExtensionClient{
		Client: client,
	}, nil
}
