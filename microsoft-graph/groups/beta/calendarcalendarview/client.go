package calendarcalendarview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewClient: %+v", err)
	}

	return &CalendarCalendarViewClient{
		Client: client,
	}, nil
}
