# Workout of the day micro api

## Functionality

This micro api serves and stores workouts for the day

## Routes

### Request api endpoint

```bash 
GET  /work-out-of-the-day

``` 

### Success Response

    Response body:
    {
        "title":"Fran",
        "workout":"21–15–9 Reps Thruster 95lbs Pull Ups"
    }

```bash 
POST /work-out-of-the-day
```

### Success Response

    Response body:
    {
        "Added wod with title: Lynne"
    }

### Error Response

    Response body:
    {
        "message": "Server Error"
    }

### Build

``` bash
 docker-compose build
 ```

### Run

``` bash
 docker-compose up
 visit http://localhost:8081
 ```
