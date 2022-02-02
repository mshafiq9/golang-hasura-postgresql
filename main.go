package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"

    "github.com/gorilla/mux"

    "hasura-pg/database"
)

type Article struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}


func createNewArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createNewArticle")

    var article Article
    reqBody, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(reqBody, &article)

    db, err := DatabaseManager.OpenDatabase()

    sqlStatement := `INSERT INTO hasurapg.articles (id, title, description, content) VALUES ($1, $2, $3, $4)`
    _, err = db.Exec(sqlStatement, article.Id, article.Title, article.Desc, article.Content)
    if err != nil {
        panic(err)
    }

    DatabaseManager.CloseDatabase(db)

    json.NewEncoder(w).Encode(article)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")

    db, err := DatabaseManager.OpenDatabase()

    rows, err := db.Query("SELECT id, title, description, content FROM hasurapg.articles")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    var articles []Article

    for rows.Next() {
        var article Article
        err := rows.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
        articles = append(articles, article)

        if err != nil {
            panic(err)
        }
    }

    DatabaseManager.CloseDatabase(db)

    json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
    // create a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

    http.Handle("/", myRouter)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    fmt.Println("Rest API - Mux Routers - GoLang Hasura PostgreSQL")
    handleRequests()
}
