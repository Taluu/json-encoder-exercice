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
Note : all example are hitting as if the domain is `localhost` and the port
`8080`. These can be configured when executing the binary with the flags 
`-domain` and `port` : `bin/json-encoder-exercise -domain localhost -port 8080`

### Encrypting data
You can encrypt data by hitting the endpoint `/encrypt` on a POST request, with
a valid json object :

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

On errors, you can have a 405 on a method other than POST, or 400 if the json
body cannot be decoded into an object.

### Decrypting data
You can decrypt data, and the server will try to decode also nested values (such
as in an object or an array) by hitting the endpoint `/decrypt` on a POST request,
with a valid json object :

```bash
curl http://localhost:8080/decrypt -X POST -d "{\"foo\": \"YmFy\", \"number\": 1, \"object\": \"eyJvbmUiOiJ0d28iLCJ0aHJlZSI6M30=\"}"
```

You should then get a 200 response with the following values :

```json
{
    "foo": "bar",
    "number": 1,
    "object": {
        "one": "two",
        "three": 3
    }
}
```

If an encoded value was a json object or array, it will be returned as such, but
if there was an encrypted value in them, they won't be decrypted.

On errors, you can have a 405 on a method other than POST, or 400 on bad json
input.
