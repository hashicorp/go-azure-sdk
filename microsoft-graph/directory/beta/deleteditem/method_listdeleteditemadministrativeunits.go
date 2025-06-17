package deleteditem

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

type ListDeletedItemAdministrativeUnitsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AdministrativeUnit
}

type ListDeletedItemAdministrativeUnitsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AdministrativeUnit
}

type ListDeletedItemAdministrativeUnitsOperationOptions struct {
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

func DefaultListDeletedItemAdministrativeUnitsOperationOptions() ListDeletedItemAdministrativeUnitsOperationOptions {
	return ListDeletedItemAdministrativeUnitsOperationOptions{}
}

func (o ListDeletedItemAdministrativeUnitsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeletedItemAdministrativeUnitsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeletedItemAdministrativeUnitsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeletedItemAdministrativeUnitsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeletedItemAdministrativeUnitsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeletedItemAdministrativeUnits - Get the items of type microsoft.graph.administrativeUnit in the
// microsoft.graph.directoryObject collection
func (c DeletedItemClient) ListDeletedItemAdministrativeUnits(ctx context.Context, options ListDeletedItemAdministrativeUnitsOperationOptions) (result ListDeletedItemAdministrativeUnitsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeletedItemAdministrativeUnitsCustomPager{},
		Path:          "/directory/deletedItems/administrativeUnit",
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
		Values *[]beta.AdministrativeUnit `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeletedItemAdministrativeUnitsComplete retrieves all the results into a single object
func (c DeletedItemClient) ListDeletedItemAdministrativeUnitsComplete(ctx context.Context, options ListDeletedItemAdministrativeUnitsOperationOptions) (ListDeletedItemAdministrativeUnitsCompleteResult, error) {
	return c.ListDeletedItemAdministrativeUnitsCompleteMatchingPredicate(ctx, options, AdministrativeUnitOperationPredicate{})
}

// ListDeletedItemAdministrativeUnitsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeletedItemClient) ListDeletedItemAdministrativeUnitsCompleteMatchingPredicate(ctx context.Context, options ListDeletedItemAdministrativeUnitsOperationOptions, predicate AdministrativeUnitOperationPredicate) (result ListDeletedItemAdministrativeUnitsCompleteResult, err error) {
	items := make([]beta.AdministrativeUnit, 0)

	resp, err := c.ListDeletedItemAdministrativeUnits(ctx, options)
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

	result = ListDeletedItemAdministrativeUnitsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
