# Variavel const em Go lang

<p>
  A ideia de criar uma  variavel constante (const) ela ser constante  onde seu valor nunca pode ser mudado ,com isso garantimos e evitamos que valor ser alterado !
 </p>

```
package main

import "fmt"

func main(){
  const taxa = 10
  taxa = 5
  fmt.Println(taxa)

}
```

<p>
  Com exemplo acima o valor taxa = 5 não vai ocorrer e retorna erro de declaração, com isso evitamos alterar valor
</p>
