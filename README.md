### David Tang — technical leader building simple, reliable systems

This repo powers `davidtang.ai`, a production site built as a Svelte SPA deployed on Firebase Hosting. It's intentionally small, well-documented, and production-ready — a representative slice of how I build pragmatic, reliable software.

### What this demonstrates
- **End-to-end ownership**: Frontend UX, backend services, CI/CD, containerization, load balancer routing, SEO optimization, technical SEO implementation.
- **Lean architecture**: Static-first SPA with a tiny JSON content contract to keep iteration fast and cost low.
- **Cloud-native deployment**: Containerized build, Cloud Run, GCLB, CDN-friendly assets, sensible defaults.
- **Operational discipline**: Clear runbooks, incident notes, and repeatable deploys.

### Highlights
- **Tech**: Go server (`main.go`) serves a built Svelte SPA; Tailwind for styling.
- **Content pipeline**: SPA reads a small JSON (`public/content.json` or a GCS URL) to render YouTube + LinkedIn without redeploys.
- **Infra**: Cloud Run (container), fronted by Google Cloud Load Balancer with serverless NEG. Docker multi-stage build. Cloud Build-ready.
- **Docs**: Notes and runbooks (in `notes/`) capture decisions, incidents, and pipeline details.

### Why this matters to your team
- **Delivers quickly with guardrails**: Small, composable pieces; strong defaults; production checks.
- **Breadth with depth**: Hands-on across frontend, backend, and cloud infra; pragmatic choices over novelty.
- **Maintains velocity**: A JSON content contract enables rapid, safe updates without touching the app.

### Live links
- Site: `https://davidtang.ai`
- Consulting: `https://tangible.healthcare`

### Contact
If you’re hiring for pragmatic builders who own problems end-to-end and keep systems running in production, I’d love to chat. Connect via LinkedIn or the consulting site above.

