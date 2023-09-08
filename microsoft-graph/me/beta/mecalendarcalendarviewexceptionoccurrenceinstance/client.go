package mecalendarcalendarviewexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &MeCalendarCalendarViewExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
