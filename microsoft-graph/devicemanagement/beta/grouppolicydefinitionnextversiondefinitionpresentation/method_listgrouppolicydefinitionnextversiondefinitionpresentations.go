package grouppolicydefinitionnextversiondefinitionpresentation

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

type ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyPresentation
}

type ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyPresentation
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

// ListGroupPolicyDefinitionNextVersionDefinitionPresentations ...
func (c GroupPolicyDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionNextVersionDefinitionPresentations(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId) (result ListGroupPolicyDefinitionNextVersionDefinitionPresentationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCustomPager{},
		Path:       fmt.Sprintf("%s/nextVersionDefinition/presentations", id.ID()),
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
		Values *[]beta.GroupPolicyPresentation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyDefinitionNextVersionDefinitionPresentationsComplete retrieves all the results into a single object
func (c GroupPolicyDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionNextVersionDefinitionPresentationsComplete(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId) (ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteResult, error) {
	return c.ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteMatchingPredicate(ctx, id, GroupPolicyPresentationOperationPredicate{})
}

// ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId, predicate GroupPolicyPresentationOperationPredicate) (result ListGroupPolicyDefinitionNextVersionDefinitionPresentationsCompleteResult, err error) {
	items := make([]beta.GroupPolicyPresentation, 0)

	resp, err := c.ListGroupPolicyDefinitionNextVersionDefinitionPresentations(ctx, id)
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
