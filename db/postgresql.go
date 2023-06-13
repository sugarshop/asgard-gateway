package db

import (
	postgrest "github.com/nedpals/postgrest-go/pkg"
	supa "github.com/nedpals/supabase-go"
	"github.com/sugarshop/env"
)

// PostgreSqlInit Postgresql 数据库初始化连接
func PostgreSqlInit() {
	supabaseUrl, ok := env.GlobalEnv().Get("SUPABASEURL")
	if !ok {
		panic("no SUPABASEURL env set")
	}
	supabaseKey, ok := env.GlobalEnv().Get("SUPABASEKEY")
	if !ok {
		panic("no SUPABASEKEY env set")
	}
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	userDB = supabase.DB
}

// completion DB
var userDB *postgrest.Client

// CompletionDB completion DB
func CompletionDB() *postgrest.Client {
	return userDB
}
