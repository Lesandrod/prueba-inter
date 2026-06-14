# prueba-inter

Dos APIs que trabajan juntas para factorizar matrices y calcular estadísticas.

## Cómo funciona

1. La **API en Go** recibe una matriz, calcula la factorización QR y le pasa los resultados a la API de Node
2. La **API en Node** recibe las matrices Q y R y calcula estadísticas (máximo, mínimo, promedio, suma, si es diagonal)
3. La API de Go devuelve todo junto al cliente

## Credenciales

```
usuario: admin123
contraseña: prueba123
```

## Endpoints

### Node API — `http://apipruebanodev1.137.184.19.146.sslip.io`

| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /login | Obtiene el token JWT |
| POST | /stats | Calcula estadísticas sobre Q y R |

### Go API — `http://apipruebagov1.137.184.19.146.sslip.io`

| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /qr | Factoriza la matriz y devuelve Q, R y estadísticas |
| POST | /qr/decompose | Solo devuelve Q y R sin estadísticas |

## Ejemplo de uso

**1. Login**
```bash
POST /login
{
  "username": "admin123",
  "password": "prueba123"
}
```

**2. Factorizar**
```bash
POST /qr
Authorization: Bearer <token>
{
  "matrix": [[1,2],[3,4],[5,6]]
}
```

## Correr con Docker

```bash
docker compose up --build
```

## Stack

- Go + Fiber
- Node.js + Express
- JWT para autenticación
- Docker + Docker Compose
