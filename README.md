# Театральная база знаний

## **Запуск:**

Пример `.env` файла:

```
POSTGRES_USER=test
POSTGRES_PASSWORD=test
POSTGRES_DB=test
POSTGRES_HOST=database
POSTGRES_PORT=5432
```

Если хотите запускать не из докера, то нужно поменять хост базы:

```
POSTGRES_HOST=localhost
```

Сначала устанавливаем loki плагин для докера:
```bash
docker plugin install grafana/loki-docker-driver:latest --alias loki   --grant-all-permissions
```

Создаем или редактируем `daemon.json`, который находится в `/etc/docker`, либо `C:/ProgramData/docker`:
```json
{
   "log-driver": "loki",
   "log-opts": {
     "loki-url": "http://localhost:3100/loki/api/v1/push",
     "loki-batch-size": "400"
   }
}
```

Перезапускаем докер:
```bash
sudo service docker restart # sudo systemctl restart docker
```

Запустить:
```
docker-compose up --build
```

## **Пути:**
```bash
http://localhost:3000 # grafana
http://localhost:16686 # jaeger
http://localhost:9090 # prometheus
http://localhost:8000 # сам сервис с документацией
```
