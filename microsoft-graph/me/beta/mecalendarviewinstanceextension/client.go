package mecalendarviewinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewInstanceExtensionClient: %+v", err)
	}

	return &MeCalendarViewInstanceExtensionClient{
		Client: client,
	}, nil
}
