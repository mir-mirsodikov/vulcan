<h1>Vulcan</h1>

A CLI tool for generating files and folders for JavaScript meta frameworks like
Next.js, Nuxt.js, SvelteKit, etc.

# Goal

I have been working with various JavaScript frameworks and I noticed that there
is a hassle to creating new pages, components, routes, etc. Each of these
frameworks have their own conventions and it can make things more difficult to
remember. I wanted to create a tool that would allow me to generate these files
and folders with ease.

Frameworks:
- [ ] Next.js
- [ ] SvelteKit
- [ ] Astro
- [ ] Nuxt.js

Features:

- [ ] Create pages
- [ ] Create components
- [ ] Create routes
- [ ] Create layouts
- [ ] Create tests
- [ ] Configuration file

# Docs

## Installation

```bash
$ git clone https://github.com/mir-mirsodikov/vulcan
$ cd vulcan
$ make build
```

> If there is a permission error, try running `sudo make build`

## Usage

### Next.js
Generate pages and components for Next.js. Only supports Next.js 13+ with the `/app` directory.

Create a new page:
```bash
$ vulcan next add page <name> [options]
```

Create a new component:
```bash
$ vulcan next add component <name> [options]
```

Options:

- `--path | -p` - The path to the directory where the file will be created. 
  - This is relative to the `/app` directory for pages.
  - This is relative to the `/components` directory for components.
- `--server-component | -s` - Creates the page or component as a server component.

## Examples

```bash
# creates a `about` folder in `app` and adds a `page.tsx` file

vulcan next add page about
```

```bash
# creates a `about/contact` folder in `app` and adds a `page.jsx` file
# as a server component

vulcan next add page about --path contact --server-component
```

```bash
# creates a `sign-up-form.jsx` file in `components`

vulcan next add component sign-up-form
```

```bash
# creates a `table.jsx` file in `components/table`
# as a server component

vulcan next add component row --path table -s
```
