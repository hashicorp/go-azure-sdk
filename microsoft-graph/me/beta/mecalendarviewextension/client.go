package mecalendarviewextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewExtensionClient: %+v", err)
	}

	return &MeCalendarViewExtensionClient{
		Client: client,
	}, nil
}
