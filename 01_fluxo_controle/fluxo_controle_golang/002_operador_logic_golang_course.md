# Controle fluxo Operadores logico


<p>
 Os operadores logico server para quando precisamos ter duas decisoes,com isso em mente podemos usar operadores logico and(&&) e or(||)
</p>
<p>
  Com isso em mente,os operadores vem para facilitar e simplementes a gente poder fazer mais de uma operação dentro do if !!
</p>


## Exemplo de como utilizar operadores logico :

```

func main(){
  salario := 5001.00


  // operador Or
  if salario < 1000.00 || salario > 5000.00 {
    // executa esse bloco 
    salario = salario - (salario * 0.08)
  }

  // operador and

    if salario < 1000.00 && salario > 5000.00 {
    // executa esse bloco 
    salario = salario - (salario * 0.08)
  }
}
```
<p>
  o operador or ele e || onde uma das condiçoes for true ele vai executar abaixo codigo 
</p>  
<p>
  jA o operador and ele e && ondeduas das condiçoes for true ele vai executar abaixo codigo 
</p>  


## Outro exemplo: 


```
// utilizando operador logic And :

func main (){

  salario := 1000.00 
  tipo := "CLT"
  

  if salario < 1000.00 && tipo == "CLT" {
    salario =  salario - (salario * 0.08)
  }else if salario 1000.00 && tipo == "PJ" {
    salario = salario - (salario * 0.05)
  }
}
```