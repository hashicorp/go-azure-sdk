package grouppolicyuploadeddefinitionfiledefinition

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

type ListGroupPolicyUploadedDefinitionFileDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyDefinition
}

type ListGroupPolicyUploadedDefinitionFileDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyDefinition
}

type ListGroupPolicyUploadedDefinitionFileDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyUploadedDefinitionFileDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyUploadedDefinitionFileDefinitions ...
func (c GroupPolicyUploadedDefinitionFileDefinitionClient) ListGroupPolicyUploadedDefinitionFileDefinitions(ctx context.Context, id DeviceManagementGroupPolicyUploadedDefinitionFileId) (result ListGroupPolicyUploadedDefinitionFileDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyUploadedDefinitionFileDefinitionsCustomPager{},
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

// ListGroupPolicyUploadedDefinitionFileDefinitionsComplete retrieves all the results into a single object
func (c GroupPolicyUploadedDefinitionFileDefinitionClient) ListGroupPolicyUploadedDefinitionFileDefinitionsComplete(ctx context.Context, id DeviceManagementGroupPolicyUploadedDefinitionFileId) (ListGroupPolicyUploadedDefinitionFileDefinitionsCompleteResult, error) {
	return c.ListGroupPolicyUploadedDefinitionFileDefinitionsCompleteMatchingPredicate(ctx, id, GroupPolicyDefinitionOperationPredicate{})
}

// ListGroupPolicyUploadedDefinitionFileDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyUploadedDefinitionFileDefinitionClient) ListGroupPolicyUploadedDefinitionFileDefinitionsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyUploadedDefinitionFileId, predicate GroupPolicyDefinitionOperationPredicate) (result ListGroupPolicyUploadedDefinitionFileDefinitionsCompleteResult, err error) {
	items := make([]beta.GroupPolicyDefinition, 0)

	resp, err := c.ListGroupPolicyUploadedDefinitionFileDefinitions(ctx, id)
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

	result = ListGroupPolicyUploadedDefinitionFileDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
