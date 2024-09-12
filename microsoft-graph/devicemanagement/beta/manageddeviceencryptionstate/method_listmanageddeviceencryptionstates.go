package manageddeviceencryptionstate

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

type ListManagedDeviceEncryptionStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedDeviceEncryptionState
}

type ListManagedDeviceEncryptionStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedDeviceEncryptionState
}

type ListManagedDeviceEncryptionStatesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListManagedDeviceEncryptionStatesOperationOptions() ListManagedDeviceEncryptionStatesOperationOptions {
	return ListManagedDeviceEncryptionStatesOperationOptions{}
}

func (o ListManagedDeviceEncryptionStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListManagedDeviceEncryptionStatesOperationOptions) ToOData() *odata.Query {
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

func (o ListManagedDeviceEncryptionStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListManagedDeviceEncryptionStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceEncryptionStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceEncryptionStates - Get managedDeviceEncryptionStates from deviceManagement. Encryption report for
// devices in this account
func (c ManagedDeviceEncryptionStateClient) ListManagedDeviceEncryptionStates(ctx context.Context, options ListManagedDeviceEncryptionStatesOperationOptions) (result ListManagedDeviceEncryptionStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListManagedDeviceEncryptionStatesCustomPager{},
		Path:          "/deviceManagement/managedDeviceEncryptionStates",
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
		Values *[]beta.ManagedDeviceEncryptionState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedDeviceEncryptionStatesComplete retrieves all the results into a single object
func (c ManagedDeviceEncryptionStateClient) ListManagedDeviceEncryptionStatesComplete(ctx context.Context, options ListManagedDeviceEncryptionStatesOperationOptions) (ListManagedDeviceEncryptionStatesCompleteResult, error) {
	return c.ListManagedDeviceEncryptionStatesCompleteMatchingPredicate(ctx, options, ManagedDeviceEncryptionStateOperationPredicate{})
}

// ListManagedDeviceEncryptionStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceEncryptionStateClient) ListManagedDeviceEncryptionStatesCompleteMatchingPredicate(ctx context.Context, options ListManagedDeviceEncryptionStatesOperationOptions, predicate ManagedDeviceEncryptionStateOperationPredicate) (result ListManagedDeviceEncryptionStatesCompleteResult, err error) {
	items := make([]beta.ManagedDeviceEncryptionState, 0)

	resp, err := c.ListManagedDeviceEncryptionStates(ctx, options)
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

	result = ListManagedDeviceEncryptionStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
