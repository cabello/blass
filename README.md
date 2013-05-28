# bloom-as-a-service

## API

### `POST /v1/filters`

Parameters: 

- name, the name you gonna use to create and check for entries, ex: spammers
- capacity, how many records do you plan to insert on bloom filter, ex: 100000
- errorRate, the chance of the filter making a mistake when checking for an entry, ex: 0.01 (1%)

Returns:

- 201 Created
- 409 Conflict

### `GET /v1/filters/{filterName}`

Parameters: 

- filterName (on URL), the name you used to create the filter

Returns:

- 200 OK, JSON with `capacity` and `errorRate`
- 404 Not Found

### `DELETE /v1/filters/{filterName}`

Parameters:

- filterName (on URL), the name you used to create the filter

Returns:

- 204 No Content
- 404 Not Found


### `POST /v1/filters/{filterName}/entries`

Parameters:

- name, the name you gonna use to check if the entry is on your bloom filter

Returns:

- 201 Created
- 400 Bad Request


### `GET /v1/filters/{filterName}/entries/{entryName}`

Parameters:

- filterName (on URL), the name you used to create the filter
- entryName (on URL), the name you use to create the entry on the bloom filter

Returns:

- 200 OK
- 404 Not found
- 400 Bad Request


