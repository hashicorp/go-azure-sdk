package grouppolicydefinitionfiledefinition

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

type ListGroupPolicyDefinitionFileDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyDefinition
}

type ListGroupPolicyDefinitionFileDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyDefinition
}

type ListGroupPolicyDefinitionFileDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyDefinitionFileDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyDefinitionFileDefinitions ...
func (c GroupPolicyDefinitionFileDefinitionClient) ListGroupPolicyDefinitionFileDefinitions(ctx context.Context, id DeviceManagementGroupPolicyDefinitionFileId) (result ListGroupPolicyDefinitionFileDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyDefinitionFileDefinitionsCustomPager{},
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

// ListGroupPolicyDefinitionFileDefinitionsComplete retrieves all the results into a single object
func (c GroupPolicyDefinitionFileDefinitionClient) ListGroupPolicyDefinitionFileDefinitionsComplete(ctx context.Context, id DeviceManagementGroupPolicyDefinitionFileId) (ListGroupPolicyDefinitionFileDefinitionsCompleteResult, error) {
	return c.ListGroupPolicyDefinitionFileDefinitionsCompleteMatchingPredicate(ctx, id, GroupPolicyDefinitionOperationPredicate{})
}

// ListGroupPolicyDefinitionFileDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyDefinitionFileDefinitionClient) ListGroupPolicyDefinitionFileDefinitionsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyDefinitionFileId, predicate GroupPolicyDefinitionOperationPredicate) (result ListGroupPolicyDefinitionFileDefinitionsCompleteResult, err error) {
	items := make([]beta.GroupPolicyDefinition, 0)

	resp, err := c.ListGroupPolicyDefinitionFileDefinitions(ctx, id)
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

	result = ListGroupPolicyDefinitionFileDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
