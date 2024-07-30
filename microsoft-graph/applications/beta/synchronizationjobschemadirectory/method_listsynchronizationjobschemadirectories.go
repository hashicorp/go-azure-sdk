package synchronizationjobschemadirectory

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

type ListSynchronizationJobSchemaDirectoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryDefinition
}

type ListSynchronizationJobSchemaDirectoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryDefinition
}

type ListSynchronizationJobSchemaDirectoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSynchronizationJobSchemaDirectoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSynchronizationJobSchemaDirectories ...
func (c SynchronizationJobSchemaDirectoryClient) ListSynchronizationJobSchemaDirectories(ctx context.Context, id ApplicationIdSynchronizationJobId) (result ListSynchronizationJobSchemaDirectoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSynchronizationJobSchemaDirectoriesCustomPager{},
		Path:       fmt.Sprintf("%s/schema/directories", id.ID()),
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
		Values *[]beta.DirectoryDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSynchronizationJobSchemaDirectoriesComplete retrieves all the results into a single object
func (c SynchronizationJobSchemaDirectoryClient) ListSynchronizationJobSchemaDirectoriesComplete(ctx context.Context, id ApplicationIdSynchronizationJobId) (ListSynchronizationJobSchemaDirectoriesCompleteResult, error) {
	return c.ListSynchronizationJobSchemaDirectoriesCompleteMatchingPredicate(ctx, id, DirectoryDefinitionOperationPredicate{})
}

// ListSynchronizationJobSchemaDirectoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SynchronizationJobSchemaDirectoryClient) ListSynchronizationJobSchemaDirectoriesCompleteMatchingPredicate(ctx context.Context, id ApplicationIdSynchronizationJobId, predicate DirectoryDefinitionOperationPredicate) (result ListSynchronizationJobSchemaDirectoriesCompleteResult, err error) {
	items := make([]beta.DirectoryDefinition, 0)

	resp, err := c.ListSynchronizationJobSchemaDirectories(ctx, id)
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

	result = ListSynchronizationJobSchemaDirectoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
