dev password: 111111

openssl genrsa -out resources/cert/echo.key 2048
openssl req -new -x509 -sha256 -key resources/cert/echo.key -out resources/cert/echo.crt -days 3650
openssl req -new -sha256 -key resources/cert/echo.key -out resources/cert/echo.csr
openssl x509 -req -sha256 -in resources/cert/echo.csr -signkey resources/cert/echo.key -out resources/cert/echo.crt -days 3650
