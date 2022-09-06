package PriverDevProvider

import "testing"

func TestPosts(t *testing.T) {
	posts, err := Posts()
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) == 0 {
		t.Fatal("Expected posts to have posts, got 0")
	}
}
