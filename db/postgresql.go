package db

import (
	postgrest "github.com/nedpals/postgrest-go/pkg"
	supa "github.com/nedpals/supabase-go"
)

// Init 数据库初始化连接
func Init() {
	supabaseUrl := "https://sfisgjpeqptcluzmtbup.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InNmaXNnanBlcXB0Y2x1em10YnVwIiwicm9sZSI6ImFub24iLCJpYXQiOjE2ODM5MzAyMzAsImV4cCI6MTk5OTUwNjIzMH0.eHhleg3ev4YGA1yHosohWwzxOZxNEh4hP1PavfMF-X0"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	userDB = supabase.DB
}

// completion DB
var userDB *postgrest.Client

// CompletionDB completion DB
func CompletionDB() *postgrest.Client {
	return userDB
}
