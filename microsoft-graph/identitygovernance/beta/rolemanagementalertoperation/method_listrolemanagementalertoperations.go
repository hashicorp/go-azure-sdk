package rolemanagementalertoperation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRoleManagementAlertOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.LongRunningOperation
}

type ListRoleManagementAlertOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.LongRunningOperation
}

type ListRoleManagementAlertOperationsOperationOptions struct {
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

func DefaultListRoleManagementAlertOperationsOperationOptions() ListRoleManagementAlertOperationsOperationOptions {
	return ListRoleManagementAlertOperationsOperationOptions{}
}

func (o ListRoleManagementAlertOperationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRoleManagementAlertOperationsOperationOptions) ToOData() *odata.Query {
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

func (o ListRoleManagementAlertOperationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRoleManagementAlertOperationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleManagementAlertOperationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleManagementAlertOperations - Get operations from identityGovernance. Represents operations on resources that
// take a long time to complete and can run in the background until completion.
func (c RoleManagementAlertOperationClient) ListRoleManagementAlertOperations(ctx context.Context, options ListRoleManagementAlertOperationsOperationOptions) (result ListRoleManagementAlertOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRoleManagementAlertOperationsCustomPager{},
		Path:          "/identityGovernance/roleManagementAlerts/operations",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.LongRunningOperation, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalLongRunningOperationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.LongRunningOperation (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListRoleManagementAlertOperationsComplete retrieves all the results into a single object
func (c RoleManagementAlertOperationClient) ListRoleManagementAlertOperationsComplete(ctx context.Context, options ListRoleManagementAlertOperationsOperationOptions) (ListRoleManagementAlertOperationsCompleteResult, error) {
	return c.ListRoleManagementAlertOperationsCompleteMatchingPredicate(ctx, options, LongRunningOperationOperationPredicate{})
}

// ListRoleManagementAlertOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementAlertOperationClient) ListRoleManagementAlertOperationsCompleteMatchingPredicate(ctx context.Context, options ListRoleManagementAlertOperationsOperationOptions, predicate LongRunningOperationOperationPredicate) (result ListRoleManagementAlertOperationsCompleteResult, err error) {
	items := make([]beta.LongRunningOperation, 0)

	resp, err := c.ListRoleManagementAlertOperations(ctx, options)
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

	result = ListRoleManagementAlertOperationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
