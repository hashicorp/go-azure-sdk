package mecalendarview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewClient: %+v", err)
	}

	return &MeCalendarViewClient{
		Client: client,
	}, nil
}
