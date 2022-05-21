# Go Workspaces
- Opinionated about how you do things
  - Done for simplicity in problem solving

- One folder, any name and any dir
  - bin - Binary
  - pkg - Packages
  - src - Source Code
    - domain.tld (e.g. Github.com, Gitlab.liamhardman.com)
      - Domain.tld user
        - Folder with code for project/repo
        - Folder with code for project/repo
  
  - Use go get to retrieve packages
    - Use GOPATH env var to point to workspace dir
      - export GOPATH="/Users/liam/Documents/learn-golang"
      - export PATH="$PATH:/Users/liam/Documents/learn-golang"