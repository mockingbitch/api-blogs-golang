# golang-jwt-template
api { \n
    http://localhost:9090/api/register => "message": created \n
    http://localhost:9090/api/login => "token": jwtToken \n
    http://localhost:9090/api/secured/ping => "message": "Pong" \n
}
