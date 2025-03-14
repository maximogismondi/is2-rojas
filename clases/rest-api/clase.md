# REST API

## Status codes

Los status codes son códigos que se envían en la respuesta de una petición HTTP para indicar el estado de la misma. Los grupos de status codes son:

- **1xx: Información**: Permite al servidor indicar que la petición ha sido recibida y está siendo procesada. Muy raramente se utiliza.
- **2xx: Éxito**: Indica que la petición fue recibida, entendida y aceptada.
- **3xx: Redirección**: Indica que el cliente necesita realizar más acciones para completar la petición.
- **4xx: Error del cliente**: Indica que hubo un error en la petición del cliente.
- **5xx: Error del servidor**: Indica que hubo un error en el servidor.

Algunos status codes comunes son:

- **200 OK**: La petición fue exitosa.
- **201 Created**: La petición fue exitosa y se creó un nuevo recurso.
- **204 No Content**: La petición fue exitosa pero no hay contenido para devolver.
- **209 Conflict**: La petición no pudo ser completada debido a un conflicto.
- **400 Bad Request**: La petición no pudo ser procesada debido a un error del cliente.
- **401 Unauthorized**: El cliente no está autorizado para acceder al recurso.
- **403 Forbidden**: El cliente no tiene permisos para acceder al recurso.
- **404 Not Found**: El recurso no fue encontrado.
- **500 Internal Server Error**: Error interno del servidor.

## Arquitectura Layer Design

- **Controller**: leer y parsear el body
- **Service**: toda la lógica de negocio (potencialmente hablar con una base de datos o algo así)
- **Router (OPCIONAL)**: manejar las rutas y llamar a los métodos del controller. Podría ser parte del controller directamente.
- **Model**: estructura de datos (puede ser una base de datos, un archivo, etc.)
