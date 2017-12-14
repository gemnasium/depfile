package depfile

var (
	pmMaven    = PackageManager{Name: "Maven", URL: "http://maven.apache.org"}
	pmBundler  = PackageManager{Name: "Bundler", URL: "http://bundler.io"}
	pmGem      = PackageManager{Name: "gem", URL: "http://guides.rubygems.org/command-reference/"}
	pmNpm      = PackageManager{Name: "npm", URL: "https://www.npmjs.com"}
	pmYarn     = PackageManager{Name: "Yarn", URL: "https://yarnpkg.com"}
	pmBower    = PackageManager{Name: "Bower", URL: "https://bower.io/"}
	pmComposer = PackageManager{Name: "Composer", URL: "http://guides.rubygems.org/command-reference/"}
	pmPip      = PackageManager{Name: "pip", URL: "https://pip.pypa.io/en/stable/"}
)

var DepFiles = []DepFile{

	// java
	{
		Name: "pom.xml",
		Filenames: []string{
			"pom.xml",
		},
		URL: "http://maven.apache.org/pom.html",
		PackageManagers: []PackageManager{
			pmMaven,
		},
	},

	// ruby
	{
		Name: "Gemfile",
		Filenames: []string{
			"Gemfile",
			"gems.rb",
		},
		URL: "http://bundler.io/man/gemfile.5.html",
		PackageManagers: []PackageManager{
			pmBundler,
		},
	},

	{
		Name: "Gemfile.lock",
		Filenames: []string{
			"Gemfile.lock",
			"gems.locked",
		},
		PackageManagers: []PackageManager{
			pmBundler,
		},
	},

	{
		Name: "gemspec",
		Patterns: []string{
			`\.gemspec$`,
		},
		URL: "http://guides.rubygems.org/specification-reference/",
		PackageManagers: []PackageManager{
			pmGem,
		},
	},

	// npm
	{
		Name: "package.json",
		Filenames: []string{
			"package.json",
		},
		URL: "https://docs.npmjs.com/files/package.json",
		PackageManagers: []PackageManager{
			pmNpm,
			pmYarn,
		},
	},

	{
		Name: "npm-shrinkwrap.json",
		Filenames: []string{
			"npm-shrinkwrap.json",
		},
		URL: "https://docs.npmjs.com/files/shrinkwrap.json",
		PackageManagers: []PackageManager{
			pmNpm,
		},
	},

	{
		Name: "package-lock.json",
		Filenames: []string{
			"package-lock.json",
		},
		URL: "https://docs.npmjs.com/files/package-lock.json",
		PackageManagers: []PackageManager{
			pmNpm,
		},
	},

	{
		Name: "yarn.lock",
		Filenames: []string{
			"yarn.lock",
		},
		URL: "https://yarnpkg.com/lang/en/docs/yarn-lock/",
		PackageManagers: []PackageManager{
			pmYarn,
		},
	},

	// bower
	{
		Name: "bower.json",
		Filenames: []string{
			"bower.json",
		},
		URL: "https://github.com/bower/spec/blob/master/json.md#name",
		PackageManagers: []PackageManager{
			pmBower,
		},
	},

	// php
	{
		Name: "composer.json",
		Filenames: []string{
			"composer.json",
		},
		URL: "https://getcomposer.org/doc/04-schema.md",
		PackageManagers: []PackageManager{
			pmComposer,
		},
	},

	{
		Name: "composer.lock",
		Filenames: []string{
			"composer.lock",
		},
		PackageManagers: []PackageManager{
			pmComposer,
		},
	},

	// python
	{
		Name: "requirements.txt",
		Filenames: []string{
			"requires.txt",
			"requirements.txt",
			"requirements.pip",
		},
		Patterns: []string{
			`requirements.*\.txt$`,
		},
		URL: "https://pip.pypa.io/en/stable/reference/pip_install/#requirements-file-format",
		PackageManagers: []PackageManager{
			pmPip,
		},
	},

	// TODO: https://github.com/naiquevin/pipdeptree

}
