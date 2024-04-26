package main

 

import (

    "database/sql"
    "fmt"
    "html/template"
    "net/http"
    _ "github.com/denisenkom/go-mssqldb"

)

 

type Produto struct {

    Nome       string
    Descricao  string
    Preco      float32
    Quantidade int

}

 

var temp = template.Must(template.ParseGlob("templates/*.html"))

 

// func main() {

//  http.HandleFunc("/", index)

//  http.ListenAndServe(":8080", nil)

// }

 

func conectaComBancoDeDados() {

    // Configurar a conexão com o banco de dados *sql.DB

    server := ""
    port := 1433
    user := ""
    password := ""
    database := ""

 

    // String de conexão

    connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",

        server, user, password, port, database)

 

    // Abrir a conexão com o banco de dados

    db, err := sql.Open("sqlserver", connectionString)

    if err != nil {

        fmt.Println("Erro ao conectar ao banco de dados:", err.Error())

        return

    }

    defer db.Close()

 

    // Verificar se a conexão está ativa

    err = db.Ping()

    if err != nil {

        fmt.Println("Erro ao pingar o banco de dados:", err.Error())

        return

    }

 

    fmt.Println("Conexão com o banco de dados bem-sucedida!")

 

}

 

func main() {

 

    http.HandleFunc("/", index)
    http.ListenAndServe(":8080", nil)

}

 

func index(w http.ResponseWriter, r *http.Request) {

 

    db := conectaComBancoDeDados()


    selectDeTodosOsProdutos, err := db.Query("select * from produtos")

    if err != nil {
        panic(err.Error())
    }

 

    p := Produto{}
    produtos := []Produto{}

 

    for selectDeTodosOsProdutos.Next() {

        var id, quantidade int

        var nome, descricao string

        var preco float64

 

        err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

        if err != nil {

            panic(err.Error())

        }

 

        p.Nome = nome

        p.Descricao = descricao

        p.Preco = preco

        p.Quantidade = quantidade

 

        produtos = append(produtos, p)

    }

 

    temp.ExecuteTemplate(w, "Index", produto)

}