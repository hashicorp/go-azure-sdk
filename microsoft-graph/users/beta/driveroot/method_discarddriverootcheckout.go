package driveroot

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscardDriveRootCheckoutOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DiscardDriveRootCheckoutOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultDiscardDriveRootCheckoutOperationOptions() DiscardDriveRootCheckoutOperationOptions {
	return DiscardDriveRootCheckoutOperationOptions{}
}

func (o DiscardDriveRootCheckoutOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DiscardDriveRootCheckoutOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DiscardDriveRootCheckoutOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DiscardDriveRootCheckout - Invoke action discardCheckout. Discard the check out of a driveItem. This action releases
// a driveItem resource that was previously checked out. Any changes made to the item while it was checked out are
// discarded. The same user that performed the checkout must discard it. Another alternative is to use application
// permissions.
func (c DriveRootClient) DiscardDriveRootCheckout(ctx context.Context, id beta.UserIdDriveId, options DiscardDriveRootCheckoutOperationOptions) (result DiscardDriveRootCheckoutOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/root/discardCheckout", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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
