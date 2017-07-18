# MicroservicePokemon
A simple example for a microservice in Golang

### Estructura

Nuestra arquitectura MVC esta compuesta de la siguiente manera:
[![N|Solid](https://cldup.com/RfII1ZbtON.png)](https://nodesource.com/products/nsolid)
 - **pokemon**: Es el nombre del repositorio del microservicio.
 - **assets**: Dentro de este folder se encuentra la carpeta *images*, aqui estarán las carpetas que contienen las fotos!. Hay un nivel mas de carpeta con el nombre *pokemons* este nombre varia dependiendo del la relación que lleven las imagenes. **Para los microservicios esta ruta estará en AWS S3, sólo como referencia se dejo en esta prueba !:**
 - **config**: Esta carpeta contiene a *database.go* el cual realiza la coneccion a la base de datos.
 - **controllers**: La carpeta contiene el archivo *pokemon.go* en donde se tiene la lógica de programación.
 - **models**: En esta carpeta se encuentran el archivo *pokemon.go* el cual contiene los structs que serán consultados a la base de datos.
 - **main.go**: Es nuestro archivo principal, el cual se encuentra en el nivel principal del repositorio, *main.go* es el archivo que se compila en la terminal !.
 - **pokemon_swagger.json**: Adicionalmente, en el nivel principal del repo se encuentra el archivo *pokemon_swagger.json*, este sirve para diseñar el API para montarlo en WSO2. swagger se codea con YAML para generar un JSON.
