package main
import "fmt"


type Entry struct {
    chave int
    valor interface{}
}


type Map struct {
    entradas []Entry
}


func main() {
    var entrada Entry
    for true {
        fmt.Println("entrada (-1 sai): ")
        fmt.Scanf("%v", &entrada.chave)
        if (entrada.chave == -1) {
            break
        } else {
            fmt.Println("saída: ", entrada.chave)
        }
    }
    mapaTeste := []Entry{
        Entry{
            chave: 1,
            valor: "Legal",
        },
        Entry{
            chave: 0,
            valor: "chato",
        },
    }
    for _, i := range mapaTeste {
        fmt.Println(i.chave, i.valor)
    }
}



