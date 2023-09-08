package mecalendareventattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventAttachmentClient: %+v", err)
	}

	return &MeCalendarEventAttachmentClient{
		Client: client,
	}, nil
}
