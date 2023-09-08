package userprofilenote

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdProfileNotesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonAnnotationCollectionResponse
}

type ListUserByIdProfileNotesCompleteResult struct {
	Items []models.PersonAnnotationCollectionResponse
}

// ListUserByIdProfileNotes ...
func (c UserProfileNoteClient) ListUserByIdProfileNotes(ctx context.Context, id UserId) (result ListUserByIdProfileNotesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/notes", id.ID()),
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
		Values *[]models.PersonAnnotationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileNotesComplete retrieves all the results into a single object
func (c UserProfileNoteClient) ListUserByIdProfileNotesComplete(ctx context.Context, id UserId) (ListUserByIdProfileNotesCompleteResult, error) {
	return c.ListUserByIdProfileNotesCompleteMatchingPredicate(ctx, id, models.PersonAnnotationCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileNotesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileNoteClient) ListUserByIdProfileNotesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PersonAnnotationCollectionResponseOperationPredicate) (result ListUserByIdProfileNotesCompleteResult, err error) {
	items := make([]models.PersonAnnotationCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileNotes(ctx, id)
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

	result = ListUserByIdProfileNotesCompleteResult{
		Items: items,
	}
	return
}
