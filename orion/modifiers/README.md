# modifiers
`import "github.com/carousell/Orion/orion/modifiers"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-imports">Imported Packages</a>

- [github.com/carousell/Orion/utils/options](./../../utils/options)

## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func GetSerilizationType(ctx context.Context) (string, bool)](#GetSerilizationType)
* [func SerializeOutJSON(ctx context.Context)](#SerializeOutJSON)
* [func SerializeOutJSONPB(ctx context.Context)](#SerializeOutJSONPB)
* [func SerializeOutProtBuf(ctx context.Context)](#SerializeOutProtBuf)

#### <a name="pkg-files">Package files</a>
[modifiers.go](./modifiers.go) [types.go](./types.go) 

## <a name="pkg-constants">Constants</a>
``` go
const (
    Request_HTTP = "OrionRequestHTTP"
    Request_gRPC = "OrionRequestGRPC"

    JSON     = "JSON"
    JSONPB   = "JSONPB"
    ProtoBuf = "PROTO"
)
```

## <a name="GetSerilizationType">func</a> [GetSerilizationType](./modifiers.go#L30)
``` go
func GetSerilizationType(ctx context.Context) (string, bool)
```

## <a name="SerializeOutJSON">func</a> [SerializeOutJSON](./modifiers.go#L18)
``` go
func SerializeOutJSON(ctx context.Context)
```

## <a name="SerializeOutJSONPB">func</a> [SerializeOutJSONPB](./modifiers.go#L22)
``` go
func SerializeOutJSONPB(ctx context.Context)
```

## <a name="SerializeOutProtBuf">func</a> [SerializeOutProtBuf](./modifiers.go#L26)
``` go
func SerializeOutProtBuf(ctx context.Context)
```

- - -
Generated by [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)