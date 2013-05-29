# bloom-as-a-service

## Instructions

It's simple: clone, set the GOPATH and run the service.

    $ go get github.com/cabello/bloom-as-a-service
    $ cd $GOPATH/github.com/cabello/bloom-as-a-service
    $ go run server/server.go

The filters you create and the entries you add will be held on memory.

## API

### Create bloom filter

Endpoint: `POST /v1/filters`

Parameters:

- `name`, the name you gonna use to create and check for entries, ex: spammers
- `capacity`, how many records do you plan to insert on bloom filter, ex: 100000
- `errorRate`, the chance of the filter making a mistake when checking for an entry, ex: 0.01 (1%)

Returns:

- `201 Created`, filter was created
- `409 Conflict`, filter already exists

### Retrieve bloom filter information

Endpoint: `GET /v1/filters/{filterName}`

Parameters:

- `filterName` (on URL), the name you used to create the filter

Returns:

- `200 OK`, filter was found, JSON with `capacity` and `errorRate`
- `404 Not Found`, filter not found

### Delete bloom filter

Endpoint: `DELETE /v1/filters/{filterName}`

Parameters:

- `filterName` (on URL), the name you used to create the filter

Returns:

- `204 No Content`, filter was deleted
- `404 Not Found`, filter not found

### Create entry on bloom filter

Endpoint: `POST /v1/filters/{filterName}/entries`

Parameters:

- `name`, the name you gonna use to check if the entry is on your bloom filter

Returns:

- `201 Created`, entry was created
- `400 Bad Request`, filter not found


### Check entry existence on bloom filter

Parameters: `GET /v1/filters/{filterName}/entries/{entryName}`

- `filterName` (on URL), the name you used to create the filter
- `entryName` (on URL), the name you use to create the entry on the bloom filter

Returns:

- `200 OK`, entry exists
- `404 Not found`, entry not found
- `400 Bad Request`, filter not found


