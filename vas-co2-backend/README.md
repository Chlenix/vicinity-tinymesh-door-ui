## Components

#### Config

Contains a struct datatype containing the environment variables. The reference is included in all the other major components.

#### Controller

`controller.go` sets up the router for all incoming API requests inside the `setupRouter()` function.

`handler.go` contains callbacks that handle each incoming API request.

#### Database

`db.go` establishes a DB connection regulated by the `dialect` constant (default: `postgres`)

#### Model
`sensor.go` struct for a door sensor

`reading.go` struct for a door event 

relation: sensor 0..* reading

#### Vicinity
`vicinity.go` creates the TD and services all requests involving the sensors/readings. The various methods are called from within the `handler.go` file of the Controller.