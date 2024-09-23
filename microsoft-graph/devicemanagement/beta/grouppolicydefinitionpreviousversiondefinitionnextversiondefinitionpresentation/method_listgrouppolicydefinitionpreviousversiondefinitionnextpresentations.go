package grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentation

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

type ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyPresentation
}

type ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyPresentation
}

type ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationOptions struct {
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

func DefaultListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationOptions() ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationOptions {
	return ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationOptions{}
}

func (o ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationOptions) ToOData() *odata.Query {
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

func (o ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentations - Get presentations from deviceManagement. The
// group policy presentations associated with the definition.
func (c GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentations(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationOptions) (result ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsCustomPager{},
		Path:          fmt.Sprintf("%s/previousVersionDefinition/nextVersionDefinition/presentations", id.ID()),
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

// ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsComplete retrieves all the results into a single object
func (c GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsComplete(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationOptions) (ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsCompleteResult, error) {
	return c.ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsCompleteMatchingPredicate(ctx, id, options, GroupPolicyPresentationOperationPredicate{})
}

// ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsOperationOptions, predicate GroupPolicyPresentationOperationPredicate) (result ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsCompleteResult, err error) {
	items := make([]beta.GroupPolicyPresentation, 0)

	resp, err := c.ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentations(ctx, id, options)
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

	result = ListGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
