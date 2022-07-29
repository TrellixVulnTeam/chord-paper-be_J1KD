package user_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/veedubyou/chord-paper-be/go-rewrite/src/lib/testing"
	authtest "github.com/veedubyou/chord-paper-be/go-rewrite/src/shared_tests/auth"
	usergateway "github.com/veedubyou/chord-paper-be/go-rewrite/src/user/gateway"
	userstorage "github.com/veedubyou/chord-paper-be/go-rewrite/src/user/storage"
	userusecase "github.com/veedubyou/chord-paper-be/go-rewrite/src/user/usecase"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("User", func() {
	var (
		userGateway usergateway.Gateway
		validator   TestingValidator
	)

	BeforeEach(func() {
		validator = TestingValidator{}
		userStorage := userstorage.NewDB(db)
		userUsecase := userusecase.NewUsecase(userStorage, validator)
		userGateway = usergateway.NewGateway(userUsecase)
	})

	BeforeEach(func() {
		ResetDB(db)
	})

	Describe("Login", func() {
		Describe("Shared auth tests", func() {
			BeforeEach(func() {
				authtest.Endpoint = userGateway.Login
			})

			authtest.ItRejectsUnauthorizedRequests("POST", "/login")
		})

		Describe("For an authorized user", func() {
			var (
				response *httptest.ResponseRecorder
			)

			BeforeEach(func() {
				request := RequestFactory{
					Method:  "POST",
					Path:    "/login",
					JSONObj: nil,
					Mods:    RequestModifiers{WithUserCred(PrimaryUser)},
				}.Make()
				response = httptest.NewRecorder()

				c := PrepareEchoContext(request, response)
				err := userGateway.Login(c)
				Expect(err).NotTo(HaveOccurred())
			})

			It("succeeds", func() {
				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("returns the correct user", func() {
				userResponse := DecodeJSON[usergateway.UserJSON](response)
				Expect(userResponse).To(BeEquivalentTo(PrimaryUser))
			})
		})
	})
})
