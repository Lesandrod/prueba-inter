# Proyecto TÉCNICAA — go-api + node-api

Resumen rápido
- API Go: expone `POST /qr` que recibe una matriz, calcula la factorización QR (Gram–Schmidt), y envía Q/R al servicio Node para obtener estadísticas. Valida JWT en entrada y reenvía el token a Node.
- API Node: expone `POST /login` (emite JWT para demo), `POST /stats` (protegido por JWT) y `GET /health`.

Estructura
- `go-api/` : código en Go (Fiber) — servidor que calcula QR y orquesta la llamada a Node.
- `node-api/` : servidor Node (Express) — calcula estadísticas sobre matrices y devuelve resultado.

Dependencias principales
- Go: `github.com/gofiber/fiber/v2`, `github.com/golang-jwt/jwt/v5`.
- Node: `express`, `jsonwebtoken`.

Variables de entorno
- `JWT_SECRET` : secreto compartido para firmar/verificar JWT (ambos servicios). Default: `secret` (solo para desarrollo).
- `NODE_STATS_URL` : URL (incluyendo /stats) que usa Go para llamar a Node. Default: `http://localhost:3000/stats`.

Arranque rápido

1) Iniciar Node (desde `node-api`):

```powershell
cd node-api
npm install
# opcional: set JWT_SECRET en PowerShell
$env:JWT_SECRET = "mi_secreto"
npm start
```

2) Iniciar Go (desde `go-api`):

```powershell
cd go-api
go get ./...
# opcional: exportar JWT_SECRET y NODE_STATS_URL
$env:JWT_SECRET = "mi_secreto"
$env:NODE_STATS_URL = "http://localhost:3000/stats"
go run .
```

Flujo JWT y ejemplos (demo)
1. Obtener token (Node):

```bash
curl -s -X POST http://localhost:3000/login -H "Content-Type: application/json" -d '{"user":{"id":"demo"}}' | jq
# responde { "token": "ey..." }
```

2. Llamar al endpoint Go `/qr` con el token:

```bash
TOKEN=ey... # token obtenido
curl -s -X POST http://localhost:8080/qr \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"matrix": [[1,2],[3,4]]}' | jq
```

Respuesta esperada (JSON): contiene `q`, `r` y `stats` (campo `max`, `min`, `average`, etc.).

Notas de seguridad
- El secreto por defecto `secret` es solo para desarrollo. En producción configurar `JWT_SECRET` fuerte y rotarlo.
- El `/login` actual es demo: acepta cualquier payload y emite token. Implementar autenticación real antes de usar en producción.

Próximos pasos recomendados
- Añadir tests E2E que obtengan token y llamen `/qr`.
- Dockerizar ambos servicios y añadir una red con variables de entorno.
- Mejorar control de errores y políticas CORS si el front accede a las APIs.

Contacto
- Si quieres, hago la demo E2E (curl) o añado Dockerfiles y tests.
