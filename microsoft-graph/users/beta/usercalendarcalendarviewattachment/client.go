package usercalendarcalendarviewattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewAttachmentClient: %+v", err)
	}

	return &UserCalendarCalendarViewAttachmentClient{
		Client: client,
	}, nil
}
