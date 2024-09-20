package driveitemversioncontent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetDriveItemVersionContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDriveItemVersionContentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetDriveItemVersionContentOperationOptions() SetDriveItemVersionContentOperationOptions {
	return SetDriveItemVersionContentOperationOptions{}
}

func (o SetDriveItemVersionContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDriveItemVersionContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDriveItemVersionContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDriveItemVersionContent - Update content for the navigation property versions in users. The content stream for
// this version of the item.
func (c DriveItemVersionContentClient) SetDriveItemVersionContent(ctx context.Context, id stable.UserIdDriveIdItemIdVersionId, input []byte, options SetDriveItemVersionContentOperationOptions) (result SetDriveItemVersionContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/content", id.ID()),
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
