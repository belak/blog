package blog

import (
	"net/http"

	"github.com/gorilla/feeds"
)

func (s *Site) handleAtom(w http.ResponseWriter, r *http.Request) {
	feed := &feeds.Feed{
		Title: s.Title,
		Link:  &feeds.Link{Href: "https://belak.io/"},
	}

	if len(s.Posts) > 0 {
		feed.Updated = s.Posts[0].Date
	}

	for _, post := range s.Posts {
		link := "https://belak.io" + post.Permalink()
		item := &feeds.Item{
			Title:   post.Title,
			Link:    &feeds.Link{Href: link},
			Id:      link,
			Created: post.Date,
			Updated: post.Date,
			Content: post.Content,
		}
		if post.Summary != "" {
			item.Description = post.Summary
		}
		feed.Items = append(feed.Items, item)
	}

	atom, err := feed.ToAtom()
	if err != nil {
		http.Error(w, "feed generation error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
	w.Write([]byte(atom))
}
