# bloom-as-a-service

## API

### `POST /v1/filters`

Parameters: 

- name
- capacity
- errorRate

Returns:

- 201 Created
- 409 Conflict

### `GET /v1/filters/{filterName}`

Parameters: 

- None

Returns:

- 200 OK, JSON with `capacity` and `errorRate`
- 404 Not Found

### `DELETE /v1/filters/{filterName}`

Parameters:

- None

Returns:

- 204 No Content
- 404 Not Found


### `POST /v1/filters/{filterName}/entries`

Parameters:

- name

Returns:

- 201 Created
- 400 Bad Request


### `GET /v1/filters/{filterName}/entries/{entryName}`

Parameters:

- None

Returns:

- 200 OK
- 404 Not found
- 400 Bad Request


