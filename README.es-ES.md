agollo_demo - Demostración de Apollo
================

Demostración del uso de [agollo](https://github.com/apolloconfig/agollo).

Está diseñada principalmente para verificar la función de actualización automática del cliente de Agollo.

Dependencias
------------

## Uso con go get

```
go mod
```

## Uso con gopm (recomendado)

```
gopm get github.com/cihub/seelog -v -g
gopm get github.com/apolloconfig/agollo -v -g
```

Se recomienda utilizar [gopm](https://github.com/gpmgo/gopm), que no requiere conexión a través de un proxy.

Instalación
------------

Si aún no tiene un entorno de desarrollo Go instalado, consulte el siguiente documento [Getting Started](http://golang.org/doc/install.html) para instalarlo. Una vez completada la instalación, ejecute el siguiente comando:

## Uso

* [Versión web](web)
* [Hola mundo](helloworld)
* [Uso personalizado](custom)
    * [Componente de registro personalizado](custom/log)
    * [Componente de caché personalizado](custom/cache)
    * [Escuchar eventos de cambio](custom/listener)

## Métodos de inicialización
### Crear aplicación
ID de aplicación: agollo-test
Nombre de la aplicación: agollo-test
#### Crear espacio de nombres privado

##### testjson.json
```json
{
  "a":"b",
  "c":"d"
}
```
##### yml
```yml
application:
    name: agollo-test
    port: 8000
```

### Crear aplicación cifrada
ID de aplicación: agollo-secret-test
Nombre de la aplicación: agollo-secret-test

##### application
```
connect: 10
quantity: 5
```

##### Activar la clave
