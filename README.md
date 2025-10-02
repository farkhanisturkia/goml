#### **Instalation**
a. create new project
```
mkdir example-app
cd example-app
go mod init example-app
```
b. install module
```
go get github.com/farkhanisturkia/goml
```
#### **Datas Preparation Tools**
models:
`FMamdani`, `FTsukamoto`

Parsing JSON:
```
datas, err := goml.ParseJSON(jsonStr)
if err != nil {
    fmt.Println("Error:", err)
    return
}
```
Validating JSON:
```
models := goml.Models{Model: goml.FMamdani}
err = goml.ValidationJSON(datas, models)
if err != nil {
    fmt.Println("Error:", err)
    return
}
```
#### **Fuzzy Mamdani**
Example JSON as Body Input:
```
jsonStr := []byte(`{
    "Input": {"A": 30, "B": 70},
    "Dataset": [
        {"A": [0, 1, 5]},
        {"B": [25, 75, 43, 12, 87]}
    ]
}`)
```
Run Fuzzy Mamdani:
```
mamdani, err := goml.UseMamdani(datas)
if err != nil {
    fmt.Println("Error:", err)
    return
}
```
