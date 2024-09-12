package calendarviewinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewInstanceClient struct {
	Client *msgraph.Client
}

func NewCalendarViewInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewInstanceClient: %+v", err)
	}

	return &CalendarViewInstanceClient{
		Client: client,
	}, nil
}
