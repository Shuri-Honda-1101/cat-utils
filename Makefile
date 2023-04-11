.PHONY: ent_generate ent_init

ent_generate:
	go generate ./ent

ent_init:
	go run -mod=mod entgo.io/ent/cmd/ent new ${name}