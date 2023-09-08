package groupcalendareventinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventInstanceAttachmentClient: %+v", err)
	}

	return &GroupCalendarEventInstanceAttachmentClient{
		Client: client,
	}, nil
}
