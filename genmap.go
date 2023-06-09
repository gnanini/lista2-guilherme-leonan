package main
import (
    "fmt"
    "strings" //precisou pra ser case insensitive, senão iria ter que fazer a minha versão, sqn!!
)


type Entry struct {
    chave int
    valor interface{}
}


// aqui não sei se era exatamente isso, mas funcionou
type Map struct {
    entradas []Entry
}


func main() {
    // variáveis que serão usadas, ignore os nomes horríveis
    var comando string
    var chave int
    var valor string
    var mapa Map

    // aqui é o terminal iterativo, quase que merecia a sua própria função
    // usei um switch -> case pra pegar as abreviações
    for {
        fmt.Printf("> ") // prompt de comando
        fmt.Scanln(&comando, &chave, &valor) // o scanln caiu bem para pegar as entradas
        if (comando == "exit") {
            break
        } else {
            comando = strings.ToLower(comando)
            switch {
                case comando == "add":
                    addEntry(chave, valor, &mapa)
                case comando == "get":
                    getEntry(chave, mapa)
                case comando == "size":
                    mapSize(mapa)
                case comando == "print":
                    printMap(mapa)
                default:
                    fmt.Println("Função não implementada!")
            }
        }
    }
}


// usei um ponteiro pra map, assim consegui passar o novo valor do mapa de maneira "suja"
func addEntry(chave int, valor string, mapa *Map) {
    mapa.entradas = append(mapa.entradas, Entry{chave: chave, valor: valor})
    // pra imprimir na formatação tive que usar um for, nada demais!
    for _,j := range mapa.entradas {
        fmt.Printf("[%d] \"%s\"\n", j.chave, j.valor)
    }
    return 
}


// Essa função itera entre todos os valores do slice e busca o índice
// caso não encontre o índice retorna chave inválida
func getEntry(chave int, mapa Map) {
    // ignorei o índice do for
    for _,j := range mapa.entradas {
        if j.chave == chave {
            fmt.Printf("[%d] \"%s\"\n", j.chave, j.valor)
            return
        }
    }
    fmt.Println("Chave inválida!")
}


// só imprime o length do slice dentro de Map
func mapSize(mapa Map) {
    fmt.Println(len(mapa.entradas))
}


// nada de mais aqui
// Exibe na saída padrão o número de entradas atualmente armazenadas no mapa.
func printMap(mapa Map) {
    fmt.Println(mapa.entradas)
    fmt.Println(len(mapa.entradas))
}


// PS: Professor, compilei o programa usando o gccgo (pq a sintaxe se parece com a do gcc "gccgo -Wall genmap.go -o genmap") larquei até um binário aí... acho que não haverá problemas 
