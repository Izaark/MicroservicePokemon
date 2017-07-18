package controllers

import (
	"fmt"
	"log"
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nichos_sociales/pokemon/config"
	"github.com/nichos_sociales/pokemon/models"
	r "gopkg.in/gorethink/gorethink.v3"
	"github.com/olahol/go-imageupload"
	"net/http"
	"time"
)

func PokemonRouter() {

	router := gin.Default()
	v1 := router.Group("api")
	{

		v1.GET("/pokemon", GetPokemons)
		v1.GET("/pokemon/:id", GetPokemon)
		v1.POST("/pokemon", PostPokemon)
		v1.GET("/index", Index)
		v1.POST("/upload", PostImage)	

	}

	router.Run(":1338")
}

var Static_url = "assets/images/pokemons/"

func GetPokemons(c *gin.Context) {
	session := config.InitDb()
	defer session.Close()

	cursor, err := r.Table("pokedex").Run(session)
	defer cursor.Close()

	if err != nil {
		c.JSON(422, gin.H{"error": "Error en la consulta de pokemones"})
		log.Panic(err)
		//fmt.Println(err)
	}

	if cursor.IsNil() {
		//fmt.Println("Row not found")
		c.JSON(404, gin.H{"error": "Pokemones no encontrados"})
		return
	}

	var pokemon []models.Pokemon
	err = cursor.All(&pokemon)
	if err != nil {
		log.Panic(err)
		//fmt.Println(err)
	} else {
		c.JSON(200, pokemon)
		//fmt.Println(pokemon)
	}
}

func GetPokemon(c *gin.Context) {
	session := config.InitDb()
	defer session.Close()

	id := c.Params.ByName("id")
	//fmt.Println("id:", id)

	cursor, err := r.Table("pokedex").Get(id).Run(session)
	defer cursor.Close()
	if err != nil {
		c.JSON(422, gin.H{"error": "Error en la consulta de pokemon"})
		//fmt.Println(err)
		log.Panic(err)
	}

	if cursor.IsNil() {
		//fmt.Println("Row not found")
		c.JSON(404, gin.H{"error": "Pokemon no encontrado"})
		return
	}

	var pokemon models.Pokemon
	err = cursor.One(&pokemon)
	if err != nil {
		log.Panic(err)
		//fmt.Println(err)
	} else {
		c.JSON(200, pokemon)
		//fmt.Println(pokemon)
	}
}

func PostPokemon(c *gin.Context) {
	session := config.InitDb()
	defer session.Close()

	var pokemon models.PokemonPost
	err := c.Bind(&pokemon)
	if err != nil {
		c.JSON(422, gin.H{"error": "Error en la captura de campos"})
		log.Panic(err)
		//fmt.Println(err)
	}

	if pokemon.Name == "" {
		c.JSON(422, gin.H{"error": "Faltan campos requeridos"})
		return
	}

	cursor, err := r.Table("pokedex").Insert(pokemon).Run(session)
	if err != nil && !cursor.IsNil() {
		log.Panic(err)
		//fmt.Println(err)
	} else {
		c.JSON(201, gin.H{"success": pokemon})
	}
}


var currentImage *imageupload.Image
func Index(c *gin.Context) {
	c.File("templates/index.html")
}

func PostImage(c *gin.Context) {
	db := config.InitDb()
	defer db.Close()

	img, err := imageupload.Process(c.Request, "file")
	r.Table("photos").Insert(img).Run(db)
	c.JSON(201, gin.H{"success": img})
	
	if err != nil {
		panic(err)
	}

	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

	if err != nil {
		panic(err)
	}
	
	thumb.Save(fmt.Sprintf(Static_url + "%d.png", time.Now().Unix()))
	fmt.Println(Static_url)
	c.Redirect(http.StatusMovedPermanently, "/api/index")
	}
