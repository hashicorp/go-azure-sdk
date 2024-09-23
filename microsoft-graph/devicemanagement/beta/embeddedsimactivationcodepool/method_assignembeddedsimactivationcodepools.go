package embeddedsimactivationcodepool

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

type AssignEmbeddedSIMActivationCodePoolsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.EmbeddedSIMActivationCodePoolAssignment
}

type AssignEmbeddedSIMActivationCodePoolsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.EmbeddedSIMActivationCodePoolAssignment
}

type AssignEmbeddedSIMActivationCodePoolsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultAssignEmbeddedSIMActivationCodePoolsOperationOptions() AssignEmbeddedSIMActivationCodePoolsOperationOptions {
	return AssignEmbeddedSIMActivationCodePoolsOperationOptions{}
}

func (o AssignEmbeddedSIMActivationCodePoolsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignEmbeddedSIMActivationCodePoolsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o AssignEmbeddedSIMActivationCodePoolsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AssignEmbeddedSIMActivationCodePoolsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignEmbeddedSIMActivationCodePoolsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignEmbeddedSIMActivationCodePools - Invoke action assign
func (c EmbeddedSIMActivationCodePoolClient) AssignEmbeddedSIMActivationCodePools(ctx context.Context, id beta.DeviceManagementEmbeddedSIMActivationCodePoolId, input AssignEmbeddedSIMActivationCodePoolsRequest, options AssignEmbeddedSIMActivationCodePoolsOperationOptions) (result AssignEmbeddedSIMActivationCodePoolsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AssignEmbeddedSIMActivationCodePoolsCustomPager{},
		Path:          fmt.Sprintf("%s/assign", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]beta.EmbeddedSIMActivationCodePoolAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignEmbeddedSIMActivationCodePoolsComplete retrieves all the results into a single object
func (c EmbeddedSIMActivationCodePoolClient) AssignEmbeddedSIMActivationCodePoolsComplete(ctx context.Context, id beta.DeviceManagementEmbeddedSIMActivationCodePoolId, input AssignEmbeddedSIMActivationCodePoolsRequest, options AssignEmbeddedSIMActivationCodePoolsOperationOptions) (AssignEmbeddedSIMActivationCodePoolsCompleteResult, error) {
	return c.AssignEmbeddedSIMActivationCodePoolsCompleteMatchingPredicate(ctx, id, input, options, EmbeddedSIMActivationCodePoolAssignmentOperationPredicate{})
}

// AssignEmbeddedSIMActivationCodePoolsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EmbeddedSIMActivationCodePoolClient) AssignEmbeddedSIMActivationCodePoolsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementEmbeddedSIMActivationCodePoolId, input AssignEmbeddedSIMActivationCodePoolsRequest, options AssignEmbeddedSIMActivationCodePoolsOperationOptions, predicate EmbeddedSIMActivationCodePoolAssignmentOperationPredicate) (result AssignEmbeddedSIMActivationCodePoolsCompleteResult, err error) {
	items := make([]beta.EmbeddedSIMActivationCodePoolAssignment, 0)

	resp, err := c.AssignEmbeddedSIMActivationCodePools(ctx, id, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = AssignEmbeddedSIMActivationCodePoolsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
