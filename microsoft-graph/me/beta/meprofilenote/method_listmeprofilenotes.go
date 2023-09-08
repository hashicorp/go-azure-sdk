package meprofilenote

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

type ListMeProfileNotesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonAnnotationCollectionResponse
}

type ListMeProfileNotesCompleteResult struct {
	Items []models.PersonAnnotationCollectionResponse
}

// ListMeProfileNotes ...
func (c MeProfileNoteClient) ListMeProfileNotes(ctx context.Context) (result ListMeProfileNotesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/notes",
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

// ListMeProfileNotesComplete retrieves all the results into a single object
func (c MeProfileNoteClient) ListMeProfileNotesComplete(ctx context.Context) (ListMeProfileNotesCompleteResult, error) {
	return c.ListMeProfileNotesCompleteMatchingPredicate(ctx, models.PersonAnnotationCollectionResponseOperationPredicate{})
}

// ListMeProfileNotesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileNoteClient) ListMeProfileNotesCompleteMatchingPredicate(ctx context.Context, predicate models.PersonAnnotationCollectionResponseOperationPredicate) (result ListMeProfileNotesCompleteResult, err error) {
	items := make([]models.PersonAnnotationCollectionResponse, 0)

	resp, err := c.ListMeProfileNotes(ctx)
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

	result = ListMeProfileNotesCompleteResult{
		Items: items,
	}
	return
}
