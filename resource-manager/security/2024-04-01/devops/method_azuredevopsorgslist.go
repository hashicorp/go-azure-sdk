package devops

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureDevOpsOrgsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AzureDevOpsOrg
}

type AzureDevOpsOrgsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AzureDevOpsOrg
}

type AzureDevOpsOrgsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AzureDevOpsOrgsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AzureDevOpsOrgsList ...
func (c DevOpsClient) AzureDevOpsOrgsList(ctx context.Context, id SecurityConnectorId) (result AzureDevOpsOrgsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AzureDevOpsOrgsListCustomPager{},
		Path:       fmt.Sprintf("%s/devops/default/azureDevOpsOrgs", id.ID()),
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
		Values *[]AzureDevOpsOrg `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AzureDevOpsOrgsListComplete retrieves all the results into a single object
func (c DevOpsClient) AzureDevOpsOrgsListComplete(ctx context.Context, id SecurityConnectorId) (AzureDevOpsOrgsListCompleteResult, error) {
	return c.AzureDevOpsOrgsListCompleteMatchingPredicate(ctx, id, AzureDevOpsOrgOperationPredicate{})
}

// AzureDevOpsOrgsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DevOpsClient) AzureDevOpsOrgsListCompleteMatchingPredicate(ctx context.Context, id SecurityConnectorId, predicate AzureDevOpsOrgOperationPredicate) (result AzureDevOpsOrgsListCompleteResult, err error) {
	items := make([]AzureDevOpsOrg, 0)

	resp, err := c.AzureDevOpsOrgsList(ctx, id)
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

	result = AzureDevOpsOrgsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
