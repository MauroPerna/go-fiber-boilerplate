# Go Fiber Backend

Este proyecto es una aplicación backend desarrollada en Go, que utiliza el framework Fiber para crear un entorno web rápido y eficiente. Siguiendo los principios de Clean Architecture, el proyecto está diseñado de manera modular, lo que garantiza un código limpio, fácilmente mantenible y extensible. MongoDB se utiliza como la base de datos principal, proporcionando un almacenamiento persistente y escalable.

## Características

- **Fiber**: Un framework web rápido y minimalista inspirado en Express.
- **MongoDB**: Base de datos NoSQL utilizada para almacenamiento persistente.
- **Fx**: Librería para la inyección de dependencias y manejo del ciclo de vida de la aplicación.
- **Modularidad**: El código está dividido en módulos, lo que facilita la escalabilidad y mantenibilidad.
- **Middleware**: Autenticación y registro de peticiones implementados como middleware.

## Estructura del Proyecto

La estructura del proyecto ha sido ampliamente influenciado por la forma de hacer las cosas de NestJS, un framework de Node.js que utiliza TypeScript como lenguaje de programación y esta inspirado en Clean Architecture. Utiliza Fiber como framework web, que es un framework web rápido y minimalista inspirado en Express.
El proyecto sigue una estructura modular que separa claramente las responsabilidades dentro de la aplicación. Cada módulo se encarga de un aspecto específico, promoviendo la cohesión interna y la independencia entre componentes. A continuación, se ofrece una visión general de la estructura del proyecto:

### Directorios

- `/common`: Contiene archivos comunes a todos los módulos. Como enums, middlewares, etc.
- `/domain`: Contiene los archivos de dominio, el core de la aplicación, la logica de negocio.
- `/services`: Contiene los servicios de la aplicación, como el servicio de autenticación, el servicio de base de datos, etc. Lo que en Clean Architecture se denomina "Infraestructura".


## Prerrequisitos

- Go 1.19+
- MongoDB
- [Air](https://github.com/cosmtrek/air) (Opcional para desarrollo)

## Instalación

1. Clona el repositorio:

```bash
    git clone https://github.com/MauroPerna/go-fiber-boilerplate.git
    cd go-fiber-boilerplate
    go mod tidy
```

2. Ejecuta el servidor:

### Desarrollo

```bash
    air
```

### Producción

```bash
    go run main.go
```

3. Abre tu navegador y accede a `http://localhost:8080`.

## Contribuir

Si deseas contribuir al proyecto, puedes hacerlo de varias maneras:

- Reportando errores o sugiriendo mejoras en la documentación.
- Abriendo un Issue en el repositorio.
- Haciendo Pull Requests para mejorar el código.
- Si tienes experiencia en Go, puedes ayudar a mejorar el código.


