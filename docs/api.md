# API DOCS

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
`GET /api/link/{link:str}`

Returns: `200` `500`
Body: JSON

Get the metadata for a single link

### Create/Update Link
`POST /api/link/{link:str}?{param}`
Returns: `200` `400` `500`

Request Body Structure: 
```json
{
    "target_url": str, // Redirect URL
    "owner": str, // User/Group that owns link
    "description": str // Short Description of purpose
}
```
Request Parameters: 
|  Key + Value | Value  |
|---|---|
|  force=true | force  |
|   |   |
|   |   |

Note: In this 

## Update Link
`PUT /{link:str}`


Returns: `200` `500`

Updates a link 

## Delete Link
`DELETE /{link:str}`


Returns: `200` `500`

Updates a link 
