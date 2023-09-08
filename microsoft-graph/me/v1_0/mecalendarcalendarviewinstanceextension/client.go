package mecalendarcalendarviewinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewInstanceExtensionClient: %+v", err)
	}

	return &MeCalendarCalendarViewInstanceExtensionClient{
		Client: client,
	}, nil
}
