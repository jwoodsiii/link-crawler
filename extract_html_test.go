package main

import (
	"testing"
)

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
