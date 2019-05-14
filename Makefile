APP_NAME := bb_matching
REPO := github.com/natsu-summer72/BbMatching
SWAGGER_DIR := ./server/swagger-ui/swagger/

# goa関連
goagen:
	@goa gen $(REPO)/design
	@cp -f ./gen/http/openapi.json ${SWAGGER_DIR}
	@cp -f ./gen/http/openapi.yaml ${SWAGGER_DIR}

regen:
	@rm -rf ./cmd
	@goa example $(REPO)/design
	@mv -n *.go controller/

run:
	@cd cmd/$(APP_NAME) && go build
	@./cmd/$(APP_NAME)/$(APP_NAME)

clean:
	@rm -rf cmd/
	@rm -rf gen/


