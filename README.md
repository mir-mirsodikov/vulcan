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

## Examples

```bash
# creates a `about` folder in `app` and adds a `page.tsx` file

vulcan next page about
```

```bash
# creates a `dashboard` folder in `pages/admin/information` and
# adds a `page.{jsx | tsx}` file

vulcan next page dashboard --path admin/information --pages-dir
```

```bash
# creates a `sign-up-form.{jsx | tsx}` file in `components` or
# user defined directory

vulcan next component sign-up-form
```
