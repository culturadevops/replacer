# Comando remplazar

Este es un comando que permite reemplazar los valores en un archivo de plantilla utilizando un archivo de leyenda.

el archivo base o aqui llamado legend puede ser igual al siguiente
```
Nombre:Juan
Lugar:xxxxx
```
el archivo plantilla o aqui llamado target puede ser igual al siguiente
```
Hola {{Nombre}}, bienvenido a {{Lugar}}.
```

el resultado seria asi 
```
Hola {{Juan}}, bienvenido a {{xxxxx}}.
```
tambien podriamos cambiar el archivo legend por los siguiente 
```
{{Nombre}}:Juan
{{Lugar}}:xxxxx
```
y el archivo final seria 

```
Hola Juan, bienvenido a xxxxx.
```
un dato mas el archivo legend tiene un separador por lineas que puede ser el que tu desees, en los ejemplos anteriores fueron ":" en el comando debe estar presente este separador para que el programa entienda que buscar.


## Uso

```bash
remplazar legend [flags]

Flags

    -h, --help: Muestra la ayuda para el comando legend.
    -l, --legend string: Especifica la ubicación del archivo de leyenda donde se encuentran los datos que se utilizarán para el reemplazo.
    -o, --output string: Especifica el nombre del archivo final que contendrá tanto los valores de la leyenda como las plantillas.
    -r, --remplace: Valor booleano que define si el target será reemplazado o si el resultado se mostrará en la pantalla.
    -s, --separated string: Especifica el separador utilizado en el archivo de leyenda para separar los datos y los valores (por ejemplo, "=").
    -t, --target string: Especifica la ubicación del archivo de plantilla donde se realizarán los reemplazos o se utilizará para crear un nuevo archivo con los datos reemplazados.
```
Ejemplos de uso

## Reemplazar en el archivo y guardar el resultado en un nuevo archivo

go run main.go puede usarse sino instala la app pero si la instala puede cambiar el nombre por  replacer


```bash

go run main.go legend -l recursos/base.vars -t recursos/archivoaRemplazar.json -s : -o archivonuevo.json

o

replacer legend -l recursos/base.vars -t recursos/archivoaRemplazar.json -s : -o archivonuevo.json

```
Este comando lee el archivo de leyenda ubicado en "recursos/base.vars" y el archivo de plantilla ubicado en "recursos/archivoaRemplazar.json". Utiliza el separador ":" y guarda el resultado en un nuevo archivo llamado "archivonuevo.json".


```bash

go run main.go legend -l recursos/base.vars -t recursos/archivoaRemplazar.json -s : -r
```
Este comando lee el archivo de leyenda ubicado en "recursos/base.vars" y el archivo de plantilla ubicado en "recursos/archivoaRemplazar.json". Utiliza el separador ":" guarda el resultado en archivo "recursos/archivoaRemplazar.json" es decir remplaza el archivo target.


```bash

go run main.go legend -l recursos/base.vars -t recursos/archivoaRemplazar.json -s :
```
Este comando lee el archivo de leyenda ubicado en "recursos/base.vars" y el archivo de plantilla ubicado en "recursos/archivoaRemplazar.json". Utiliza el separador ":" y muestra el resultado en la pantalla en lugar de guardarlos en un nuevo archivo.
Reemplazar sin guardar el resultado
