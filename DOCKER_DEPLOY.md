Instrucciones rápidas para Docker y despliegue

1) Construir y levantar localmente

En la carpeta raíz (donde está `docker-compose.yml`):

```bash
# build y levantar
docker compose build
docker compose up -d

# ver logs
docker compose logs -f go-api
```

2) Probar endpoints

- Login (Node):
```bash
curl -s -X POST http://localhost:3000/login -H "Content-Type: application/json" -d '{"username":"admin123","password":"prueba123"}'
```

- QR (Go):
```bash
curl -s -X POST http://localhost:8080/qr -H "Authorization: Bearer <TOKEN>" -H "Content-Type: application/json" -d '{"matrix":[[1,2],[3,4]]}'
```

3) Desplegar en VPS (ejemplo mínimo)

- En tu VPS instalar Docker y Docker Compose. Luego copiar este repo y en la carpeta raíz ejecutar:

```bash
docker compose build
docker compose up -d
```

El VPS expondrá los puertos 3000 y 8080.

4) Opciones de hosting gratuito

- Fly.io: tiene un workflow con `flyctl` que puede desplegar contenedores (lee su doc y crea una app por servicio).
- Render / Railway / Render free tier: pueden ejecutar Docker images o repositorios directamente.

Si quieres, puedo:
- Crear `fly.toml` y comandos para `flyctl`.
- Subir imágenes a Docker Hub y darte los comandos `docker push`.
- Preparar un `Dockerfile` más pequeño o multi-service para un solo contenedor (no recomendado).

Dime qué prefieres y lo preparo.
