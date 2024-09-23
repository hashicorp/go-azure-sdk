package b2xuserflowuserattributeassignment

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

type ListB2xUserFlowUserAttributeAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityUserFlowAttributeAssignment
}

type ListB2xUserFlowUserAttributeAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityUserFlowAttributeAssignment
}

type ListB2xUserFlowUserAttributeAssignmentsOperationOptions struct {
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

func DefaultListB2xUserFlowUserAttributeAssignmentsOperationOptions() ListB2xUserFlowUserAttributeAssignmentsOperationOptions {
	return ListB2xUserFlowUserAttributeAssignmentsOperationOptions{}
}

func (o ListB2xUserFlowUserAttributeAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListB2xUserFlowUserAttributeAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListB2xUserFlowUserAttributeAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListB2xUserFlowUserAttributeAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowUserAttributeAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlowUserAttributeAssignments - List userAttributeAssignments. Get the identityUserFlowAttributeAssignment
// resources from the userAttributeAssignments navigation property in a b2xIdentityUserFlow.
func (c B2xUserFlowUserAttributeAssignmentClient) ListB2xUserFlowUserAttributeAssignments(ctx context.Context, id stable.IdentityB2xUserFlowId, options ListB2xUserFlowUserAttributeAssignmentsOperationOptions) (result ListB2xUserFlowUserAttributeAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListB2xUserFlowUserAttributeAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/userAttributeAssignments", id.ID()),
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
		Values *[]stable.IdentityUserFlowAttributeAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListB2xUserFlowUserAttributeAssignmentsComplete retrieves all the results into a single object
func (c B2xUserFlowUserAttributeAssignmentClient) ListB2xUserFlowUserAttributeAssignmentsComplete(ctx context.Context, id stable.IdentityB2xUserFlowId, options ListB2xUserFlowUserAttributeAssignmentsOperationOptions) (ListB2xUserFlowUserAttributeAssignmentsCompleteResult, error) {
	return c.ListB2xUserFlowUserAttributeAssignmentsCompleteMatchingPredicate(ctx, id, options, IdentityUserFlowAttributeAssignmentOperationPredicate{})
}

// ListB2xUserFlowUserAttributeAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c B2xUserFlowUserAttributeAssignmentClient) ListB2xUserFlowUserAttributeAssignmentsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityB2xUserFlowId, options ListB2xUserFlowUserAttributeAssignmentsOperationOptions, predicate IdentityUserFlowAttributeAssignmentOperationPredicate) (result ListB2xUserFlowUserAttributeAssignmentsCompleteResult, err error) {
	items := make([]stable.IdentityUserFlowAttributeAssignment, 0)

	resp, err := c.ListB2xUserFlowUserAttributeAssignments(ctx, id, options)
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

	result = ListB2xUserFlowUserAttributeAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
