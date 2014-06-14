# Boolr

_Finally_, a web-scale boolean-switching service. Informally created for
[SOHacks](http://sohacks.com/).

## Endpoints

* `POST /users`
* `POST /users/login`
* `GET /users/:id`
* `POST /booleans`
* `PUT /booleans/:id/switch`
  * `url` - url to issue callback to
  * `limit` - number of times to switch boolean, default `-1` (unlimited)
* `DELETE /booleans/:id/switch`
* `DELETE /booleans/:id`

## License

See the [LICENSE](LICENSE) file.
