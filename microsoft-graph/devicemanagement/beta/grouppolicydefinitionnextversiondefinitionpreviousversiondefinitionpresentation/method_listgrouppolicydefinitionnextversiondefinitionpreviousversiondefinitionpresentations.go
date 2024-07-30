package grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentation

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

type ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyPresentation
}

type ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyPresentation
}

type ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentations ...
func (c GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient) ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentations(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId) (result ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsCustomPager{},
		Path:       fmt.Sprintf("%s/nextVersionDefinition/previousVersionDefinition/presentations", id.ID()),
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

// ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsComplete retrieves all the results into a single object
func (c GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient) ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsComplete(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId) (ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsCompleteResult, error) {
	return c.ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsCompleteMatchingPredicate(ctx, id, GroupPolicyPresentationOperationPredicate{})
}

// ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient) ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId, predicate GroupPolicyPresentationOperationPredicate) (result ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsCompleteResult, err error) {
	items := make([]beta.GroupPolicyPresentation, 0)

	resp, err := c.ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentations(ctx, id)
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

	result = ListGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
