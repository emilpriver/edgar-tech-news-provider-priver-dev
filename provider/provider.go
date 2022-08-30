/*
This need to be changed to something unique
We recommend changeing to your domain and "provider".
ex:
example.com = ExampleComProvider
*/
package ExampleComProvider

import (
	"time"

	providers "github.com/edgar-systems/tech-news-website/Backend/providers"
)

type PostsSolution struct{}

/*
* Config returns information used for the application when it seeds information to the database.
accountType: can be one of "personal" or "company". This is used to split posts if it's a company that wrote it or if it come from a personal blog
contactEmail: used if needed for contact
*/

func (s *PostsSolution) Config() providers.Provider {
	return providers.Provider{
		AccountType:  "personal",
		ContactEmail: "example@example.com",
	}
}

/*
Here goes your logic.
Your Posts function should return all posts you want to add to the portal.
The important is that Posts should return a array of provider.Post which is all your posts.
How you build the logic is up to you.

We only read the /provider map
*/
func (s *PostsSolution) Posts() ([]providers.Post, error) {
	var posts []providers.Post

	posts = append(posts, providers.Post{
		Link:        "https://example.com",
		ID:          "4d41a4e5-0539-4823-89f6-6024aea8c7de",
		Excerpt:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent semper non leo vitae porttitor. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae",
		Title:       "Lorem ipsum",
		PublishedAt: time.Now(),
		Cover:       "https://picsum.photos/1000",
		Categories:  []string{"Golang"},
		Tags:        []string{"Tag"},
	})

	return posts, nil
}
