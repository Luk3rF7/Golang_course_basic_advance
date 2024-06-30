# Operadores em Go lang

<p>
  Assim como em outra lang temos operadores onde podemos fazer ações aritmetricas etc..
</p>

```
// exemplo
package main

import ("fml","reflect")

func main() {
  num1 := 10
  num2 := 20

  fml.Println(num1 + num2) // = 30
}
```

## operadores que tem :

<li> adicao / concatenação de string -> +
<li> subtracao -> - 
<li> multiplicação -> *
<li> divisão -> /

<p>
  Para saber um tipo de variavel ou dados pode usar a biblioteca reflect
</p>

```
import ("fmt""reflect")

func main(){
  num1 := 10.0
  num2 := 20.0

  result:= num1 / num2

  fmt.Println(result)
  fmt.Println(reflect.TypeOf(result))


}

```
