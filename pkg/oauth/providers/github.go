package providers

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/golang/glog"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/mayadata-io/kubera-auth/pkg/models"
	"github.com/mayadata-io/kubera-auth/pkg/types"
	controller "github.com/mayadata-io/kubera-auth/versionedController/v1"
)

func getUserFromToken(c *gin.Context, token *oauth2.Token) (*models.UserCredentials, error) {
	ctx := c.Request.Context()
	ts := oauth2.StaticTokenSource(token)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	githubUser, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	githubUserEmails, _, err := client.Users.ListEmails(ctx, &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	currTime := time.Now()
	user := models.UserCredentials{
		Name:         githubUser.Name,
		Kind:         models.GithubAuth,
		Role:         models.RoleUser,
		State:        models.StateActive,
		SocialAuthID: githubUser.ID,
		LoggedIn:     &types.TrueValue,
		CreatedAt:    &currTime,
	}

	for _, githubUserEmail := range githubUserEmails {
		if *githubUserEmail.Primary {
			user.Email = githubUserEmail.Email
			user.OnBoardingState = models.BoardingStateEmailVerified
			break
		}
	}

	return &user, err
}

//GetGithubUser gives the details of the user fetched as from github
func GetGithubUser(c *gin.Context) (*models.UserCredentials, error) {
	token, err := controller.Server.GithubConfig.GetToken(c)
	if err != nil {
		log.Errorln("Error getting token from github", err)
		return nil, err
	}

	githubUser, err := getUserFromToken(c, token)
	if err != nil {
		log.Errorln("Error getting user from github", err)
		return nil, err
	}

	return githubUser, nil
}
