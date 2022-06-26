
# Micobian Party

This service is responsible for managing events and employees in Micobo .


## Run Locally

Clone the project

```bash
  git clone https://github.com/tinaaliakbarpour/micoboParty
```

Go to the project directory

```bash
  cd micoboParty
```

Start the server

```bash
  docker-compose up -d 
  docker exec -it container_name(micobo) bash
  go run main.go migrate
```
## Run tests

```bash
  for testing the repository part you have to go inside the container and then run 
  go test -v . -cover

```


## API Reference
### Employee API
#### Post new employee

```http
  Post /api/v1/employee/
```

| Parameter   | Type     | Description                       |
| :--------   | :------- | :-------------------------------- |
| `firstname` | `string` | **Required**.                     |
| `lastname`  | `string` | **Required**.                     |
| `gender`    | `string` | **Required**.                     |
| `event_id`  | `int64`  |                                   |
| `birth_day` | `string` | **Required**.                     |

 ** the birthday format should be like this -> 2021-02-02T00:00:00Z
#### Get all employees

```http
  GET /api/v1/employees/
```

#### Update employee
```http
  PUT /api/v1/employees/{employee_id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**.                     |

#### Delete employee
```http
  DELETE /api/v1/employees/{employee_id}
```

### Event API

#### Get all events

```http
  GET /api/v1/events/
```

#### Get specific event by event_id

```http
  GET /api/v1/events/{event_id}
```

#### Get list of the employees that are assisting to the event 
##### this endpoint accepts query parameters

```http
  GET /api/v1/events/{event_id}/employees
```
```curl
curl -X GET http://localhost:8080/api/v1/events/1/employees?filter="first_name:John"
```
## TODO

- test cases are not impemented so the next step for me is to complete the unit tests

- it should be some kind of mechnism to prevent admin from registering repetitive records
(maybe adding some extra unique fields like user_id or identification id to make them unique)
- and also should prepare some kind of mechanism for filter parameters as if someone wanted to make a damage they couldn't do it via filter parameters(if we make these endpoints private so no body can access them outside our internal network also makes it better)
- the test cases scenarios are really lazy i know but actually it takes time and  ... :((

- we can also have a make file and config.test.yaml and config.example.yaml for testing purposes

- i couldn't test the repository with sql mock as it doesn't support the gorm V2 package so i used the real connection to db to test the functionality

- i choose gorm package because i wanted to test it with sql mock and it doesn't work in the end if i have to choose the package again i will use go-pg as it is really simple to use and also really efficient for postgres





## Authors

- [@tinaaliakbarpour](https://www.github.com/tinaaliakbarpour)

