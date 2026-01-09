package calendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarClient: %+v", err)
	}

	return &CalendarClient{
		Client: client,
	}, nil
}
