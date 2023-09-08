package mecalendarviewinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &MeCalendarViewInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
