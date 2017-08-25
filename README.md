
## Cap'n Proto VPC Flow Logs Compression

Any code, applications, scripts, templates, proofs of concept,
documentation and other items are provided for illustration purposes only.

Copyright 2017 Amazon Web Services

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


### Setup

[Install Cap'n Proto](https://capnproto.org/install.html)

Make sure you have $GOPATH/bin in your PATH

```
go get -u -t zombiezen.com/go/capnproto2/...
```

```
capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std -ogo vpc/flowlogs.capnp
```

```
go build
```

### Results

Raw: 81.644 MB
Rows: 775k

| Format                            | File Size |
| --------------------------------- | --------- |
| Compressed Text                   | 400 KB    |
| Cap'n Proto Packed and Compressed | 208 KB    |

