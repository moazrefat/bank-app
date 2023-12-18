module github.com/moazrefat/bankapp

go 1.21.5

require (
	github.com/go-sql-driver/mysql v1.7.1
	golang.org/x/crypto v0.17.0
)

replace github.com/hardw01f/Vulnerability-goapp/pkg/cookei => ./pkg/cookie

replace github.com/hardw01f/Vulnerability-goapp/pkg/login => ./pkg/login

replace github.com/hardw01f/Vulnerability-goapp/pkg/logout => ./pkg/logout

replace github.com/hardw01f/Vulnerability-goapp/pkg/register => ./pkg/register
