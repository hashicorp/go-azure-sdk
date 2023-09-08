package mecalendarcalendarviewinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &MeCalendarCalendarViewInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
