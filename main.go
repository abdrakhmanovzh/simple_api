package main

import (
	"encoding/json"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

type List struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Items *[]Item `json:"items"`
}

type Item struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

func getLists(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(lists)
}

func deleteList(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for index, item := range lists {
		if item.ID == id {
			lists = append(lists[:index], lists[index+1:]...)
			break
		}
	}
	json.NewEncoder(c.Writer).Encode(lists)
}

func getList(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for _, item := range lists {
		if item.ID == id {
			json.NewEncoder(c.Writer).Encode(item)
			return
		}
	}
}

func createList(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var list List
	_ = json.NewDecoder(c.Request.Body).Decode(&list)
	list.ID = strconv.Itoa(rand.Intn(100000000))
	lists = append(lists, list)
	json.NewEncoder(c.Writer).Encode(list)
}

func updateList(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for index, item := range lists {
		if item.ID == id {
			lists = append(lists[:index], lists[index+1:]...)
			var list List
			_ = json.NewDecoder(c.Request.Body).Decode(&list)
			list.ID = id
			lists = append(lists, list)
			json.NewEncoder(c.Writer).Encode(list)
			return
		}
	}
}

var lists []List

func main() {
	r := gin.Default()

	item1 := Item{ID: "52233", Content: "Monster", Done: true}
	item2 := Item{ID: "68843", Content: "Oyasumi Punpun", Done: false}
	lists = append(lists, List{ID: "2321", Title: "Favourite Manga", Items: &[]Item{item1, item2}})

	r.GET("/lists", getLists)
	r.POST("/lists", createList)
	r.GET("/lists/:id", getList)
	r.PUT("/lists/:id", updateList)
	r.DELETE("/lists/:id", deleteList)

	r.Run()
}
