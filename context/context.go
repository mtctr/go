package main

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data, err := store.Fetch(ctx)

		if err != nil {
			fmt.Printf("an error ocurred")
			return
		}

		fmt.Fprintf(w, data)
	}

}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
