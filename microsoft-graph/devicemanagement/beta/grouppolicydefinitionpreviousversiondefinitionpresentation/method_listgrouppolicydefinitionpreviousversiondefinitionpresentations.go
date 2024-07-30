package grouppolicydefinitionpreviousversiondefinitionpresentation

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

type ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyPresentation
}

type ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyPresentation
}

type ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyDefinitionPreviousVersionDefinitionPresentations ...
func (c GroupPolicyDefinitionPreviousVersionDefinitionPresentationClient) ListGroupPolicyDefinitionPreviousVersionDefinitionPresentations(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId) (result ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsCustomPager{},
		Path:       fmt.Sprintf("%s/previousVersionDefinition/presentations", id.ID()),
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

// ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsComplete retrieves all the results into a single object
func (c GroupPolicyDefinitionPreviousVersionDefinitionPresentationClient) ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsComplete(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId) (ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsCompleteResult, error) {
	return c.ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsCompleteMatchingPredicate(ctx, id, GroupPolicyPresentationOperationPredicate{})
}

// ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyDefinitionPreviousVersionDefinitionPresentationClient) ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyDefinitionId, predicate GroupPolicyPresentationOperationPredicate) (result ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsCompleteResult, err error) {
	items := make([]beta.GroupPolicyPresentation, 0)

	resp, err := c.ListGroupPolicyDefinitionPreviousVersionDefinitionPresentations(ctx, id)
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

	result = ListGroupPolicyDefinitionPreviousVersionDefinitionPresentationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
