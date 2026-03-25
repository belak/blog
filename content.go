package blog

import (
	"bytes"
	"fmt"
	"io/fs"
	"sort"
	"strings"
	"time"

	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

// Post represents a single blog post.
type Post struct {
	Title   string
	Date    time.Time
	Updated time.Time
	Tags    []string
	Draft   bool
	Slug    string

	// Summary is the content before <!--more-->, rendered to HTML.
	Summary string
	// Content is the full post content rendered to HTML.
	Content string
}

// Permalink returns the URL path for this post.
func (p *Post) Permalink() string {
	return fmt.Sprintf("/%d/%02d/%s/", p.Date.Year(), p.Date.Month(), p.Slug)
}

// Page represents a standalone page (like about).
type Page struct {
	Title       string
	Description string
	Updated     time.Time
	TOC         bool
	Menu        string
	Content     string
	Slug        string
}

type frontmatter struct {
	Title       string   `yaml:"title"`
	Date        string   `yaml:"date"`
	Updated     string   `yaml:"updated"`
	Tags        []string `yaml:"tags"`
	Draft       bool     `yaml:"draft"`
	Description string   `yaml:"description"`
	TOC         bool     `yaml:"toc"`
	Menu        string   `yaml:"menu"`
}

var md = goldmark.New(
	goldmark.WithExtensions(
		highlighting.NewHighlighting(
			highlighting.WithFormatOptions(
				html.WithClasses(true),
			),
		),
	),
	goldmark.WithRendererOptions(
		goldmarkhtml.WithUnsafe(),
	),
)

func parseFrontmatter(data []byte) (frontmatter, string, error) {
	var fm frontmatter

	content := string(data)
	if !strings.HasPrefix(content, "---\n") {
		return fm, content, nil
	}

	end := strings.Index(content[4:], "\n---\n")
	if end == -1 {
		return fm, content, nil
	}

	if err := yaml.Unmarshal([]byte(content[4:4+end]), &fm); err != nil {
		return fm, "", fmt.Errorf("parsing frontmatter: %w", err)
	}

	body := content[4+end+5:]
	return fm, body, nil
}

func renderMarkdown(source string) (string, error) {
	var buf bytes.Buffer
	if err := md.Convert([]byte(source), &buf); err != nil {
		return "", fmt.Errorf("rendering markdown: %w", err)
	}
	return buf.String(), nil
}

// processNoticeShortcodes replaces {{< notice >}}...{{< /notice >}} with
// <p class="notice">...</p>.
func processNoticeShortcodes(content string) string {
	content = strings.ReplaceAll(content, "{{< notice >}}", `<p class="notice">`)
	content = strings.ReplaceAll(content, "{{< /notice >}}", "</p>")
	return content
}

// LoadPosts reads all posts from the blog/ embedded filesystem.
func LoadPosts(blogFS fs.FS) ([]*Post, error) {
	entries, err := fs.ReadDir(blogFS, "blog")
	if err != nil {
		return nil, fmt.Errorf("reading blog directory: %w", err)
	}

	var posts []*Post
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		data, err := fs.ReadFile(blogFS, "blog/"+entry.Name())
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", entry.Name(), err)
		}

		fm, body, err := parseFrontmatter(data)
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %w", entry.Name(), err)
		}

		var postDate time.Time
		if fm.Date != "" {
			postDate, err = time.Parse(time.RFC3339, fm.Date)
			if err != nil {
				return nil, fmt.Errorf("parsing date in %s: %w", entry.Name(), err)
			}
		}

		slug := strings.TrimSuffix(entry.Name(), ".md")

		var summary string
		body = processNoticeShortcodes(body)

		if idx := strings.Index(body, "<!--more-->"); idx != -1 {
			summaryMd := strings.TrimSpace(body[:idx])
			summary, err = renderMarkdown(summaryMd)
			if err != nil {
				return nil, fmt.Errorf("rendering summary for %s: %w", entry.Name(), err)
			}
		}

		fullContent, err := renderMarkdown(body)
		if err != nil {
			return nil, fmt.Errorf("rendering %s: %w", entry.Name(), err)
		}

		var updatedDate time.Time
		if fm.Updated != "" {
			updatedDate, err = time.Parse(time.RFC3339, fm.Updated)
			if err != nil {
				return nil, fmt.Errorf("parsing updated in %s: %w", entry.Name(), err)
			}
		}

		posts = append(posts, &Post{
			Title:   fm.Title,
			Date:    postDate,
			Updated: updatedDate,
			Tags:    fm.Tags,
			Draft:   fm.Draft,
			Slug:    slug,
			Summary: summary,
			Content: fullContent,
		})
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	return posts, nil
}

// LoadPages reads standalone pages from the pages/ embedded filesystem.
// It skips index.md which is loaded separately.
func LoadPages(pagesFS fs.FS) ([]*Page, error) {
	entries, err := fs.ReadDir(pagesFS, "pages")
	if err != nil {
		return nil, fmt.Errorf("reading pages directory: %w", err)
	}

	var pages []*Page
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		if entry.Name() == "index.md" {
			continue
		}

		data, err := fs.ReadFile(pagesFS, "pages/"+entry.Name())
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", entry.Name(), err)
		}

		fm, body, err := parseFrontmatter(data)
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %w", entry.Name(), err)
		}

		body = processNoticeShortcodes(body)

		content, err := renderMarkdown(body)
		if err != nil {
			return nil, fmt.Errorf("rendering %s: %w", entry.Name(), err)
		}

		slug := strings.TrimSuffix(entry.Name(), ".md")

		var updatedDate time.Time
		if fm.Updated != "" {
			updatedDate, err = time.Parse(time.RFC3339, fm.Updated)
			if err != nil {
				return nil, fmt.Errorf("parsing updated in %s: %w", entry.Name(), err)
			}
		}

		pages = append(pages, &Page{
			Title:       fm.Title,
			Description: fm.Description,
			Updated:     updatedDate,
			TOC:         fm.TOC,
			Menu:        fm.Menu,
			Content:     content,
			Slug:        slug,
		})
	}

	return pages, nil
}

// LoadIndexContent reads pages/index.md for the home page.
func LoadIndexContent(pagesFS fs.FS) (frontmatter, string, error) {
	data, err := fs.ReadFile(pagesFS, "pages/index.md")
	if err != nil {
		return frontmatter{}, "", fmt.Errorf("reading index.md: %w", err)
	}

	fm, body, err := parseFrontmatter(data)
	if err != nil {
		return frontmatter{}, "", fmt.Errorf("parsing index.md: %w", err)
	}

	body = processNoticeShortcodes(body)

	content, err := renderMarkdown(body)
	if err != nil {
		return frontmatter{}, "", fmt.Errorf("rendering index.md: %w", err)
	}

	return fm, content, nil
}
