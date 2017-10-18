# Angular + Roda + Sequel

## Requirements

- docker
- docker-compose

## Start app

- docker-compose up -d db
- docker-compose run app -m migrations postgres://postgres@db/recipes
- docker-compose up -d web

Open http://localhost:9292
