package mecalendareventextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventExtensionClient: %+v", err)
	}

	return &MeCalendarEventExtensionClient{
		Client: client,
	}, nil
}
