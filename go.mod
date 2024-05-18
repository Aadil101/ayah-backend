module github.com/Aadil101/ayah-backend

go 1.22

replace github.com/Aadil101/ayah-backend/pkg/handler => /home/code_is_life/ayah-backend/pkg/handler

replace github.com/Aadil101/ayah-backend/pkg/internal => /home/code_is_life/ayah-backend/pkg/internal

require github.com/stretchr/testify v1.8.4

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
