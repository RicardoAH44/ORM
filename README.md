Luis Ricardo Arredondo Huerta 191288

Ejecutar Aplicación:

go run main.go

Esto iniciará el servidor API en http://localhost:8080.

Conexión a la Base de Datos MySQL:
La aplicación se conecta a una base de datos MySQL. Asegúrese de configurar la cadena de conexión en el archivo config.yml con los detalles de la base de datos.

Operaciones CRUD:
La API proporciona las siguientes rutas para operaciones CRUD:

GET /products: Obtiene todos los productos.
GET /products/:id: Obtiene un producto por ID.
POST /products: Crea un nuevo producto.
PUT /products/:id: Actualiza un producto existente.
DELETE /products/:id: Elimina un producto por ID.

La aplicación utiliza GORM como ORM y depende de las bibliotecas 
gorm.io/gorm y gorm.io/driver/mysql 