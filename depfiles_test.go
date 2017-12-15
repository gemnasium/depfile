package depfile

import (
	"bytes"
	"encoding/json"
	"testing"
)

var jsonDepFiles string = `[
  {
    "name": "pom.xml",
    "filenames": [
      "pom.xml"
    ],
    "url": "http://maven.apache.org/pom.html",
    "package_type": "maven",
    "package_managers": [
      {
        "name": "Maven",
        "url": "http://maven.apache.org"
      }
    ]
  },
  {
    "name": "maven-dependencies.json",
    "filenames": [
      "maven-dependencies.json",
      "gemnasium-maven-plugin.json"
    ],
    "package_type": "maven",
    "package_managers": [
      {
        "name": "Maven",
        "url": "http://maven.apache.org"
      }
    ]
  },
  {
    "name": "Gemfile",
    "filenames": [
      "Gemfile",
      "gems.rb"
    ],
    "url": "http://bundler.io/man/gemfile.5.html",
    "package_type": "gem",
    "package_managers": [
      {
        "name": "Bundler",
        "url": "http://bundler.io"
      }
    ]
  },
  {
    "name": "Gemfile.lock",
    "filenames": [
      "Gemfile.lock",
      "gems.locked"
    ],
    "package_type": "gem",
    "package_managers": [
      {
        "name": "Bundler",
        "url": "http://bundler.io"
      }
    ]
  },
  {
    "name": "gemspec",
    "patterns": [
      "\\.gemspec$"
    ],
    "url": "http://guides.rubygems.org/specification-reference/",
    "package_type": "gem",
    "package_managers": [
      {
        "name": "gem",
        "url": "http://guides.rubygems.org/command-reference/"
      }
    ]
  },
  {
    "name": "package.json",
    "filenames": [
      "package.json"
    ],
    "url": "https://docs.npmjs.com/files/package.json",
    "package_type": "npm",
    "package_managers": [
      {
        "name": "npm",
        "url": "https://www.npmjs.com"
      },
      {
        "name": "Yarn",
        "url": "https://yarnpkg.com"
      }
    ]
  },
  {
    "name": "npm-shrinkwrap.json",
    "filenames": [
      "npm-shrinkwrap.json"
    ],
    "url": "https://docs.npmjs.com/files/shrinkwrap.json",
    "package_type": "npm",
    "package_managers": [
      {
        "name": "npm",
        "url": "https://www.npmjs.com"
      }
    ]
  },
  {
    "name": "package-lock.json",
    "filenames": [
      "package-lock.json"
    ],
    "url": "https://docs.npmjs.com/files/package-lock.json",
    "package_type": "npm",
    "package_managers": [
      {
        "name": "npm",
        "url": "https://www.npmjs.com"
      }
    ]
  },
  {
    "name": "yarn.lock",
    "filenames": [
      "yarn.lock"
    ],
    "url": "https://yarnpkg.com/lang/en/docs/yarn-lock/",
    "package_type": "npm",
    "package_managers": [
      {
        "name": "Yarn",
        "url": "https://yarnpkg.com"
      }
    ]
  },
  {
    "name": "bower.json",
    "filenames": [
      "bower.json"
    ],
    "url": "https://github.com/bower/spec/blob/master/json.md#name",
    "package_type": "bower",
    "package_managers": [
      {
        "name": "Bower",
        "url": "https://bower.io/"
      }
    ]
  },
  {
    "name": "composer.json",
    "filenames": [
      "composer.json"
    ],
    "url": "https://getcomposer.org/doc/04-schema.md",
    "package_type": "packagist",
    "package_managers": [
      {
        "name": "Composer",
        "url": "http://guides.rubygems.org/command-reference/"
      }
    ]
  },
  {
    "name": "composer.lock",
    "filenames": [
      "composer.lock"
    ],
    "package_type": "packagist",
    "package_managers": [
      {
        "name": "Composer",
        "url": "http://guides.rubygems.org/command-reference/"
      }
    ]
  },
  {
    "name": "requirements.txt",
    "filenames": [
      "requires.txt",
      "requirements.txt",
      "requirements.pip"
    ],
    "patterns": [
      "requirements.*\\.txt$"
    ],
    "url": "https://pip.pypa.io/en/stable/reference/pip_install/#requirements-file-format",
    "package_type": "pypi",
    "package_managers": [
      {
        "name": "pip",
        "url": "https://pip.pypa.io/en/stable/"
      }
    ]
  }
]`

func TestMarshalToJSON(t *testing.T) {
	b, err := json.Marshal(DepFiles)
	if err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")

	if out.String() != jsonDepFiles {
		t.Errorf("Wrong JSON export, expecting:\n%s", out.String())
	}
}
