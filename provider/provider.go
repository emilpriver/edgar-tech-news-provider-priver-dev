/*
This need to be changed to something unique
We recommend changeing to your domain and "provider".
ex:
example.com = ExampleComProvider
*/
package PriverDevProvider

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	providers "github.com/edgar-systems/tech-news-provider"
)

type PostsSolution struct{}

var postsUrl = "https://priver.dev/blog/index.json"

/*
* Config returns information used for the application when it seeds information to the database.
accountType: can be one of "personal" or "company". This is used to split posts if it's a company that wrote it or if it come from a personal blog
contactEmail: used if needed for contact
*/

func (s *PostsSolution) Config() providers.Provider {
	return providers.Provider{
		AccountType:         "personal",
		ContactEmail:        "emil@priver.dev",
		AuthorUrl:           "https://priver.dev",
		AuthorDefaultTopics: []string{},
		AuthorDefaultCover:  "https://cdn.priver.dev/?src=https://priver.dev/images/me.png&w=1000",
	}
}

type BlogJson struct {
	Data []struct {
		Title     string    `json:"title"`
		Date      time.Time `json:"date"`
		Type      string    `json:"type"`
		Permalink string    `json:"permalink"`
		Summary   string    `json:"summary"`
		Hero      string    `json:"hero"`
		Tags      string    `json:"tags"`
	} `json:"data"`
}

func (s *PostsSolution) Posts() ([]providers.Post, error) {
	var posts []providers.Post
	var data BlogJson

	resp, err := http.Get(postsUrl)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range data.Data {
		tags := []string{}

		for _, tag := range strings.Split(item.Tags, ",") {
			newTag := strings.ReplaceAll(strings.Trim(tag, " "), ",", "")

			if len(newTag) > 0 {
				tags = append(tags, newTag)
			}
		}

		posts = append(posts, providers.Post{
			Link:        item.Permalink,
			ID:          item.Permalink,
			Excerpt:     item.Summary,
			Title:       item.Title,
			PublishedAt: item.Date,
			Cover:       item.Hero,
			Categories:  []string{},
			Tags:        tags,
		})
	}

	return posts, nil
}
