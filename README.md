# Resume generator
Generate resume from templates by editing your information in form of JSON files.

## Usage
### Build
You need to have Go installed to build this program.
```terminal
go build main.go
```
### Run
```terminal
Usage: main template_folder input output
```
For example:
```terminal
./main templates/simple configs/sample.json dist/index.html
```