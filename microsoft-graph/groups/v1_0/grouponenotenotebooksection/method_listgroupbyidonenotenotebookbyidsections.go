package grouponenotenotebooksection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdOnenoteNotebookByIdSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenoteSectionCollectionResponse
}

type ListGroupByIdOnenoteNotebookByIdSectionsCompleteResult struct {
	Items []models.OnenoteSectionCollectionResponse
}

// ListGroupByIdOnenoteNotebookByIdSections ...
func (c GroupOnenoteNotebookSectionClient) ListGroupByIdOnenoteNotebookByIdSections(ctx context.Context, id GroupOnenoteNotebookId) (result ListGroupByIdOnenoteNotebookByIdSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/sections", id.ID()),
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
		Values *[]models.OnenoteSectionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdOnenoteNotebookByIdSectionsComplete retrieves all the results into a single object
func (c GroupOnenoteNotebookSectionClient) ListGroupByIdOnenoteNotebookByIdSectionsComplete(ctx context.Context, id GroupOnenoteNotebookId) (ListGroupByIdOnenoteNotebookByIdSectionsCompleteResult, error) {
	return c.ListGroupByIdOnenoteNotebookByIdSectionsCompleteMatchingPredicate(ctx, id, models.OnenoteSectionCollectionResponseOperationPredicate{})
}

// ListGroupByIdOnenoteNotebookByIdSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupOnenoteNotebookSectionClient) ListGroupByIdOnenoteNotebookByIdSectionsCompleteMatchingPredicate(ctx context.Context, id GroupOnenoteNotebookId, predicate models.OnenoteSectionCollectionResponseOperationPredicate) (result ListGroupByIdOnenoteNotebookByIdSectionsCompleteResult, err error) {
	items := make([]models.OnenoteSectionCollectionResponse, 0)

	resp, err := c.ListGroupByIdOnenoteNotebookByIdSections(ctx, id)
	if err != nil {
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

	result = ListGroupByIdOnenoteNotebookByIdSectionsCompleteResult{
		Items: items,
	}
	return
}
