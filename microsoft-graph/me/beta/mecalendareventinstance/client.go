package mecalendareventinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventInstanceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventInstanceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventInstanceClient: %+v", err)
	}

	return &MeCalendarEventInstanceClient{
		Client: client,
	}, nil
}
