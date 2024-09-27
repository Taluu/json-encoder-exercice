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
bin/json-encrypter-exercise -key "my secret key"
```

The `-key` flag is important here, to set the key used to generate and verify
signatures for the `/sign` and `/verify` endpoints. 

Other available flags are `-domain` and `-port`, with respective default values
being `localhost` and `8080`.

### Fron binary release
No releases yet :P

Endpoints
---------
Note : all example are hitting as if the domain is `localhost` and the port
`8080`, and the key `test`.

For security concerns, the key should of course not be plaintext when launching
the binary, but for the sake of simplicity and this being out of scope for this
exercice, let's pass this as plaintext.

`bin/json-encoder-exercise -domain localhost -port 8080 -key test`

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

### Signing data
You can sign data (any data) as long as it's in a json format, by hitting the
endpoint `/sign` like the following :

```bash
curl http://localhost:8080/sign -X POST -d "{\"foo\": \"YmFy\"}"
```

You should get a 200 response with the following json body :

```json
{
    "signature": "384854f7b73ed17f1107006aa49e3813a7d1f9f3ce75fa0b770bc00e35c8ea82",
    "data": {
        "foo":"YmFy"
    }
}
```

The `data` property of the returned objet will be of the type of the given data
(if it's a scalar, it will be scalar, an array an array, ... and so on)

On errors, you can have a 405 for a method other than POST, or 400 on bad json
input (or empty input).
