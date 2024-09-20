package resourceaccessprofileassignment

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

type ListResourceAccessProfileAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementResourceAccessProfileAssignment
}

type ListResourceAccessProfileAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementResourceAccessProfileAssignment
}

type ListResourceAccessProfileAssignmentsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListResourceAccessProfileAssignmentsOperationOptions() ListResourceAccessProfileAssignmentsOperationOptions {
	return ListResourceAccessProfileAssignmentsOperationOptions{}
}

func (o ListResourceAccessProfileAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListResourceAccessProfileAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListResourceAccessProfileAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListResourceAccessProfileAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListResourceAccessProfileAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListResourceAccessProfileAssignments - Get assignments from deviceManagement. The list of assignments for the device
// configuration profile.
func (c ResourceAccessProfileAssignmentClient) ListResourceAccessProfileAssignments(ctx context.Context, id beta.DeviceManagementResourceAccessProfileId, options ListResourceAccessProfileAssignmentsOperationOptions) (result ListResourceAccessProfileAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListResourceAccessProfileAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.DeviceManagementResourceAccessProfileAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListResourceAccessProfileAssignmentsComplete retrieves all the results into a single object
func (c ResourceAccessProfileAssignmentClient) ListResourceAccessProfileAssignmentsComplete(ctx context.Context, id beta.DeviceManagementResourceAccessProfileId, options ListResourceAccessProfileAssignmentsOperationOptions) (ListResourceAccessProfileAssignmentsCompleteResult, error) {
	return c.ListResourceAccessProfileAssignmentsCompleteMatchingPredicate(ctx, id, options, DeviceManagementResourceAccessProfileAssignmentOperationPredicate{})
}

// ListResourceAccessProfileAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceAccessProfileAssignmentClient) ListResourceAccessProfileAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementResourceAccessProfileId, options ListResourceAccessProfileAssignmentsOperationOptions, predicate DeviceManagementResourceAccessProfileAssignmentOperationPredicate) (result ListResourceAccessProfileAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementResourceAccessProfileAssignment, 0)

	resp, err := c.ListResourceAccessProfileAssignments(ctx, id, options)
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

	result = ListResourceAccessProfileAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
