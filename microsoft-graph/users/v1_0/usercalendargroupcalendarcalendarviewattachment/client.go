package usercalendargroupcalendarcalendarviewattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewAttachmentClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewAttachmentClient{
		Client: client,
	}, nil
}
