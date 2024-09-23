package embeddedsimactivationcodepoolassignment

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

type ListEmbeddedSIMActivationCodePoolAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.EmbeddedSIMActivationCodePoolAssignment
}

type ListEmbeddedSIMActivationCodePoolAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.EmbeddedSIMActivationCodePoolAssignment
}

type ListEmbeddedSIMActivationCodePoolAssignmentsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListEmbeddedSIMActivationCodePoolAssignmentsOperationOptions() ListEmbeddedSIMActivationCodePoolAssignmentsOperationOptions {
	return ListEmbeddedSIMActivationCodePoolAssignmentsOperationOptions{}
}

func (o ListEmbeddedSIMActivationCodePoolAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEmbeddedSIMActivationCodePoolAssignmentsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListEmbeddedSIMActivationCodePoolAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEmbeddedSIMActivationCodePoolAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEmbeddedSIMActivationCodePoolAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEmbeddedSIMActivationCodePoolAssignments - Get assignments from deviceManagement. Navigational property to a list
// of targets to which this pool is assigned.
func (c EmbeddedSIMActivationCodePoolAssignmentClient) ListEmbeddedSIMActivationCodePoolAssignments(ctx context.Context, id beta.DeviceManagementEmbeddedSIMActivationCodePoolId, options ListEmbeddedSIMActivationCodePoolAssignmentsOperationOptions) (result ListEmbeddedSIMActivationCodePoolAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEmbeddedSIMActivationCodePoolAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/assignments", id.ID()),
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

// ListEmbeddedSIMActivationCodePoolAssignmentsComplete retrieves all the results into a single object
func (c EmbeddedSIMActivationCodePoolAssignmentClient) ListEmbeddedSIMActivationCodePoolAssignmentsComplete(ctx context.Context, id beta.DeviceManagementEmbeddedSIMActivationCodePoolId, options ListEmbeddedSIMActivationCodePoolAssignmentsOperationOptions) (ListEmbeddedSIMActivationCodePoolAssignmentsCompleteResult, error) {
	return c.ListEmbeddedSIMActivationCodePoolAssignmentsCompleteMatchingPredicate(ctx, id, options, EmbeddedSIMActivationCodePoolAssignmentOperationPredicate{})
}

// ListEmbeddedSIMActivationCodePoolAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EmbeddedSIMActivationCodePoolAssignmentClient) ListEmbeddedSIMActivationCodePoolAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementEmbeddedSIMActivationCodePoolId, options ListEmbeddedSIMActivationCodePoolAssignmentsOperationOptions, predicate EmbeddedSIMActivationCodePoolAssignmentOperationPredicate) (result ListEmbeddedSIMActivationCodePoolAssignmentsCompleteResult, err error) {
	items := make([]beta.EmbeddedSIMActivationCodePoolAssignment, 0)

	resp, err := c.ListEmbeddedSIMActivationCodePoolAssignments(ctx, id, options)
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

	result = ListEmbeddedSIMActivationCodePoolAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
