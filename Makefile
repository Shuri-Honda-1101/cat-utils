.PHONY: ent_generate ent_init

ent_generate:
	go generate ./ent && go run github.com/hedwigz/entviz/cmd/entviz ./ent/schema

ent_init:
	go run -mod=mod entgo.io/ent/cmd/ent new ${name}