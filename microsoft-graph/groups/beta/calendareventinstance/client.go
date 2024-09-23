package calendareventinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventInstanceClient struct {
	Client *msgraph.Client
}

func NewCalendarEventInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventInstanceClient: %+v", err)
	}

	return &CalendarEventInstanceClient{
		Client: client,
	}, nil
}
