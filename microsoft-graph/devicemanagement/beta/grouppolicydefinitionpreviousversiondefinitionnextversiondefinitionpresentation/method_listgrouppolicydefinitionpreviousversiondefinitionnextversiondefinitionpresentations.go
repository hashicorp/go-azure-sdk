package grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentation

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

type ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyPresentation
}

type ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyPresentation
}

type ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentations ...
func (c GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentations(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId) (result ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsCustomPager{},
		Path:       fmt.Sprintf("%s/previousVersionDefinition/nextVersionDefinition/presentations", id.ID()),
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

// ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsComplete retrieves all the results into a single object
func (c GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsComplete(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId) (ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsCompleteResult, error) {
	return c.ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsCompleteMatchingPredicate(ctx, id, GroupPolicyPresentationOperationPredicate{})
}

// ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient) ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId, predicate GroupPolicyPresentationOperationPredicate) (result ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsCompleteResult, err error) {
	items := make([]beta.GroupPolicyPresentation, 0)

	resp, err := c.ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentations(ctx, id)
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

	result = ListGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
