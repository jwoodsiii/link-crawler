package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  PageData
	}{
		{
			name:     "all fields extract",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
	        <h1>Test Title</h1>
	        <p>This is the first paragraph.</p>
	        <a href="/link1">Link 1</a>
	        <img src="/image1.jpg" alt="Image 1">
	    </body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{"https://crawler-test.com/link1"},
				ImageURLs:      []string{"https://crawler-test.com/image1.jpg"},
			},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: error parsing url", i, tc.name)
			}
			actual := extractPageData(tc.inputBody, url)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected: %v, inputURL: %v, actual: %v", i, tc.name, tc.expected, tc.inputURL, actual)
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "relative img",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:      "absolute url",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="https://crawler-test.com/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:      "no images",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body>Just a body</body></html>`,
			expected:  []string{},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: error parsing url", i, tc.name)
			}
			actual, err := getImagesFromHTML(tc.inputBody, url)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected: %v, inputURL: %v, actual: %v", i, tc.name, tc.expected, tc.inputURL, actual)
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "absolute url",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com"},
		},
		{
			name:      "mutliple urls",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a><a href="/xyz"></a></body></html>`,
			expected:  []string{"https://crawler-test.com", "https://crawler-test.com/xyz"},
		},
		{
			name:      "single relative url",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="/xyz"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com/xyz"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: error parsing url", i, tc.name)
			}
			actual, err := getURLsFromHTML(tc.inputBody, url)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected: %v, inputURL: %v, actual: %v", i, tc.name, tc.expected, tc.inputURL, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name: "first paragraph present",
			inputHTML: `<html>
  <body>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>`,
			expected: "Learn to code by building real projects.",
		},
		{
			name: "no paragraphs",
			inputHTML: `<html>
  <body>
    <main>
    </main>
  </body>
</html>`,
			expected: "",
		},
		{
			name: "paragraph with no main",
			inputHTML: `<html>
  <body>
    <p>Learn to code by building real projects.</p>
  </body>
</html>`,
			expected: "Learn to code by building real projects.",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name: "<h1> present",
			inputHTML: `<html>
  <body>
    <h1>Welcome to Boot.dev</h1>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>`,
			expected: "Welcome to Boot.dev",
		},
		{
			name: "<h1> not present",
			inputHTML: `<html>
  <body>
  <h2>You can't learn without practice</h2>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>`,
			expected: "You can't learn without practice",
		},
		{
			name: "<h1>, <h2> missing",
			inputHTML: `<html>
  <body>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>`,
			expected: "",
		},
		{
			name: "no content",
			inputHTML: `<html>
  <body>
  </body>
</html>`,
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
