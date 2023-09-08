package mecalendarviewexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewExceptionOccurrenceClient: %+v", err)
	}

	return &MeCalendarViewExceptionOccurrenceClient{
		Client: client,
	}, nil
}
