# Variavel Em Go lang

## O que é Variavel ?

<p>
  A Variavel e um lugar para poder armazenar um tipo de dado / valor e eles nos ajuda  a cria nosso programa no
  go lang temos conceitos:
</p>

<li> var -> Palavra reservada que serve para criar variavel

### Exemplo :

```
package main

import 'fmt'

func main (){

   // sintaxe: -> var nomeDaVar  tipoDaVar
   // criando variavel
  var peso int ;
  // atribuição
  peso = 10;
   fmt.Println(peso)
}

```

<p>
  As variavel elas geralmente tem que ter um tipo ,e existe tipos especifico de cada lang ,no go lang usamos int,float,string,boolean e varios outro tipos que podemos criando e depende do que programa precisa
</p>

```
package main

import 'fmt'

func main (){
// criar e tipa
var peso int ;
var message string;

// atribuição
peso = 10;
message = "isso e uma mensagem!"
fmt.Println(peso)
}

```

<p>
 O go lang tbm pode fazer criação de variavel com atribuição usando := como no exemplo:
</p>

```
package main

import 'fmt'

func main (){
    mensagem := "criação var + atribuição"
   fmt.Println(peso)
}

```

<p>
  := Com isso não precisamos passar  o tipo da variavel!
</p>

## Não podemos fazer em go

<li> utilizar palavras reservada como nome de variaveis/funcao
