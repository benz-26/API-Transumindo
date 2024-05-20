# API-Transumindo
## Train Station Location API

A simple API to check if a user has arrived at a specified train station, airport, bus station, or harbour in Indonesia.

## Important Note
*This API is created and deployed for educational purposes only*.

## Endpoints

### 1. Get Stations

**Endpoint:** `/stations`  
**Method:** `GET`  
**Description:** Retrieves a list of train stations, airports, bus stations, and harbours in Indonesia. This endpoint returns detailed information about various stations, including their name, latitude, longitude, and type.#### Response Example



### 2. Check Location

Endpoint: /check-location
Method: POST
Description: Checks if the user has arrived at the specified destination based on their current location.







### Request Example
```json
{
    "latitude": -6.200000,
    "longitude": 106.816666
}
```
### Response Example
Success (200 OK)

```json
"User has arrived at the destination"
```
Failure (404 Not Found)

```json
"User is not at the destination"
```
## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/API-Transumindo.git

cd API-Transumindo
```

Install dependencies:
```bash
go mod tidy
```

Run the server:
```bash
go run main.go
```
The server will start running at http://localhost:8080.

## License

[MIT](https://choosealicense.com/licenses/mit/)


## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.

