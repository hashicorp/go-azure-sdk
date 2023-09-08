package mecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarClient: %+v", err)
	}

	return &MeCalendarClient{
		Client: client,
	}, nil
}
