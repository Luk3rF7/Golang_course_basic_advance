# Fluxo de controle Else :


<p>
  o fluxo de controle e um do mais utilizado para tomar decisão :
</p>


```

// quando utilizamos !em uma condição estamo negando ela 

if ehCarro {
  valorDoCarro += 55.50
}
if !ehCarro {
    valorDoCarro += 20.50
}
```


<p>
  com controle else podemos criar seguinte jeito :
</p>

```

// quando utilizamos else uma condição onde se nao for a primeira vai para proximo

if ehCarro {
  valorDoCarro += 55.50
}
else {
    valorDoCarro += 20.50
}

```

## Podemos usar tambem o else if : 

```
func main(){
  salario := 1000.00

  if salario < 1000.00 {
     salario = salario - (salario * 0.08)
  } else if salario <= 1200.00{
    salario = salario - (salario * 0.10)
  }else {
    salario = salario - ( salario * 0.15)
  }
}
```