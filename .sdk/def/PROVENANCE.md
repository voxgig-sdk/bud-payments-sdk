# API definition provenance

## bud-payments-openapi.json

- **Source:** https://docs.thisisbud.com/llms.txt → the `docs.thisisbud.com/reference`
  pages for the **Payments API**. Bud's documentation site publishes OpenAPI
  **one operation per page**, each page carrying the full
  `openapi`/`info`/`servers`/`security`/`components` header.
- **Publisher:** Bud Financial
- **Retrieved:** 2026-09-17
- **Format:** OpenAPI 3.x
- **Size:** 146738 bytes
- **Coverage:** 14 paths, 17 methods, 0 component schemas —
  assembled from the 17 reference page(s) Bud publishes for this API.

## Why this is its own SDK, and how the file was assembled

Bud publishes **fourteen** separate APIs across 211 reference pages, each with
its own `info.title`: Assess, Authentication, Characteristics, Connect,
Customers, Drive, Enrichment, Financial Data, Goals, Insights, Intelligent
Search, MCP, Payments and Smart Finders. They are separate API surfaces, so
each is its own SDK rather than one client claiming "the Bud API".

Bud offers no single downloadable document, so this file is that API's
reference pages' embedded OpenAPI fragments merged: `paths` and
`components` unioned, `info`/`servers`/`security` taken as published.
Nothing is hand-written and no name is rewritten — the fragments for one API
are one document served in pieces.

Rebuild with `admin/scripts/bud-merge.sh`; do not hand-edit this file.
