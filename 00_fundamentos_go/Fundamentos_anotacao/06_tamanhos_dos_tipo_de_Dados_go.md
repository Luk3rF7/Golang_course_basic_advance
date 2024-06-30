# Tamanho dos tipo de dados em Go lang

<p>
  Os tipo de inteiro ,em qualquer lang vai ter um tamanho maximo que podemos atribuir, onde execmplo so e permitido colocar 127
</p>

## Tipos de tamanho

<li> int8    8bit-signed integer
<li> int16   16-bit signed integer 
<li> int32   32-bit signed integer
<li> int64   64-bit singed integer

<li> uint8   8-bit unsinged integer
<li> uint16  16-bit unsinged integer
<li> uint32  32-bit unsinged integer
<li> uint64  64-bit unsinged integer

<li> int  both in and unit contain same size, either 32 or 64 bit

<li> uint  both in and unit contain same size, either 32 or 64 bit

<li>float32 32-bit IEE 754 floating-point number 
<li>float64 64-bit IEE 754 floating-point number

<p>
  cada um dele tem tamanho maximo que podemos atribuir seja 8,16,32 a 64 bit onde existe alguma diferenças:
</p>
<p>
   int8 e diferente do uint8 por um e assinado e outro não, por exemplo :

    - com exemplo abaixo unit e valor não assinado
    - e não podemos atribuir valores negativos


    ```

    criando var uint8
    var num uint8;
    num = -1 // (x) o uint não pode atribuir valores negativos
    num = 1 // vai atribuir e ler normal

    ```

    - ja com int que e valor assinado conseguimos
        passar valores negativo :


    ```
    criando var com int8 :

    var numberint8;
    number = -1
    ```

 </p>

| Tipo   | Faixa Valores                                              | Tamanho |
| ------ | ---------------------------------------------------------- | ------- |
| int8   | -127 a 127                                                 | 8-bits  |
| uint8  | 0 a 255                                                    | 8-bits  |
| int16  | -32.768 a32.767                                            | 16-bits |
| uint16 | 0 a 65.535                                                 | 16-bits |
| int32  | -2.147.483.648 a 2.147.483.647                             | 32-bits |
| uint32 | 0 a 4.294.967.295                                          | 32-bits |
| int64  | t64 -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807 | 64-bits |
| uint64 | 4 0 a 18.446.744.073.709.551.615                           | 64-bits |
