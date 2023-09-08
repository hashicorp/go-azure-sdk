package groupcalendarcalendarviewattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewAttachmentClient: %+v", err)
	}

	return &GroupCalendarCalendarViewAttachmentClient{
		Client: client,
	}, nil
}
