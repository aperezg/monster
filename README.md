# Monster API
The Monster API, is a simple CRUD API for development purpose, this API can be used to create a demo projects, katas, etc...

The API request/response is based on [JSON:API](https://jsonapi.org/) specification.

## How can I use it?

### Install 

```sh
$ go get -u github.com/aperezg/monster/cmd/monster
```

### Usage 

Launch server with predefined data
```sh
$ monster --withData
```

Launch server with custom host and port
```sh
$ monster --port 8080 --host monster.io
```

## Endpoints

Create a new monster

```
POST /monsters
```

Update a monster

```
PATCH /monsters/{monster_id}
```

Delete a monster

```
DELETE /monsters/{monster_id}
```

Fetch all monsters

```
GET /monsters
```

Fetch a monster by ID
```sh
GET /monsters/{monster_id}
```

## Contributing

If you think that you can improve with new endpoints, and functionallities the API feel free to contribute with this project with
fork this repo and send your Pull Request.

## License
MIT License, see [LICENSE](https://github.com/aperezg/monster/blob/master/LICENSE)