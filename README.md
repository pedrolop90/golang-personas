# golang-personas
Servidor web usando golang y chi para el manejo de controladores

### Ejecución
Podemos Ejecutar el proyecto de dos formar:
- Docker: Ya está preparado un
  docker-compose para ejecutarlo con el comando
  **docker-compose -d up**.
- Gradle: usando el comando **gradle run**.
- Release: Esta preprado un release, para descargar y
  ejecutarlo **go_personas.exe**

### Pruebas

- #### Dependencias
  Para poder probar el funcionamiento correcto 
  debemos tener un navegador.

- #### Guia de Pruebas
    1.  Obtener Personas:
        Probaremos ingresando a la siguiente url
        *<http://localhost:3000/personas/>* 
        desde el navegador, la salida seria
        muy similar a la siguiente:
        ```text
        Hola mundo!!!
        ```

Como comprobamos pudimos obtener el mensaje
"Hola mundo!!!" en el puerto que hemos definido
>> "Mientras estás pensando si deberías hacerlo o no, hay personas que lo están haciendo realidad."
