# jsonsort

Sort for ndjson.

## Usage

Sort each lines composed of JSON by the value pointed by jsonpath.

```
cat output.json | jsonsort -p $.post.record.createdAt
```

## Installation

```
go install github.com/mattn/jsonsort
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
