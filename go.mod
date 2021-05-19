module go-fiber

go 1.15

require (
	book/book v0.0.0
	database/database v0.0.0
	github.com/andybalholm/brotli v1.0.2 // indirect
	github.com/gofiber/fiber/v2 v2.9.0
	github.com/klauspost/compress v1.12.2 // indirect
	github.com/valyala/fasthttp v1.24.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
	gorm.io/driver/postgres v1.1.0 // indirect
	gorm.io/gorm v1.21.10 // indirect
)

replace (
	book/book v0.0.0 => ./book
	database/database v0.0.0 => ./database
)
