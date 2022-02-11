build: bin/astviz

bin/astviz: main.go astviz/astviz.go astviz/model.go
	mkdir -p bin || true
	go build -o bin/astviz .
test:
	watchexec -r -- go test -v gitlab.com/fawad/ast-viz/astviz
