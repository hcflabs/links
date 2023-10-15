# API DOCS

## Model
### POST Link
```json
{
    "target_url": str, // Redirect URL
    // "owner": str, // User/Group that owns link
    "description": str // Short Description of purpose
}
```
### API Link
```json
{
    "target_url": str, // Redirect URL
    "owner": str, // User/Group that owns link
    "description": str, // Short Description of purpose
    "creation_time": int, // timestamp
    "last_modified": int, // timestamp
}
```

## User Interface

Here I will document the basic functions and features of the service.

### Redirect User
 `GET /{link:str}?{wildcard params}`

Returns: `301` `302` `404` `500`

Redirects a browser/request to a different configured location.

### Enter Admin Portal
`GET /admin`
Returns: `200` `500`
Body: HTML

This is a rendered admin portal for users to search, edit and update links.

## Programmable Interface
### Get Link Metadata
`GET /api/links/{link:str}`

Returns: `200` `500`
Body: JSON

Get the metadata for a single link

### Get Links Paginated
`GET /api/links?{params}`

Returns: `200` `500`
Response Body Structure: [[]Link](#link)


Request Parameters: 
|  Key + Value | Description  |
|---|---|
|  pagesize=X | force  |
|   offset=Y|   |
|   |   |

### Create/Update Link
`POST /api/links/{link:str}?{param}`
Returns: `200` `400` `500`

Request Body Structure: [Link](#link)
Response Body Structure: [Link](#link)

Request Parameters: 
|  Key + Value | Description  |
|---|---|
|  force=true | force  |
|   |   |
|   |   |

Note: In this 

## Delete Link
`DELETE /api/links/{link:str}`


Returns: `200` `500`

Updates a link 
