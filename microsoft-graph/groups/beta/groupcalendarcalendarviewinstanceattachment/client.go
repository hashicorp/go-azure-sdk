package groupcalendarcalendarviewinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewInstanceAttachmentClient: %+v", err)
	}

	return &GroupCalendarCalendarViewInstanceAttachmentClient{
		Client: client,
	}, nil
}
