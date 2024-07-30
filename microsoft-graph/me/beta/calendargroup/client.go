package calendargroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupClient: %+v", err)
	}

	return &CalendarGroupClient{
		Client: client,
	}, nil
}
