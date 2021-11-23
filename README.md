# Bill-Center

# Structure


- [x] pdf bill qrcode extract and scan
- [ ] bill parse web service **In Progress**
- [ ] User CRUD **Almost**

[Project Layout Reference](https://github.com/golang-standards/project-layout)

```text
├── README.md
├── cmd
│   ├── bill-decode // cli pdf qrcode scan
│   └── bill-server // provide web service
├── go.mod
├── go.sum
├── pkg
│   └── bill-decode // decode lib 
└── resource
    ├── fapiao.png
    └── 凉茶发票.pdf
```


# Installation

1. You need a go development environment setup before everything starts taking off.
2. Use git clone cloing the repo to your local folder.
   `git clone https://github.com/hades300/bill-center`
3. Import document/sql/create.sql to your database.
4. Create configuration file from `cmd/bill-server/config/config.toml.bak`.
   `cp config/config.toml.bak config/config.toml`
   Update config.toml according to your local configurations if necessary.

5. Run command go run main.go `cd cmd/bill-server && go run main.go`

# Usage

## bill qrcode decode

download release

```bash
bill-decode -pdf "invoce.pdf"
```