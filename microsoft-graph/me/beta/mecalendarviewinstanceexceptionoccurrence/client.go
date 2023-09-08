package mecalendarviewinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &MeCalendarViewInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
