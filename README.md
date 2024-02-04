# links 
Human Redirector and admin portal. Similar to GoLinks and Trotto.

# Current Features 
- json api
- Inmem storage with hardcoded values!
- nextjs frontend admin portal
    - edit lints


# Planned features
- Okta SSO Integration
- postgres

# Architecture
```mermaid
flowchart LR
    user -- /:shorturl: --> links
    user -- /metrics --> admin
    user -- /api --> admin
    user --> node(frontendapp)
    subgraph linksApp
        direction LR
        links(host:8080)
        admin( host:8081 )
    end

    node --> admin
    linksApp --> db[(postgres)]
```