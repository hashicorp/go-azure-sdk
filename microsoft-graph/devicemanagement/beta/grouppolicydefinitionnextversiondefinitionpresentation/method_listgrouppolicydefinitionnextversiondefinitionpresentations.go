package grouppolicydefinitionnextversiondefinitionpresentation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyPresentation
}

type ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyPresentation
}

type ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationOptions() ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationOptions {
	return ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationOptions{}
}

func (o ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationOptions) ToOData() *odata.Query {
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

func (o ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyDefinitionNextVersionDefinitionPresentations - Get presentations from deviceManagement. The group
// policy presentations associated with the definition.
func (c GroupPolicyDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionNextVersionDefinitionPresentations(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationOptions) (result ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCustomPager{},
		Path:          fmt.Sprintf("%s/nextVersionDefinition/presentations", id.ID()),
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

	temp := make([]beta.GroupPolicyPresentation, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalGroupPolicyPresentationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.GroupPolicyPresentation (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListGroupPolicyDefinitionNextVersionDefinitionPresentationsComplete retrieves all the results into a single object
func (c GroupPolicyDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionNextVersionDefinitionPresentationsComplete(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationOptions) (ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteResult, error) {
	return c.ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteMatchingPredicate(ctx, id, options, GroupPolicyPresentationOperationPredicate{})
}

// ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationOptions, predicate GroupPolicyPresentationOperationPredicate) (result ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteResult, err error) {
	items := make([]beta.GroupPolicyPresentation, 0)

	resp, err := c.ListGroupPolicyDefinitionNextVersionDefinitionPresentations(ctx, id, options)
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

	result = ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
