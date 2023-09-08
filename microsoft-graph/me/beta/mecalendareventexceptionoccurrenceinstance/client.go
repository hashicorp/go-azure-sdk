package mecalendareventexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &MeCalendarEventExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
