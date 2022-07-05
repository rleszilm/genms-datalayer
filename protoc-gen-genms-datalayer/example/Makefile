PB_OPTIONS := --experimental_allow_proto3_optional
PB_GO_INCLUDE := -I ../../pkg -I .
PB_GO_COMPILE := --go_opt=paths=source_relative \
	--go_out=. \
	--genms-datalayer_opt=paths=source_relative \
	--genms-datalayer_out=.
PB_GO_SRC := .

all:
	protoc $(PB_OPTIONS) $(PB_GO_INCLUDE) $(PB_GO_COMPILE) `ls $(PB_GO_SRC)/*.proto`
