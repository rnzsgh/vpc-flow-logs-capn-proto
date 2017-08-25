
[Install Cap'n Proto](https://capnproto.org/install.html)

Make sure you have $GOPATH/bin in your PATH

```
capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std -ogo vpc/flowlogs.capnp
```

```
go build
```

Text
Entries: 775,000
Raw: 81.644 MB
Compressed (zlib): 400 KB

Cap'n Proto Packed
Entries: 775,000
Raw: NA
Compressed (zlib): 208 KB


