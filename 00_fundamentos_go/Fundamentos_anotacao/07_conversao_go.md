# Conversões de dados em Go lang

<p>
  Construimos sistema para ser integrado não só com outro 
  sistema mas outros usuario se comunica e vem tipo de dados 
  diferente exemplo:
</p>

```
package main
import "fmt"

func main (){
  var num int8 = 127
  // as vezes essa var vai receber mais que 127 var
  // entao criamos segunda variavel onde vai fazer conversão
  //  numInt = int(num)
  var numInt int;

  // não sera atribuira a conversão :
    numInt = num

  //sera atribuido a conversão:
    numInt = int(num)

  fmt.Println(num)
}
```

<p>
  conversão nada mais é que eu transformar aquele valor para 
  aquele tipo de dado que estou esperando e ao caso contrario
</p>

## conversoes numericas

 <p>
  Com pacote strconv conseguimos ter conversoes de tipo comun 
  como string para int usando Atoi e int para string com Itoa
 </p>

```
err := strconv.Atoi("-42") => -42 int
s := strconv.Itoa(-42) => "-42" string
```

## usando funções Parse

- b,err := strconv.ParseBool("valor boolean")
- f,err := strconv.ParseFloat("valor string",bit) bit 32/64
- i,err := strconv.ParseInt("valor",base,bit) int bit ->8,16,32,64
- u,err := strconv.ParseUint("valor",base,bit) uint bit -> 8,16,32,64

### Sintaxe code

```

  b, err := strconv.ParseBool();

  o b e uma variavel para armazenar valor convertido
  o err => erro caso conversor encontre
  o strconv -> pacote de conversão
  o ParseBool -> metodo/fn do pacote onde fara conversão do tipo recebendo argumento como string

  nos Parse int e uint são 3 argumento uma dele e valor string , base e bit que serão convertido

  nos Parse float ja recebe 2 argumentos um valor e bit de conversão a ser realizado
```
