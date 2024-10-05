package jwt

import (
	"log/slog"
	"net/http"
	"time"

	e "hackathon/exceptions"
	"hackathon/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/valyala/fasthttp"
)

type Response struct {
	Error   string `json:"error"`
	Content any    `json:"content"`
}

type JWT struct {
	secret  string
	expTime time.Duration
}

func New(cfg *config.Config) *JWT {
	return &JWT{
		secret:  cfg.JWTsecret,
		expTime: time.Duration(cfg.ExpTime) * time.Second,
	}
}

func (j *JWT) CreateToken(guid string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"guid": guid,
		"exp":  time.Now().Add(j.expTime).Unix(),
	})

	stringToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		slog.Error(err.Error())
	}

	return stringToken
}

func (j *JWT) ValidateToken(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := string(c.Request().Header.Peek(fasthttp.HeaderAuthorization))

		if len(authHeader) == 0 {
			return c.Status(http.StatusUnauthorized).JSON(Response{
				Error:   e.ErrInvalidToken.Error(),
				Content: nil,
			})
		}
		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, e.ErrInvalidSigningMethod
			}
			return []byte(j.secret), nil
		})

		if err != nil {
			slog.Error(err.Error())
			return c.Status(http.StatusUnauthorized).JSON(Response{
				Error:   e.ErrInvalidToken.Error(),
				Content: nil,
			})
		} else if !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(Response{
				Error:   e.ErrInvalidToken.Error(),
				Content: nil,
			})
		}

		return next(c)
	}
}
