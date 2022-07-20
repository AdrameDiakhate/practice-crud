package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type category struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// album represents data about a record album.
type product struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Category category
}

var Category = []category{
	{
		ID:    "1",
		Title: "Category 1",
	},
	{
		ID:    "2",
		Title: "Category 2",
	},
	{
		ID:    "3",
		Title: "Category 3",
	},
	{
		ID:    "4",
		Title: "Category 4",
	},
	{
		ID:    "5",
		Title: "Category 5",
	},
}

// albums slice to seed record album data.
var products = []product{
	{
		ID:       "1",
		Title:    "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Body:     "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
		Category: Category[1],
	},
	{
		ID:       "2",
		Title:    "qui est esse",
		Body:     "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
		Category: Category[1],
	},
	{
		ID:       "3",
		Title:    "ea molestias quasi exercitationem repellat qui ipsa sit aut",
		Body:     "et iusto sed quo iure\nvoluptatem occaecati omnis eligendi aut ad\nvoluptatem doloribus vel accusantium quis pariatur\nmolestiae porro eius odio et labore et velit aut",
		Category: Category[1],
	},
	{
		ID:       "4",
		Title:    "eum et est occaecati",
		Body:     "ullam et saepe reiciendis voluptatem adipisci\nsit amet autem assumenda provident rerum culpa\nquis hic commodi nesciunt rem tenetur doloremque ipsam iure\nquis sunt voluptatem rerum illo velit",
		Category: Category[1],
	},
	{
		ID:       "5",
		Title:    "nesciunt quas odio",
		Body:     "repudiandae veniam quaerat sunt sed\nalias aut fugiat sit autem sed est\nvoluptatem omnis possimus esse voluptatibus quis\nest aut tenetur dolor neque",
		Category: Category[1],
	},
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)        //yes
	router.GET("/products/:id", getProductByID) //yes
	router.GET("/products", searchByTitle)      //yes
	router.POST("/products", postProducts)      //yes
	router.PUT("/products/:id", updateProduct)  //yes
	//router.DELETE("/products/:id",deleteProduct)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

// postAlbums adds an album from JSON received in the request body.
func postProducts(c *gin.Context) {
	var newProduct product

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	// Add the new album to the slice.
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getProductByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

// Update a product
func updateProduct(c *gin.Context) {
	// Get model if exist
	var updateProduct product
	if err := c.BindJSON(&updateProduct); err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updateProduct})
}

/*func deleteProduct(){

  }*/
func searchByTitle(c *gin.Context) {
	title := c.Param("title")
	c.String(http.StatusOK, "Hello %s", title)
	for _, a := range products {
		if a.Title == title {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}
