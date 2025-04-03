# arago

## How to Run and Test the Service

### Prerequisites
- Docker

### Steps

1. **Run MongoDB Docker Container**
   ```sh
   docker run --name mongodb -d -p 27017:27017 mongo
   ```

2. **Run Dragonfly or Redis Docker Container**
   - To run Dragonfly:
     ```sh
     docker run -p 6379:6379 --ulimit memlock=-1 docker.dragonflydb.io/dragonflydb/dragonfly
     ```
   - To run Redis:
     ```sh
     docker run --name redis -d -p 6379:6379 redis
     ```

3. **Run Ad Server**
    Change values in `ad/config/config.go` to match your MongoDB and Dragonfly/Redis configurations.
   ```sh
   cd path/to/ad
   ```
   build and/or run

4. **Run Tracker Server**
   Change values in `ad/config/config.go` to match your MongoDB
   ```sh
   cd path/to/tracker
   ```
   build and/or run

5. **Test**
   `ad/client/main.go` is a simple client to test the service. You can run it to test the endpoints.
   `tracker/client/main.go` is a simple client to test the service. You can run it to test the endpoints.

### Testing the Service
- Ensure both MongoDB and Dragonfly/Redis containers are running.
- Ensure both the ad server and tracker server are running.
- Use appropriate tools or scripts to test the endpoints and functionality of the service.
