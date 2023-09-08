package mecalendargroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupClient: %+v", err)
	}

	return &MeCalendarGroupClient{
		Client: client,
	}, nil
}
