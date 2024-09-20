package mobileapptroubleshootingeventapplogcollectionrequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateMobileAppTroubleshootingEventAppLogCollectionRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateMobileAppTroubleshootingEventAppLogCollectionRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateMobileAppTroubleshootingEventAppLogCollectionRequestOperationOptions() UpdateMobileAppTroubleshootingEventAppLogCollectionRequestOperationOptions {
	return UpdateMobileAppTroubleshootingEventAppLogCollectionRequestOperationOptions{}
}

func (o UpdateMobileAppTroubleshootingEventAppLogCollectionRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateMobileAppTroubleshootingEventAppLogCollectionRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateMobileAppTroubleshootingEventAppLogCollectionRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateMobileAppTroubleshootingEventAppLogCollectionRequest - Update appLogCollectionRequest. Update the properties of
// a appLogCollectionRequest object.
func (c MobileAppTroubleshootingEventAppLogCollectionRequestClient) UpdateMobileAppTroubleshootingEventAppLogCollectionRequest(ctx context.Context, id stable.DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId, input stable.AppLogCollectionRequest, options UpdateMobileAppTroubleshootingEventAppLogCollectionRequestOperationOptions) (result UpdateMobileAppTroubleshootingEventAppLogCollectionRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
