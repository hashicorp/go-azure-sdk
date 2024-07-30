package grouppolicycategorydefinition

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

type ListGroupPolicyCategoryDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyDefinition
}

type ListGroupPolicyCategoryDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyDefinition
}

type ListGroupPolicyCategoryDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyCategoryDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyCategoryDefinitions ...
func (c GroupPolicyCategoryDefinitionClient) ListGroupPolicyCategoryDefinitions(ctx context.Context, id DeviceManagementGroupPolicyCategoryId) (result ListGroupPolicyCategoryDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyCategoryDefinitionsCustomPager{},
		Path:       fmt.Sprintf("%s/definitions", id.ID()),
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
		Values *[]beta.GroupPolicyDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyCategoryDefinitionsComplete retrieves all the results into a single object
func (c GroupPolicyCategoryDefinitionClient) ListGroupPolicyCategoryDefinitionsComplete(ctx context.Context, id DeviceManagementGroupPolicyCategoryId) (ListGroupPolicyCategoryDefinitionsCompleteResult, error) {
	return c.ListGroupPolicyCategoryDefinitionsCompleteMatchingPredicate(ctx, id, GroupPolicyDefinitionOperationPredicate{})
}

// ListGroupPolicyCategoryDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyCategoryDefinitionClient) ListGroupPolicyCategoryDefinitionsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyCategoryId, predicate GroupPolicyDefinitionOperationPredicate) (result ListGroupPolicyCategoryDefinitionsCompleteResult, err error) {
	items := make([]beta.GroupPolicyDefinition, 0)

	resp, err := c.ListGroupPolicyCategoryDefinitions(ctx, id)
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

	result = ListGroupPolicyCategoryDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
