package user

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"loyalid-test/domain"
)

func (s *Service) CreateUser(ctx context.Context, user *domain.User) (err error) {
	user.GenerateId()

	if err = user.HashPassword(); err != nil {
		return
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *Service) Authenticate(ctx context.Context, user *domain.User) (token string, err error) {
	reqPassword := user.Password()

	if *user, err = s.repo.GetUserByUsername(ctx, user.Username); err != nil {
		return
	}

	if err = user.ValidatePassword(reqPassword); err != nil {
		return
	}

	url := "https://dev-waygvmnclik5ibp8.us.auth0.com/oauth/token"

	payload := strings.NewReader("grant_type=client_credentials&client_id=UshkHutwQ8D3l1aXN1MRvBB9LO4ICLRn&client_secret=1AHiu7EB1ngC1eQ_zdwY_2wUSNcYS_2X8NW3YMpRIkbb_pvdRLluldHHV3NzwlVq&audience=https://loyalid-test/get-token")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	token = string(body)
	return
}

func (s *Service) CurrentUser(ctx context.Context) (user domain.User, err error) {
	return s.repo.GetUserByUsername(ctx, domain.UsernameCtx(ctx))
}
