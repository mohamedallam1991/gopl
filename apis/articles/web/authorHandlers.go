package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedallam1991/gopl/apis/articles/resources"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var authors []resources.Author = resources.GetAllAuthors()

func AuthorRetriveAllEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	json.NewEncoder(response).Encode(authors)
}

func AuthorRetriveEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)

	for _, author := range authors {
		if author.Id == params["id"] {
			json.NewEncoder(response).Encode(author)
			return
		}
	}

	json.NewEncoder(response).Encode(resources.Author{})
}

func AuthorStoreEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var author resources.Author
	json.NewDecoder(request.Body).Decode(&author)

	// hashing the incoming password
	hash, _ := bcrypt.GenerateFromPassword([]byte(author.Password), 10)
	// assigning the uuid
	author.Id = uuid.Must(uuid.NewV4(), errors.New("cant uuid must")).String()
	// assigning the hased password
	author.Password = string(hash)
	authors = append(authors, author)
	json.NewEncoder(response).Encode(author)
}

func AuthorLoginEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var data resources.Author
	json.NewDecoder(request.Body).Decode(&data)
	for _, author := range authors {
		if author.Username == data.Username {
			err := bcrypt.CompareHashAndPassword([]byte(author.Password), []byte(data.Password))
			if err != nil {
				response.WriteHeader(500)
				response.Write([]byte(`{"message": "invalid password"}`))
			}
			json.NewEncoder(response).Encode(author)
			return
		}
	}
	response.Write([]byte(`{"message": "invalid username"}`))
}

func AuthorUpdateEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	var changes resources.Author
	json.NewDecoder(request.Body).Decode(&changes)
	for index, author := range authors {
		if author.Id == params["id"] {
			if changes.Firstname != "" {
				author.Firstname = changes.Firstname
			}
			if changes.Lastname != "" {
				author.Lastname = changes.Lastname
			}
			if changes.Username != "" {
				author.Username = changes.Username
			}
			if changes.Password != "" {
				author.Password = changes.Password
			}
			authors[index] = author
			json.NewEncoder(response).Encode(authors)
			return
		}
	}

	json.NewEncoder(response).Encode(resources.Author{})
}

func AuthorDeleteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	for index, author := range authors {
		if author.Id == params["id"] {
			authors = append(authors[:index], authors[index+1:]...)
			json.NewEncoder(response).Encode(authors)
			return
		}
	}
	json.NewEncoder(response).Encode(resources.Author{})
}
