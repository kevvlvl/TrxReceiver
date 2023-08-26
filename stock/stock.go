package stock

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stock struct {
	Id     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Value  float32 `json:"value"`
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var s Stock
	parseTransactionBody(r, &s)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {

	stock := r.Context().Value("stock").(*Stock)

	fmt.Println("PUT/Update trx: ", stock.Id)

	var s Stock
	parseTransactionBody(r, &s)
}

func TransactionCtx(next http.Handler) http.Handler {
	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	var article *Article
	//	var err error
	//
	//	if articleID := chi.URLParam(r, "articleID"); articleID != "" {
	//		article, err = dbGetArticle(articleID)
	//	} else if articleSlug := chi.URLParam(r, "articleSlug"); articleSlug != "" {
	//		article, err = dbGetArticleBySlug(articleSlug)
	//	} else {
	//		render.Render(w, r, ErrNotFound)
	//		return
	//	}
	//	if err != nil {
	//		render.Render(w, r, ErrNotFound)
	//		return
	//	}
	//
	//	ctx := context.WithValue(r.Context(), "article", article)
	//	next.ServeHTTP(w, r.WithContext(ctx))
	//})
	return nil
}

func parseTransactionBody(r *http.Request, s *Stock) {

	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		fmt.Println("Error parsing Request Body: ", err)
	}

	fmt.Printf("Parsed JSON successfully: %+v", s)
}
