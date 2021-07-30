# ID64: Quick non-distributed ordered 64-bit ID generator

Quick ordered 64-bit ID generator, single thread only, can generate hundreds of millions (~276M) IDs per second.

Consist of 2 segments:
- time/second segment (32 bit, offsetted to 2021-01-01)
- counter segment (32 bit, reset to 0 every second)

**warning** this is not for multi-node, since there's no identity segment.

## Benchmark

```
cpu: AMD Ryzen 3 3100 4-Core Processor    
BenchmarkShortuuid-8      118908      8572 ns/op
BenchmarkKsuid-8          760924      1493 ns/op
BenchmarkNanoid-8         759548      1485 ns/op
BenchmarkUuid-8           935152      1304 ns/op
BenchmarkTime-8          1690483       720.0 ns/op
BenchmarkSnowflake-8     4911249       244.7 ns/op
BenchmarkLexIdNano-8     8483720       138.8 ns/op
BenchmarkLexIdNoSep-8   10396551       116.3 ns/op
BenchmarkLexIdNoLex-8   10590300       115.1 ns/op
BenchmarkLexId-8         9991906       114.9 ns/op
BenchmarkXid-8          13754178        86.02 ns/op
BenchmarkId64-8        276799974         4.362 ns/op
```

## Usage

```
import "github.com/kokizzu/id64"

func main() {
   id := id64.ID()
   
   // or object-oriented version (eg. when you need different generator per table)
   gen := id64.Generator{}
   id = gen.ID()
   
   u64 := uint64(id) // get uint64 value
   
   t := id.Time() // get time.Time of that id
   
   i32 := id.Counter() // get counter of that id
   
   s := id.String() // string representation (base64-like encoding) 
}

```

## Gotchas

- can be duplicate if your processor can increment faster than 4.2 billion times per second
- might overflow at 2156-07-28 02:08:00 UTC

## See also

[lexid](//github.com/kokizzu/lexid) - if you need distributed ID (have identity segment)
