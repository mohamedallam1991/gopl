package resources

type Author struct {
	Id        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

func GetAllAuthors() []Author {
	// a := make([]Author, 0)
	return getauthors()
}

func GetAuthor() Author {
	author := Author{
		Id:        "author-1",
		Firstname: "Christiano",
		Lastname:  "Ronaldo",
		Username:  "nraboy",
		Password:  "pass",
	}

	return author

}

func getauthors() []Author {
	authors := make([]Author, 0)
	author1 := Author{
		Id:        "1",
		Firstname: "Christiano",
		Lastname:  "Ronaldo",
		Username:  "nraboy",
		Password:  "pass",
	}

	author2 := Author{
		Id:        "2",
		Firstname: "Maria",
		Lastname:  "Rabboy",
		Username:  "mraboy",
		Password:  "abc123",
	}

	authors = []Author{author1, author2}
	return authors
}
