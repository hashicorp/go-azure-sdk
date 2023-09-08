package mecalendareventinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventInstanceExtensionClient: %+v", err)
	}

	return &MeCalendarEventInstanceExtensionClient{
		Client: client,
	}, nil
}
