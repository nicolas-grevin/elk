# ELK

This project was inspired by [deviantony/docker-elk](https://github.com/deviantony/docker-elk/blob/release-7.x/kibana/Dockerfile)

## Start project

```bash
docker compose up setup
docker compose up -d
```

## Start Apps

```bash
docker compose exec -d app_1 generate-logs --name=app_1 --format=json --output=file --interval=1
docker compose exec -d app_2 generate-logs --name=app_2 --format=text --output=file --interval=1
```

## Kibana interface

http://localhost:5601
