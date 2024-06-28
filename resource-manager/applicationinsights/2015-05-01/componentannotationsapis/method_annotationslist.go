package componentannotationsapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnnotationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *AnnotationsListResult
}

type AnnotationsListOperationOptions struct {
	End   *string
	Start *string
}

func DefaultAnnotationsListOperationOptions() AnnotationsListOperationOptions {
	return AnnotationsListOperationOptions{}
}

func (o AnnotationsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AnnotationsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o AnnotationsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.End != nil {
		out.Append("end", fmt.Sprintf("%v", *o.End))
	}
	if o.Start != nil {
		out.Append("start", fmt.Sprintf("%v", *o.Start))
	}
	return &out
}

// AnnotationsList ...
func (c ComponentAnnotationsAPIsClient) AnnotationsList(ctx context.Context, id ComponentId, options AnnotationsListOperationOptions) (result AnnotationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/annotations", id.ID()),
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

	var model AnnotationsListResult
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
