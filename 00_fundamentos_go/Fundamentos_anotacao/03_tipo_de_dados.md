# Tipo de Dados em Go lang

<p> 
  No Go lang os tipos são estaticos uma vez declarada não podemos mudar o tipo
</p>

```
// se definimos variavel como string nao Conseguimos passar um int,boolean,float nela

 var texto string;

 texto = 10 // (x)  go vai ler e retorna erro de tipagem
 texto = "texto" // (V) go ira ler texto sem retorna erro
 texto = "texto atualizado" // (V) go ira ler texto normal sem retorna erro
```

<li> Com double // conseguimos cometra em go lang

### Tipos Primitivos em Go lang

<li>string -> Texto ou characteres
<li> int -> numero inteiro 1,10,20
<li> float32 -> numero de ponto fluuantes 10.50
<li> bool -> valores booleanos que são true ou falso.
