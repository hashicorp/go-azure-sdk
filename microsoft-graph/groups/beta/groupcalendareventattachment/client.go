package groupcalendareventattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventAttachmentClient: %+v", err)
	}

	return &GroupCalendarEventAttachmentClient{
		Client: client,
	}, nil
}
