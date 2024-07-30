package calendarcalendarviewinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewInstanceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewInstanceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewInstanceExtensionClient: %+v", err)
	}

	return &CalendarCalendarViewInstanceExtensionClient{
		Client: client,
	}, nil
}
