# Fluxo de Controle

<p>
  o fluxo de controle e um do mais utilizado para tomar decisão :
</p>

 ```
 func main() {
    salario := 990.00
    var salarioMaisBonus float64

    salarioMaisBonus = salario

    if salario < 1000 {
      salariomaisBonus = (salarioMaisBonus + 100)
    }

 }
 ```

 ## Existe mais maneira de voce usar if :  


 ```
 if salario > 1000 { 
  salarioMaisBonus = (salarioMaisBonus + 100)
 }

 if salario < 1000 {
  salarioMaisBonus = (salarioMaisBonus)
 }
 ```


 ### operadores existente:

<li> <   menor  
<li> >   maior
<li> <=  menor  igual
<li> >=  maior igual



## Exemplo 2:

```
func main(){
  var ehCarro bool = false
  var valorDoAutomovel = 1000

  if ehCarro {
    valorDoAutomovel = valorDoAutomovel + 55.50
  }
}
```