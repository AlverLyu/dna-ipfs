# dna-ipfs

A JSON-RPC service for accessing IPFS network.

## Usage

Make sure you have Go installed and run:

    go get -u -d github.com/AlverLyu/dna-ipfs

This will download the source into $GOPATH/src/github.com/AlverLyu/dna-ipfs.
Switch to the directory and build dna-ipfs:

    go build

Run `./dna-ipfs` to start the service, then you can access IPFS network
via the JSON-RPC API.

By default, this RPC service will listen on http://localhost:8080/rpc/ipfs and
access IPFS via http://localhost:5001, which means you should have an IPFS
service running on your machine and the HTTP API listening on port 5001. This
can be customized in the config file, which is `dnaipfs.cfg` by default and can
be specified by `-c` option.

You can also customize the logging system by specifying a .xml config file via
`-lc` option. There is a sample config file in `conf/sample/`.

## API

### addfile

Add a file to the IPFS network.

```
{
    "Method": "addfile",
    "Params": {
        "name": "THE_FILE_NAME",
        "data": "THE_FILE_DATA"
    }
}
```

The whole json object in `Params` will be stored into IPFS.

If succeeded, return the file ID in result. Otherwise the error code and a
simple description will be returned.
```
{
    "id": "",
    "jsonrpc": "2.0",
    "errrorcode": 0,
    "result": {
        "id": "THE_FILE_ID"
    }
}   
```

### getfile

Get a file from IPFS network.
```
{
    "Method": "getfile",
    "Params": {
        "id": "THE_FILE_ID"
    }
}
```

If succeeded, this api will return the json that passed to the `Params` of
`addfile` in result.
```
{
    "id": "",
    "jsonrpc": "2.0",
    "errrorcode": 0,
    "result": {
        "name": "THE_FILE_NAME",
        "data": "THE_FILE_DATA"
    }
}   
```

## LISENCE

Apache Licence Version 2.0
