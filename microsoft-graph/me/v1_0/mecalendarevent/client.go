package mecalendarevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventClient: %+v", err)
	}

	return &MeCalendarEventClient{
		Client: client,
	}, nil
}
