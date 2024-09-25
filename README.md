JSON encoder / Decoder exercice
===============================

This is a simple exercise, the goal being to be able to encrypt some values on a
JSON body (only on a first level), or decrypt values in it, sign a JSON body or 
check if a provided signature matches the signature used in the program.

How to build
------------
### From Source

```bash
git clone https://github.com/Taluu/json-encoder-exercise
cd json-encrypter-exercise
go build -o bin/json-encrypter-exercise
bin/json-encrypter-exercise
```

### Fron binary release
No releases yet :P

Endpoints
---------
### Encrypting data
You can encrypt data by hitting the endpoint (with default values for host and
ports, `localhost` and `8080`) on a POST request, with a valid json object :

```bash
curl http://localhost:8080/encrypt -X POST -d "{\"foo\": \"bar\", \"number\": 1, \"object\": {\"one\": \"two\", \"three\": 3}}"
```

You should get a 200 response with the following values :

```json
{
    "foo": "YmFy",
    "number":"MQ==",
    "object":"eyJvbmUiOiJ0d28iLCJ0aHJlZSI6M30="
}
```

On errors, you can have a 405 on a method other than POST, or 400 if the json body cannot be decoded into an object.
