package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/tomhockett/learn-go-with-tests/blogposts"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := blogposts.Post{Title: "Post 1"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestNewBlogPosts_fails(t *testing.T) {
	_, err := blogposts.NewPostsFromFS(StubFailingFS{})
	if err == nil {
		t.Fatal("expected error but got none")
	}
}
