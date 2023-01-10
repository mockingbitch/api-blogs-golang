# golang-jwt-template
api {
  http://localhost:9090/api/register => "message": created
  http://localhost:9090/api/login => "token": jwtToken
  http://localhost:9090/api/secured/ping => "message": "Pong"
}
