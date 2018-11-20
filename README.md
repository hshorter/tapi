# README

Skeleton OAS server implementation for TAPI Latest, note this was built on 20th Nov 2018 .

## Generate

Read this [page](https://github.com/MEF-GIT/MEF-LSO-Presto-SDK/tree/master/published/swagger) to understand how to generate a single OAS specification from a set of YANG files.

```bash
java -cp ~/Downloads/yang2swagger-tapi-cli-1.0-cli.jar com.amartus.y2s.Generator -yang-dir TAPI/YANG -output tapi-2.1.0.yml
```

Followed by 

```bash
swagger generate server -f tapi-2.1.0.yml -A tapi
```

## Install

```bash
go install github.com/damianoneill/tapi/...
```

## Run

```bash
$ tapi-server --port=8080
2018/11/14 18:44:36 Serving tapi at http://127.0.0.1:8080
```

## Test

```bash
$ curl http://127.0.0.1:8080/restconf/data/context
<string>operation tapi_common.GetDataContext has not yet been implemented</string>
```

## Docker

Build image 

```bash
docker build -t tapi-server .
```

Run image

```bash
docker run -p 8080:8080 --name tapi tapi-server
```