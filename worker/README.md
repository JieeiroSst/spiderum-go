go get github.com/gocraft/work/cmd/workwebui

go install github.com/gocraft/work/cmd/workwebui

workwebui -redis="redis:6379" -ns="work" -listen=":5040"