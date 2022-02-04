# Channels en GO (Diálogo entre GORoutines)

Un canal es un mecanismo de comunicación que le permite a una Gorutine enviar información a otra Gorutine a través de él. Cada canal tiene un tipo especial, que es el tipo de datos que los canales pueden enviar. Un canal que puede enviar datos de tipo int generalmente se escribe como ```chan int```.

Declarar un canal es muy simple, podemos usar la palabra clave chan. Además, también debemos especificar el tipo de datos enviados y recibidos en el canal para que podamos saber qué tipo de datos enviar al canal y también saber de este canal. Qué tipo de datos se pueden recibir aquí.
```
canal := make(chan int)
```
Aquí inicializamos un canal de tipo ```chan int```, por lo que solo podemos enviar datos int a este canal, por supuesto, el receptor solo puede resibir datos de tipo int.

Un canal tiene dos operaciones principales: enviar y recibir datos. Una instrucción de envío envía un valor de una rutina a través de un canal a otra rutina que realiza una operación de recepción. Tanto las operaciones de envío como de recepción usan el operator <-. 
```
ch <- 2 //Envía el valor 2 al canal ch
x := <- ch //Lee el valor del canal ch y asigna el valor de lectura a la variable x
<- ch //Lee el valor del canal ch y no lo guarda en ninguna variable
```
