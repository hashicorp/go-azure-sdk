package driveitemlistitemversion

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

type GetDriveItemListItemVersionsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetDriveItemListItemVersionsCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetDriveItemListItemVersionsCountOperationOptions() GetDriveItemListItemVersionsCountOperationOptions {
	return GetDriveItemListItemVersionsCountOperationOptions{}
}

func (o GetDriveItemListItemVersionsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveItemListItemVersionsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetDriveItemListItemVersionsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveItemListItemVersionsCount - Get the number of the resource
func (c DriveItemListItemVersionClient) GetDriveItemListItemVersionsCount(ctx context.Context, id beta.GroupIdDriveIdItemId, options GetDriveItemListItemVersionsCountOperationOptions) (result GetDriveItemListItemVersionsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/listItem/versions/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
