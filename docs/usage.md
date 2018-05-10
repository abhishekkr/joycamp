
## Usage

* running a cmd (pointless)

```
go run main.go -cmd "ls -lah /tmp"
```


* running a cmd downloaded from provided `src`

```
go run main.go -src "http://127.0.0.1:5678/ls-lah"
```


* running a cmd from source with args provided

```
go run main.go -src "http://127.0.0.1:5678/ls" -args "-lah"
```


* running a cmd from source with args provided ad custom additional environment

```
go run main.go -src "http://127.0.0.1:5678/print-env" -args "-env EXAMPLE" -env '{"EXAMPLE": "think"}'
```

* to run all above variations using [config json](./job01.json)

```
go run main.go -cfg ./job01.json
```

---
