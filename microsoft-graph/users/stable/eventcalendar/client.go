package eventcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventCalendarClient struct {
	Client *msgraph.Client
}

func NewEventCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*EventCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventCalendarClient: %+v", err)
	}

	return &EventCalendarClient{
		Client: client,
	}, nil
}
