package drivelistcontenttype

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

type CreateDriveListContentTypeUnpublishOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateDriveListContentTypeUnpublishOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDriveListContentTypeUnpublishOperationOptions() CreateDriveListContentTypeUnpublishOperationOptions {
	return CreateDriveListContentTypeUnpublishOperationOptions{}
}

func (o CreateDriveListContentTypeUnpublishOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDriveListContentTypeUnpublishOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDriveListContentTypeUnpublishOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDriveListContentTypeUnpublish - Invoke action unpublish
func (c DriveListContentTypeClient) CreateDriveListContentTypeUnpublish(ctx context.Context, id beta.GroupIdDriveIdListContentTypeId, options CreateDriveListContentTypeUnpublishOperationOptions) (result CreateDriveListContentTypeUnpublishOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/unpublish", id.ID()),
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
