package blog

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/a-h/templ"
)

type MenuItem struct {
	Name string
	URL  string
}

type TagInfo struct {
	Name  string
	Count int
}

type Site struct {
	Title     string
	Posts     []*Post
	Pages     []*Page
	Tags      map[string][]*Post
	MenuItems []MenuItem

	IndexContent string

	assetsFS fs.FS
}

func LoadSite(blogFS, pagesFS, assetsFS fs.FS) (*Site, error) {
	posts, err := LoadPosts(blogFS)
	if err != nil {
		return nil, fmt.Errorf("loading posts: %w", err)
	}

	var published []*Post
	for _, p := range posts {
		if !p.Draft {
			published = append(published, p)
		}
	}

	pages, err := LoadPages(pagesFS)
	if err != nil {
		return nil, fmt.Errorf("loading pages: %w", err)
	}

	indexFM, indexContent, err := LoadIndexContent(pagesFS)
	if err != nil {
		return nil, fmt.Errorf("loading index: %w", err)
	}

	tags := make(map[string][]*Post)
	for _, p := range published {
		for _, t := range p.Tags {
			tags[t] = append(tags[t], p)
		}
	}

	var menuItems []MenuItem
	menuItems = append(menuItems, MenuItem{Name: "Blog", URL: "/posts/"})
	for _, p := range pages {
		if p.Menu != "" {
			menuItems = append(menuItems, MenuItem{
				Name: p.Title,
				URL:  "/" + p.Slug + "/",
			})
		}
	}

	return &Site{
		Title:        indexFM.Title,
		Posts:        published,
		Pages:        pages,
		Tags:         tags,
		MenuItems:    menuItems,
		IndexContent: indexContent,
		assetsFS:     assetsFS,
	}, nil
}

type basePage struct {
	SiteTitle string
	PageTitle string
	IsHome    bool
	Section   string
	MenuItems []MenuItem
	Year      int
}

func (s *Site) newBasePage(pageTitle, section string) basePage {
	return basePage{
		SiteTitle: s.Title,
		PageTitle: pageTitle,
		Section:   section,
		MenuItems: s.MenuItems,
		Year:      time.Now().Year(),
	}
}

func (s *Site) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /", s.handleIndex)
	mux.HandleFunc("GET /posts/", s.handlePostList)
	mux.HandleFunc("GET /tags/", s.handleTags)
	mux.HandleFunc("GET /tags/{tag}/", s.handleTagPosts)
	mux.HandleFunc("GET /feed.xml", s.handleAtom)
	mux.Handle("GET /assets/", http.FileServerFS(s.assetsFS))

	for _, post := range s.Posts {
		p := post
		mux.HandleFunc("GET "+p.Permalink(), s.handlePost(p))
	}

	for _, page := range s.Pages {
		pg := page
		mux.HandleFunc("GET /"+pg.Slug+"/", s.handlePage(pg))
	}
}

func (s *Site) renderTempl(w http.ResponseWriter, r *http.Request, bp basePage, main func() templ.Component) {
	component := base(bp, main())
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("render error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (s *Site) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.handle404(w, r)
		return
	}

	bp := s.newBasePage(s.Title, "")
	bp.IsHome = true
	s.renderTempl(w, r, bp, func() templ.Component {
		return indexPage(s.IndexContent)
	})
}

func (s *Site) handlePostList(w http.ResponseWriter, r *http.Request) {
	const perPage = 10
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if page < 1 {
		page = 1
	}

	start := (page - 1) * perPage
	if start > len(s.Posts) {
		start = len(s.Posts)
	}
	end := start + perPage
	if end > len(s.Posts) {
		end = len(s.Posts)
	}

	totalPages := (len(s.Posts) + perPage - 1) / perPage

	posts := s.Posts[start:end]
	bp := s.newBasePage("Blog", "/posts/")
	s.renderTempl(w, r, bp, func() templ.Component {
		return postList(posts, page, totalPages)
	})
}

func (s *Site) handlePost(post *Post) http.HandlerFunc {
	var prev, next *Post
	for i, p := range s.Posts {
		if p == post {
			if i > 0 {
				next = s.Posts[i-1]
			}
			if i < len(s.Posts)-1 {
				prev = s.Posts[i+1]
			}
			break
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		bp := s.newBasePage(post.Title, "/posts/")
		s.renderTempl(w, r, bp, func() templ.Component {
			return postSingle(post, prev, next)
		})
	}
}

func (s *Site) handlePage(page *Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bp := s.newBasePage(page.Title, "/"+page.Slug+"/")
		s.renderTempl(w, r, bp, func() templ.Component {
			return pageView(page)
		})
	}
}

func (s *Site) handleTags(w http.ResponseWriter, r *http.Request) {
	var tagList []TagInfo
	for name, posts := range s.Tags {
		tagList = append(tagList, TagInfo{Name: name, Count: len(posts)})
	}
	sort.Slice(tagList, func(i, j int) bool {
		return tagList[i].Name < tagList[j].Name
	})

	bp := s.newBasePage("Tags", "")
	s.renderTempl(w, r, bp, func() templ.Component {
		return tagsPage(tagList)
	})
}

func (s *Site) handleTagPosts(w http.ResponseWriter, r *http.Request) {
	tag := r.PathValue("tag")
	posts, ok := s.Tags[tag]
	if !ok {
		s.handle404(w, r)
		return
	}

	bp := s.newBasePage("#"+tag, "")
	s.renderTempl(w, r, bp, func() templ.Component {
		return tagPostsPage(tag, posts)
	})
}

func (s *Site) handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	bp := s.newBasePage("404", "")
	s.renderTempl(w, r, bp, func() templ.Component {
		return notFoundPage()
	})
}
